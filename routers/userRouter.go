package routers

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/get", controllers.GetAllUsers)
		userGroup.POST("/create",controllers.CreateUser)
	}
}
