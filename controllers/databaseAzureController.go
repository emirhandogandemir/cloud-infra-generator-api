package controllers

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2021-06-01/postgresqlflexibleservers"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetDatabaseAzureHandler(c *gin.Context) {

	ctx:=context.Background()

	subscriptionID := "15608984-3c5b-41dc-9e79-5b17be37947a"
	tenantID := "9b4786c5-38d8-4442-b63f-d2c8d41d0e95"
	clientID := "e2c85ab7-e254-43da-ad29-a4799f45f4fc"
	clientSecret := "f_A8Q~GYHeMSmJ-yImc_roVeVseNZ-zoe1I5KdBd"

	config := auth.NewClientCredentialsConfig(clientID,clientSecret,tenantID)

	resourceGroupName:="bitirme"

	authorizer,err := config.Authorizer()
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Authorization Error": err.Error()})
		return
	}
	if err == nil{
		fmt.Println("err şu anda boş gardaşım")
	}
	psqlClient := postgresqlflexibleservers.NewServersClient(subscriptionID)
	psqlClient.Authorizer = authorizer

	dbListResult, err := psqlClient.ListByResourceGroupComplete(ctx,resourceGroupName)
	if err != nil {
		log.Fatalf("Failed to get database list: %v", err)
	}

	for dbListResult.NotDone() {
		fmt.Println("Values", *dbListResult.Value().Name)
		err = dbListResult.NextWithContext(ctx)
		if err != nil {
			log.Fatalf("Failed to get next server in list: %v", err)
		}
	}

	c.JSON(http.StatusOK,dbListResult.NotDone())

}

func CreateDatabaseAzureHandler(c *gin.Context) {
	var params models.DatabaseAzure
	ctx:=context.Background()

	subscriptionID := "15608984-3c5b-41dc-9e79-5b17be37947a"
	tenantID := "9b4786c5-38d8-4442-b63f-d2c8d41d0e95"
	clientID := "e2c85ab7-e254-43da-ad29-a4799f45f4fc"
	clientSecret := "f_A8Q~GYHeMSmJ-yImc_roVeVseNZ-zoe1I5KdBd"

	config := auth.NewClientCredentialsConfig(clientID,clientSecret,tenantID)

	authorizer,err := config.Authorizer()
	if err != nil{
		log.Fatal("Authorization hatası ,",err)
	}
	if err == nil{
		fmt.Println("err şu anda boş gardaşım")
	}
	// PostgreSQL esnek sunucu istemci oluşturma
	psqlClient := postgresqlflexibleservers.NewServersClient(subscriptionID)
	psqlClient.Authorizer = authorizer

	// PostgreSQL sunucusu için yapılandırma ayarları
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
 	var params models.DatabaseAzure
	ctx:=context.Background()

	subscriptionID := "15608984-3c5b-41dc-9e79-5b17be37947a"
	tenantID := "9b4786c5-38d8-4442-b63f-d2c8d41d0e95"
	clientID := "e2c85ab7-e254-43da-ad29-a4799f45f4fc"
	clientSecret := "f_A8Q~GYHeMSmJ-yImc_roVeVseNZ-zoe1I5KdBd"

	config := auth.NewClientCredentialsConfig(clientID,clientSecret,tenantID)

	authorizer,err := config.Authorizer()
	if err != nil{
		log.Fatal("Authorization hatası ,",err)
	}
	if err == nil{
		fmt.Println("err şu anda boş gardaşım")
	}

	psqlClient := postgresqlflexibleservers.NewServersClient(subscriptionID)
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
