package routers

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/controllers"
	"github.com/gin-gonic/gin"
)

func SetupStorageRoutes(router *gin.Engine) {
	storageGroup := router.Group("/storage")
	{
		storageGroup.POST("/createaws", controllers.CreateStorageAwsHandler)
		storageGroup.GET("/getlistaws",controllers.ListStorageAwsHandler)
		storageGroup.POST("/deleteaws",controllers.DeleteStorageAwsHandler)
	}
}
