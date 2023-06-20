package main

import (
	"fmt"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/controllers"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // frontend'inizin çalıştığı adresi buraya ekleyin.
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}

	router.Use(cors.New(config))

	router.GET("/getinstancetypesaws",controllers.GetInstanceTypeHandler)
	routers.SetupUserRoutes(router)
	routers.SetupVirtualMachinesRoutes(router)
	routers.SetupKubernetesClusterRoutes(router)
	routers.SetupBillingRoutes(router)
	routers.SetupNodeGroupRoutes(router)
	routers.SetupStorageRoutes(router)
	routers.SetupDatabaseRoutes(router)
	routers.SetupAwsAccessRouter(router)
	routers.SetupAzureAccessRouter(router)
	routers.SetupAuthRoutes(router)
	if err := router.Run(":7070"); err != nil {
		fmt.Println("API sunucusu başlatılırken hata oluştu: " + err.Error())
	}
}
