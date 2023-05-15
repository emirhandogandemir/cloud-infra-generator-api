package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/getvmlist",controllers.GetInstancesHandler)

	router.POST("/createvm",controllers.CreateInstanceHandlers)
	router.GET("/getinstancetypes",controllers.GetInstanceTypeHandler)
	// API sunucusunu başlatın
	if err := router.Run(":7070"); err != nil {
		fmt.Println("API sunucusu başlatılırken hata oluştu: " + err.Error())
	}
}
