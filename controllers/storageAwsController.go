package controllers

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/db"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateStorageAwsHandler(c *gin.Context) {
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


	var params models.StorageAws

	if err := c.ShouldBindJSON(&params); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error : ":err.Error()})
		return
	}
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx,config.WithRegion(params.Location),config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccess.AccessKey,awsAccess.SecretKey,"")))
	if err != nil {
		fmt.Println("Couldn't load default configuration.",err)
	}
	client := s3.NewFromConfig(cfg)

	createBucketOutput, err := client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(params.StorageName),
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create bucket: " + err.Error()})
		return
	}

	log.Println("Bucket Created Successfully!", createBucketOutput)

	c.JSON(http.StatusOK, createBucketOutput)
}

func ListStorageAwsHandler(c *gin.Context) {
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

	client := s3.NewFromConfig(cfg)

	listBucketsOutput, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to list storage s3 list: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, listBucketsOutput.Buckets)
}

func DeleteStorageAwsHandler(c *gin.Context) {
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


	var params models.StorageAws
	if err := c.ShouldBindJSON(&params); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error : ":err.Error()})
		return
	}
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx,config.WithRegion(params.Location),config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccess.AccessKey,awsAccess.SecretKey,"")))
	if err != nil {
		fmt.Println("Couldn't load default configuration.",err)
	}
	client := s3.NewFromConfig(cfg)

	bucketName := params.StorageName

	_, err = client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: &bucketName,


	})

	if err != nil {
		c.JSON(500, gin.H{"error": "Deleted to bucket: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, bucketName)
}