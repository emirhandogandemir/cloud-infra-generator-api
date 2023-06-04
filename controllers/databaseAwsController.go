package controllers

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/db"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDatabaseAwsHandler(c *gin.Context) {
	userId := c.Param("userid")
	db,err := db.Connect()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}
	var user models.User
	if err := db.Preload("AwsAccessModel").First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	if len(user.AwsAccessModel) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "AWS Access not found for the user"})
		return
	}

	awsAccess := user.AwsAccessModel[0]
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx,config.WithRegion("us-east-1"),config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccess.AccessKey,awsAccess.SecretKey,"")))
	if err != nil {
		fmt.Println("Couldn't load default configuration.",err)
	}

	client := rds.NewFromConfig(cfg)
	input := &rds.DescribeDBInstancesInput{}
	output, err := client.DescribeDBInstances(context.TODO(), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status error to access to RDS database": err.Error()})
		return
	}

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
	userId := c.Param("userid")
	db,err := db.Connect()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}
	var user models.User
	if err := db.Preload("AwsAccessModel").First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	if len(user.AwsAccessModel) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "AWS Access not found for the user"})
		return
	}

	awsAccess := user.AwsAccessModel[0]

	var params models.DatabaseAws
	if err := c.ShouldBindJSON(&params); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error : ":err.Error()})
		return
	}

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx,config.WithRegion("us-east-1"),config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccess.AccessKey,awsAccess.SecretKey,"")))
	if err != nil {
		fmt.Println("Couldn't load default configuration.",err)
	}
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
	userId := c.Param("userid")
	db,err := db.Connect()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}
	var user models.User
	if err := db.Preload("AwsAccessModel").First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	if len(user.AwsAccessModel) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "AWS Access not found for the user"})
		return
	}

	awsAccess := user.AwsAccessModel[0]

	var params models.DatabaseAws
	if err := c.ShouldBindJSON(&params); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error : ":err.Error()})
		return
	}
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx,config.WithRegion("us-east-1"),config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccess.AccessKey,awsAccess.SecretKey,"")))
	if err != nil {
		fmt.Println("Couldn't load default configuration.",err)
	}

	rdsClient := rds.NewFromConfig(cfg)

	input := &rds.DeleteDBInstanceInput{
		DBInstanceIdentifier:      aws.String(params.DbInstanceIdentifier),
		SkipFinalSnapshot:         bool(true),
	}

	_, err = rdsClient.DeleteDBInstance(context.TODO(), input)
	if err != nil {
		fmt.Println("Rds database silinirken bir hata meydana geldi")
	}

	fmt.Println("RDS veritabanı başarıyla silindi!")

	c.JSON(http.StatusOK,params.DbInstanceIdentifier)

}