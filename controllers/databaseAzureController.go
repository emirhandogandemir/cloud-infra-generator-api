package controllers

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2021-06-01/postgresqlflexibleservers"
	"github.com/Azure/go-autorest/autorest/azure/auth"
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