package main

import (
	"fmt"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/controllers"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//router.GET("/getbillingaws",controllers.GetBillingAwsHandler)
	//router.GET("/getvmlistaws",controllers.GetInstancesHandler)
	//router.POST("/createeks",controllers.CreateKubernetesClusterAwsHandlers)
	//router.POST("/createvmaws",controllers.CreateInstanceHandlers)
	router.GET("/getinstancetypesaws",controllers.GetInstanceTypeHandler)
	//router.POST("/createnodegroupaws",controllers.NodeGroupEksHandlers)
	//router.POST("/createvmazure",controllers.CreateVmAzureInstanceHandlers)
	//router.POST("/createaks",controllers.CreateKubernetesClusterAzureHandlers)
	// API sunucusunu başlatın

	routers.SetupUserRoutes(router)
	routers.SetupVirtualMachinesRoutes(router)
	routers.SetupKubernetesClusterRoutes(router)
	routers.SetupBillingRoutes(router)
	routers.SetupNodeGroupRoutes(router)
	if err := router.Run(":7070"); err != nil {
		fmt.Println("API sunucusu başlatılırken hata oluştu: " + err.Error())
	}
}
