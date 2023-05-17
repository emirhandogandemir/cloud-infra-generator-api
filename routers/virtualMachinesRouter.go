package routers

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/controllers"
	"github.com/gin-gonic/gin"
)

func SetupVirtualMachinesRoutes(router *gin.Engine) {
	virtualMachinesGroup := router.Group("/vm")
	{
		virtualMachinesGroup.POST("/createazure",controllers.CreateVmAzureInstanceHandlers)
		virtualMachinesGroup.POST("/createaws",controllers.CreateVmAwsInstanceHandlers)
		virtualMachinesGroup.GET("/getlistaws",controllers.GetInstancesHandler)
		virtualMachinesGroup.GET("/getlistazure",controllers.GetVmAzureInstanceHandlers)
	}
}