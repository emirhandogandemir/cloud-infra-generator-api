package main

import (
	"fmt"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/controllers"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/getinstancetypesaws",controllers.GetInstanceTypeHandler)
	routers.SetupUserRoutes(router)
	routers.SetupVirtualMachinesRoutes(router)
	routers.SetupKubernetesClusterRoutes(router)
	routers.SetupBillingRoutes(router)
	routers.SetupNodeGroupRoutes(router)
	routers.SetupStorageRoutes(router)
	routers.SetupDatabaseRoutes(router)
	if err := router.Run(":7070"); err != nil {
		fmt.Println("API sunucusu başlatılırken hata oluştu: " + err.Error())
	}
}
