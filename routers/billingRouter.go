package routers

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/controllers"
	"github.com/gin-gonic/gin"
)

func SetupBillingRoutes(router *gin.Engine) {
	billingGroup := router.Group("/billing")
	{
		billingGroup.GET("/getbillingaws/:userid",controllers.GetBillingAwsHandler)
		//billingGroup.POST("/createaks",controllers.CreateKubernetesClusterAzureHandlers)
	}
}
