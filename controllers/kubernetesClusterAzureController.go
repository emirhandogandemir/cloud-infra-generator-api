package controllers

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/containerservice/mgmt/containerservice"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func CreateKubernetesClusterAzureHandlers (c *gin.Context){
	var params models.KubernetesParamsAzure

	if err := c.ShouldBindJSON(&params); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscriptionID := "15608984-3c5b-41dc-9e79-5b17be37947a"
	tenantID := "9b4786c5-38d8-4442-b63f-d2c8d41d0e95"
	clientID := "e2c85ab7-e254-43da-ad29-a4799f45f4fc"
	clientSecret := "f_A8Q~GYHeMSmJ-yImc_roVeVseNZ-zoe1I5KdBd"

	config := auth.NewClientCredentialsConfig(clientID, clientSecret, tenantID)

	authorizer, err := config.Authorizer()
	if err != nil {
		log.Fatal("Authorization hatası ,", err)
	}

	aksClient := containerservice.NewManagedClustersClient(subscriptionID)
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
				ClientID: to.StringPtr(clientID),
				Secret:   to.StringPtr(clientSecret),
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