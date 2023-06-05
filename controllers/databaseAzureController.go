package controllers

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2021-06-01/postgresqlflexibleservers"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/db"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetDatabaseAzureHandler(c *gin.Context) {

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

	config := auth.NewClientCredentialsConfig(azureAccess.ClientID,azureAccess.ClientSecret,azureAccess.TenantId)

	resourceGroupName:="bitirme"

	authorizer,err := config.Authorizer()
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Authorization Error": err.Error()})
		return
	}
	if err == nil{
		fmt.Println("err şu anda boş gardaşım")
	}
	psqlClient := postgresqlflexibleservers.NewServersClient(azureAccess.SubscriptionId)
	psqlClient.Authorizer = authorizer

	dbListResult, err := psqlClient.ListByResourceGroupComplete(ctx,resourceGroupName)
	if err != nil {
		log.Fatalf("Failed to get database list: %v", err)
	}

	var dblist []postgresqlflexibleservers.Server
	for dbListResult.NotDone() {
		//fmt.Println("Values", *dbListResult.Value().Name)
		db := dbListResult.Value()
		dblist = append(dblist,db)
		err = dbListResult.NextWithContext(ctx)
		if err != nil {
			log.Fatalf("Failed to get next server in list: %v", err)
		}
	}

	c.JSON(http.StatusOK,dblist)

}

func CreateDatabaseAzureHandler(c *gin.Context) {
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

	config := auth.NewClientCredentialsConfig(azureAccess.ClientID,azureAccess.ClientSecret,azureAccess.TenantId)

	authorizer,err := config.Authorizer()
	if err != nil{
		log.Fatal("Authorization hatası ,",err)
	}
	if err == nil{
		fmt.Println("err şu anda boş gardaşım")
	}
	var params models.DatabaseAzure
	if err := c.ShouldBindJSON(&params); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error : ":err.Error()})
		return
	}

	psqlClient := postgresqlflexibleservers.NewServersClient(azureAccess.SubscriptionId)
	psqlClient.Authorizer = authorizer

	serverProperties := postgresqlflexibleservers.ServerProperties{
		AdministratorLogin:         to.StringPtr(params.AdminName),
		AdministratorLoginPassword: to.StringPtr(params.AdminPassword),
		Version: postgresqlflexibleservers.ServerVersion(params.Version),
		Storage: &postgresqlflexibleservers.Storage{
			StorageSizeGB: to.Int32Ptr(128),
		},
	}
	future, err := psqlClient.Create(ctx, params.ResourceGroup, params.ServerName, postgresqlflexibleservers.Server{
		Location: to.StringPtr(params.Location),
		ServerProperties: &serverProperties,
		Sku: &postgresqlflexibleservers.Sku{
			Name: to.StringPtr("Standard_D2s_v3"),
			Tier: postgresqlflexibleservers.SkuTierGeneralPurpose,
		},
	})
	if err != nil {
		log.Fatalf("Failed to start create server operation: %s", err)
	}
	if err ==nil{
		fmt.Println("Future : ", future)
	}

	err = future.WaitForCompletionRef(context.Background(), psqlClient.Client)
	if err != nil {
		log.Fatalf("Failed to finish create server operation: %s", err)
	}

	fmt.Printf("Successfully created PostgreSQL server: %s", params.ServerName)

	c.JSON(http.StatusOK,"Successfull created Postgresql Server")

}

func DeleteDatabaseAzureHandler(c *gin.Context) {
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

	config := auth.NewClientCredentialsConfig(azureAccess.ClientID,azureAccess.ClientSecret,azureAccess.TenantId)

	authorizer,err := config.Authorizer()
	if err != nil{
		log.Fatal("Authorization hatası ,",err)
	}
	if err == nil{
		fmt.Println("err şu anda boş gardaşım")
	}
	var params models.DatabaseAzure
	if err := c.ShouldBindJSON(&params); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error : ":err.Error()})
		return
	}

	psqlClient := postgresqlflexibleservers.NewServersClient(azureAccess.SubscriptionId)
	psqlClient.Authorizer = authorizer

	future, err := psqlClient.Delete(ctx, params.ResourceGroup, params.ServerName)
	if err != nil {
		log.Fatalf("Failed to start delete server operation: %v", err)
	}

	err = future.WaitForCompletionRef(ctx, psqlClient.Client)
	if err != nil {
		log.Fatalf("Failed to finish delete server operation: %v", err)
	}

	fmt.Printf("Successfully deleted PostgreSQL server: %s", params.ServerName)

}
