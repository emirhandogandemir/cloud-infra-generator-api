package controllers

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateStorageAwsHandler(c *gin.Context) {
  var params models.StorageAws

	if err := c.ShouldBindJSON(&params); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error : ":err.Error()})
		return
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create cluster: " + err.Error()})
		return
	}
	cfg.Region="us-east-1"
	client := s3.NewFromConfig(cfg)

	createBucketOutput, err := client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String("my-bucket-emirhandgndmr51-bitirme2"),
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create cluster: " + err.Error()})
		return
	}

	log.Println("Bucket Created Successfully!", createBucketOutput)

	c.JSON(http.StatusOK, createBucketOutput)
}

