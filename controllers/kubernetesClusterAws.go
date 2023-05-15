package controllers

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/gin-gonic/gin"
)
func CreateKubernetesClusterAwsHandlers (c *gin.Context){
  var paramsKubernetes models.KubernetesParamsAws

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		c.JSON(500, gin.H{"error": "Error Loading Default Config: " + err.Error()})
		return
	}
	cfg.Region = "us-east-1"

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