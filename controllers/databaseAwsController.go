package controllers

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDatabaseAwsHandler(c *gin.Context) {
	cfg,err:= config.LoadDefaultConfig(context.TODO())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status load configuration error": err.Error()})
		return
	}
	cfg.Region="us-east-1"

	// RDS istemcisini oluştur
	client := rds.NewFromConfig(cfg)

	// RDS veritabanlarını listeleyen sorguyu oluştur
	input := &rds.DescribeDBInstancesInput{}

	// RDS veritabanlarını al
	output, err := client.DescribeDBInstances(context.TODO(), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status error to access to RDS database": err.Error()})
		return
	}

	// RDS veritabanlarını döngüyle gezerek bilgileri yazdır
	for _, db := range output.DBInstances {
		fmt.Println("DB Name:", aws.ToString(db.DBInstanceIdentifier))
		fmt.Println("Engine:", aws.ToString(db.Engine))
		fmt.Println("Status:", aws.ToString(db.DBInstanceStatus))
		fmt.Println("Endpoint:", aws.ToString(db.Endpoint.Address))
		fmt.Println("------------------------------")
	}

	c.JSON(http.StatusOK,output.DBInstances)

}
