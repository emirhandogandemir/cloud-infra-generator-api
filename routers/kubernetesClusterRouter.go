package routers

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/controllers"
	"github.com/gin-gonic/gin"
)

func SetupKubernetesClusterRoutes(router *gin.Engine) {
	kubernetesClusterGroup := router.Group("/k8s")
	{
		kubernetesClusterGroup.POST("/createeks",controllers.CreateKubernetesClusterAwsHandlers)
		kubernetesClusterGroup.POST("/createaks",controllers.CreateKubernetesClusterAzureHandlers)
	}
}