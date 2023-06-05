package routers

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/controllers"
	"github.com/gin-gonic/gin"
)

func SetupKubernetesClusterRoutes(router *gin.Engine) {
	kubernetesClusterGroup := router.Group("/k8s")
	{
		kubernetesClusterGroup.POST("/createeks/:userid",controllers.CreateKubernetesClusterAwsHandlers)
		kubernetesClusterGroup.GET("/geteks/:userid",controllers.GetKubernetesClusterAwsHandlers)
		kubernetesClusterGroup.POST("/deleteeks/:userid",controllers.DeleteKubernetesClusterAwsHandlers)
		//kubernetesClusterGroup.POST("/createaks/:userid",controllers.CreateKubernetesClusterAzureHandlers)
	}
}