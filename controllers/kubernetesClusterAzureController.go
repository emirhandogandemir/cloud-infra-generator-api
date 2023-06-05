package controllers

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/containerservice/mgmt/containerservice"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/db"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func CreateKubernetesClusterAzureHandlers (c *gin.Context){
	userId := c.Param("userid")
	ctx:=context.Background()

	db, err := db.Connect()

	if err != nil {
		fmt.Println("db baglantısında sorun oluştu")
	}
	var user models.User
	if err := db.Preload("AzureAccessModel").First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	if len(user.AzureAccessModel) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Azure Access not found for the user"})
		return
	}

	azureAccess := user.AzureAccessModel[0]

	var params models.KubernetesParamsAzure

	if err := c.ShouldBindJSON(&params); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	config := auth.NewClientCredentialsConfig(azureAccess.ClientID, azureAccess.ClientSecret, azureAccess.TenantId)

	authorizer, err := config.Authorizer()
	if err != nil {
		log.Fatal("Authorization hatası ,", err)
	}

	aksClient := containerservice.NewManagedClustersClient(azureAccess.SubscriptionId)
	aksClient.Authorizer = authorizer

	resourceGroupName := params.ResourceGroupName
	clusterName := params.ClusterName
	location := params.Location
	vmSize := params.VmSize

	agentPoolProfiles := []containerservice.ManagedClusterAgentPoolProfile{
		{
			Name:   to.StringPtr("systempool"),
			Count:  to.Int32Ptr(2),
			VMSize: &vmSize,
			// AKS kümesi için sistem havuzu olarak işaretleyin
			Mode: containerservice.System,
		},
	}


	// AKS kümesi için yapılandırmaları sağlayın
	cluster := containerservice.ManagedCluster{
		Location: &location,
		ManagedClusterProperties: &containerservice.ManagedClusterProperties{
			KubernetesVersion: to.StringPtr(params.KubernetesVersion),
			DNSPrefix:         to.StringPtr(params.DnsPrefixName),
			AgentPoolProfiles: &agentPoolProfiles,
			ServicePrincipalProfile: &containerservice.ManagedClusterServicePrincipalProfile{
				ClientID: to.StringPtr(azureAccess.ClientID),
				Secret:   to.StringPtr(azureAccess.ClientSecret),
			},
		},
	}
	// AKS kümesi oluşturma işlemini başlatın
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	future, err := aksClient.CreateOrUpdate(ctx, resourceGroupName, clusterName, cluster)
	if err != nil {
		log.Fatal("AKS kümesi oluşturma başlatılırken hata oluştu: ", err)
	}

	// Oluşturma işlemi tamamlanana kadar bekleyin
	err = future.WaitForCompletionRef(ctx, aksClient.Client)
	if err != nil {
		log.Fatal("AKS kümesi oluşturma işlemi beklenirken hata oluştu: ", err)
	}

	// Oluşturulan AKS kümesinin ayrıntılarını alın
	createdCluster, err := future.Result(aksClient)
	if err != nil {
		log.Fatal("AKS kümesi oluştu")

	}

	fmt.Println("Created ClusterName :",createdCluster.Name)

	c.JSON(http.StatusOK, createdCluster.Name)

}