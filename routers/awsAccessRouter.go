package routers

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/controllers"
	"github.com/gin-gonic/gin"
)

func SetupAwsAccessRouter(router *gin.Engine) {
	userGroup := router.Group("/awsaccess")
	{
		userGroup.GET("/get", controllers.GetAllUsers)
		userGroup.POST("/create",controllers.CreateUser)
	}
}