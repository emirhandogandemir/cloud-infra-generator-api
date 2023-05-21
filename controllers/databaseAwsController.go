package controllers

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
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
func CreateDatabaseAwsHandler(c *gin.Context) {
	var params models.DatabaseAws
	cfg,err:= config.LoadDefaultConfig(context.TODO())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Could not load configuration": err.Error()})
		return
	}
	cfg.Region="us-east-1"

	rdsClient := rds.NewFromConfig(cfg)

	input := &rds.CreateDBInstanceInput{
		DBInstanceIdentifier: &params.DbInstanceIdentifier,
		DBInstanceClass:      &params.DbInstanceClass,
		Engine:               &params.Engine,
		MasterUsername:       &params.MasterUsername,
		MasterUserPassword:   &params.MasterPassword,
		AllocatedStorage:     aws.Int32(20),
		DBName:               &params.DbName,
		AvailabilityZone:     aws.String("us-east-1a"),
		MultiAZ:              aws.Bool(false),
	}

	createOutput, err := rdsClient.CreateDBInstance(context.Background(), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Failed to create instance Rds": err.Error()})
		return
	}

	fmt.Printf("Created RDS instance: %s\n", *createOutput.DBInstance.DBInstanceIdentifier)

	// Oluşturulan RDS örneği hakkında bilgileri alın
	describeInput := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: &params.DbInstanceIdentifier,
	}

	describeOutput, err := rdsClient.DescribeDBInstances(context.Background(), describeInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Failed to describe instances RDS": err.Error()})
		return
	}

	// RDS örneği hakkında bilgileri yazdırın
	for _, instance := range describeOutput.DBInstances {
		fmt.Printf("DB Instance ID: %s\n", *instance.DBInstanceIdentifier)
		fmt.Printf("Engine: %s\n", *instance.Engine)
		//fmt.Printf("Endpoint: %s\n", *instance.Endpoint.Address)
		fmt.Printf("Status: %s\n", *instance.DBInstanceStatus)
		fmt.Println("-----")
	}

	c.JSON(http.StatusOK,describeOutput.DBInstances)

}

func DeleteDatabaseAwsHandler(c *gin.Context) {
	var params models.DatabaseAws
	cfg,err:= config.LoadDefaultConfig(context.TODO())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Could not load configuration": err.Error()})
		return
	}
	cfg.Region="us-east-1"

	rdsClient := rds.NewFromConfig(cfg)

	input := &rds.DeleteDBInstanceInput{
		DBInstanceIdentifier:      aws.String(params.DbInstanceIdentifier), // Silinecek RDS veritabanının kimliği
		SkipFinalSnapshot:         bool(true), // Son anlık görüntü oluşturmadan doğrudan silme
	}


	// RDS veritabanını sil
	_, err = rdsClient.DeleteDBInstance(context.TODO(), input)
	if err != nil {
		fmt.Println("Rds database silinirken bir hata meydana geldi")
	}

	fmt.Println("RDS veritabanı başarıyla silindi!")

	c.JSON(http.StatusOK,params.DbInstanceIdentifier)

}