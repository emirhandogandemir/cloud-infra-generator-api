package routers

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/controllers"
	"github.com/gin-gonic/gin"
)

func SetupAwsAccessRouter(router *gin.Engine) {
	awsAccessGroup := router.Group("/awsaccess")
	{
		awsAccessGroup.GET("/get", controllers.GetAllAwsAccess)
		awsAccessGroup.POST("/create",controllers.CreateAwsAccess)
	}
}