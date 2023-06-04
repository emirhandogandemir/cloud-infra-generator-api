package routers

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/controllers"
	"github.com/gin-gonic/gin"
)

func SetupVirtualMachinesRoutes(router *gin.Engine) {
	virtualMachinesGroup := router.Group("/vm")
	{
		virtualMachinesGroup.POST("/createazure/:userid",controllers.CreateVmAzureInstanceHandlers)
		virtualMachinesGroup.POST("/createaws/:userid",controllers.CreateVmAwsInstanceHandlers)
		virtualMachinesGroup.GET("/getlistaws/:userid",controllers.GetInstancesHandler)
		virtualMachinesGroup.GET("/getlistazure/:userid",controllers.GetVmAzureInstanceHandlers)
	}
}