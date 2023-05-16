package controllers

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NodeGroupEksHandlers(c *gin.Context) {
	var params models.NodeGroup

	if err := c.ShouldBindJSON(&params); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cfg, err:= config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error Loading Default Config",err)
	}
	cfg.Region="us-east-1"

	nodeGroup:= eks.NewFromConfig(cfg)

	createNodeGroupInput:= &eks.CreateNodegroupInput{
		ClusterName: aws.String(params.ClusterName),
		NodegroupName: aws.String(params.NodegroupName),
		NodeRole: aws.String("arn:aws:iam::416011088332:role/csm_role"),
		ScalingConfig: &types.NodegroupScalingConfig{
			DesiredSize: aws.Int32(params.DesiredState),
			MaxSize:     aws.Int32(params.MaxSize),
			MinSize:     aws.Int32(params.MinSize),
		},
		Subnets: []string{
			"subnet-0951bebfbd524097b",
			"subnet-0b5238adedeb97cbd",
		},
		InstanceTypes: []string{
			"t2.micro",
			"t3.small",
		},
		RemoteAccess: &types.RemoteAccessConfig{
			Ec2SshKey: &params.KeyName,
		},
	}

	createNodeGroupOutput, err := nodeGroup.CreateNodegroup(context.Background(),createNodeGroupInput)
	if err != nil {
		fmt.Println("Failed to create nodeGroup",err)
	}

	fmt.Println("NodeGroup oluşturma işlemi başarılı şekilde olmuştur",createNodeGroupOutput.Nodegroup.NodegroupName)

	c.JSON(http.StatusOK, createNodeGroupOutput)

}