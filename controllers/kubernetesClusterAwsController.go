package controllers

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/db"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/gin-gonic/gin"
	"net/http"
)
func CreateKubernetesClusterAwsHandlers (c *gin.Context){
	userId:= c.Param("userid")
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

  var paramsKubernetes models.KubernetesParamsAws

	if err := c.ShouldBindJSON(&paramsKubernetes); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error : ":err.Error()})
		return
	}

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx,config.WithRegion("us-east-1"),config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccess.AccessKey,awsAccess.SecretKey,"")))
	if err != nil {
		fmt.Println("Couldn't load default configuration.",err)
	}
	svc := eks.NewFromConfig(cfg)

	params := &eks.CreateClusterInput{
		Name: aws.String(paramsKubernetes.ClusterName),
		Version: aws.String(paramsKubernetes.ClusterVersion),
		//RoleArn: aws.String("arn:aws:iam::416011088332:role/aws-service-role/eks.amazonaws.com/AWSServiceRoleForAmazonEKS"),
		RoleArn: aws.String("arn:aws:iam::416011088332:role/csm_role"),
		ResourcesVpcConfig: &types.VpcConfigRequest{
			SecurityGroupIds: []string{"sg-016d835889a7cc312"},
			SubnetIds: []string{"subnet-0951bebfbd524097b","subnet-0b5238adedeb97cbd"},
		},


	}
	resp, err := svc.CreateCluster(context.Background(), params)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create cluster: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Cluster created successfully",
		"arn":     *resp.Cluster.Arn,
		"name":    *resp.Cluster.Name,
	})

}

func GetKubernetesClusterAwsHandlers (c *gin.Context){
	userId:= c.Param("userid")
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
	svc := eks.NewFromConfig(cfg)
	resp, err := svc.ListClusters(context.Background(), &eks.ListClustersInput{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list clusters: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"clusters": resp.Clusters,
	})


}

func DeleteKubernetesClusterAwsHandlers (c *gin.Context){

	userId:= c.Param("userid")
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
	var paramsKubernetes models.KubernetesParamsAws

	if err := c.ShouldBindJSON(&paramsKubernetes); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error : ":err.Error()})
		return
	}
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx,config.WithRegion("us-east-1"),config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccess.AccessKey,awsAccess.SecretKey,"")))
	if err != nil {
		fmt.Println("Couldn't load default configuration.",err)
	}
	svc := eks.NewFromConfig(cfg)
	_, err = svc.DeleteCluster(context.Background(), &eks.DeleteClusterInput{
		Name: &paramsKubernetes.ClusterName,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete cluster: " + err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete cluster: " + err.Error()})
		return
	}


}