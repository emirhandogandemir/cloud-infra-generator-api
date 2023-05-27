package routers

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/controllers"
	"github.com/gin-gonic/gin"
)

func SetupAzureAccessRouter(router *gin.Engine) {
	azureAccessGroup := router.Group("/azureaccess")
	{
		azureAccessGroup.GET("/get", controllers.GetAllAzureAccess)
		azureAccessGroup.POST("/create",controllers.CreateAzureAccess)
	}
}

