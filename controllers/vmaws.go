package controllers

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
)

func GetInstancesHandler(c *gin.Context) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Couldn't load default configuration.", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AWS yapılandırması yüklenirken hata oluştu"})
		return
	}
	cfg.Region = "us-east-1"

	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeInstancesInput{}
	result, err := client.DescribeInstances(context.TODO(), input)
	if err != nil {
		log.Fatal("Instances listelenirken hata: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "VM listesi alınırken hata oluştu"})
		return
	}
	instances := []map[string]interface{}{}
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			vm := map[string]interface{}{
				"InstanceID":     aws.ToString(instance.InstanceId),
				"InstanceType":   string(instance.InstanceType),
				"State":          string(instance.State.Name),
				"PublicDNSName":  aws.ToString(instance.PublicDnsName),
				// Daha fazla istediğiniz bilgiyi buraya ekleyebilirsiniz
			}
			instances = append(instances, vm)
		}
	}

	c.JSON(http.StatusOK, gin.H{"instances": instances})

}
func GetInstanceTypeHandler(c *gin.Context){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AWS yapılandırması yüklenirken hata oluştu"})
		return
	}
	cfg.Region = "us-east-1"

	svc := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeInstanceTypesInput{}

	result, err := svc.DescribeInstanceTypes(context.TODO(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "EC2 örnek tipleri alınırken hata oluştu"})
		return
	}

	instanceTypes := []string{}
	for _, instanceType := range result.InstanceTypes {
		instanceTypes = append(instanceTypes, string(instanceType.InstanceType))
	}
	c.JSON(http.StatusOK, gin.H{"instance_types": instanceTypes})
}


func CreateInstanceHandlers (c *gin.Context){
	var details models.VirtualMachine
	if err := c.ShouldBindJSON(&details); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error : ":err.Error()})
		return
	}
	cfg,err:= config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Couldn't load default configuration.",err)
	}
	cfg.Region="us-east-1"

	client := ec2.NewFromConfig(cfg)

	input:= &ec2.RunInstancesInput{
		ImageId:      aws.String(details.ImageId),
		InstanceType: types.InstanceTypeT2Micro,
		MaxCount:     aws.Int32(1),
		MinCount:     aws.Int32(1),
	}

	result ,err := client.RunInstances(context.TODO(),input)
	if err != nil{
		fmt.Println("Error launching instances",err)
		return
	}
	c.JSON(200, gin.H{
		"message": "VM successfully created",
		"instanceId": result.Instances[0].InstanceId,
	})

}