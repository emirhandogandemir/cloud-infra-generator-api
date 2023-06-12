package controllers

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/compute/mgmt/compute"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/db"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func CreateVmAzureInstanceHandlers (c *gin.Context) {
	userId := c.Param("userid")
	var details models.VirtualMachineAzure
	if err := c.ShouldBindJSON(&details); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error : ": err.Error()})
		return}
	ctx:=context.Background()
	 // Kullanıcının ID'sini burada belirtin
	db, err := db.Connect()
	if err != nil {
		fmt.Println("db baglantısında sorun oluştu")
	}
	var user models.User
	if err := db.Preload("AzureAccessModel").First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return}
	if len(user.AzureAccessModel) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Azure Access not found for the user"})
		return}
	azureAccess := user.AzureAccessModel[0]
	config := auth.NewClientCredentialsConfig(azureAccess.ClientID,azureAccess.ClientSecret,azureAccess.TenantId)
	authorizer,err := config.Authorizer()
	if err != nil{
		log.Fatal("Authorization hatası ,",err)}
	if err == nil{
		fmt.Println("err şu anda boş gardaşım")}
	vmClient := compute.NewVirtualMachinesClient(azureAccess.SubscriptionId)
	vmClient.Authorizer = authorizer
	nicId := "/subscriptions/15608984-3c5b-41dc-9e79-5b17be37947a/resourceGroups/bitirme/providers/Microsoft.Network/networkInterfaces/bitirme"
	vmParameters := compute.VirtualMachine{
		Location: &details.Location,
		Name:     &details.VmName,
		Type:     to.StringPtr("Microsoft.Compute/virtualMachines"),
		VirtualMachineProperties: &compute.VirtualMachineProperties{
			HardwareProfile: &compute.HardwareProfile{
				VMSize: compute.VirtualMachineSizeTypes(details.VmSize),
			},
			OsProfile: &compute.OSProfile{
				ComputerName:  &details.VmName,
				AdminUsername: &details.AdminUsername,
				AdminPassword: &details.AdminPassword,
			},
			StorageProfile: &compute.StorageProfile{
				ImageReference: &compute.ImageReference{
					Publisher: to.StringPtr("Canonical"),
					Offer:     to.StringPtr("UbuntuServer"),
					Sku:       to.StringPtr("18.04-LTS"),
					Version:   to.StringPtr("latest"),
				},
				OsDisk: &compute.OSDisk{
					Name:         to.StringPtr(fmt.Sprintf("%s_os_disk", details.VmName)),
					Caching:      compute.CachingTypesReadWrite,
					CreateOption: compute.DiskCreateOptionTypesFromImage,
					DiskSizeGB:   to.Int32Ptr(100),
					ManagedDisk: &compute.ManagedDiskParameters{
						StorageAccountType: compute.StorageAccountTypesStandardLRS,
					},
				},
			},
			NetworkProfile: &compute.NetworkProfile{
				NetworkInterfaces: &[]compute.NetworkInterfaceReference{
					{
						ID: &nicId,
					},
				},
			},
		},
	}
	_, err = vmClient.CreateOrUpdate(ctx, details.ResourceGroupName, details.VmName, vmParameters)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "VM created successfully"})

}
func GetVmAzureInstanceHandlers (c *gin.Context) {
	userId := c.Param("userid")
	ctx:=context.Background()

	db, err := db.Connect()

	if err != nil {
		fmt.Println("db baglantısında sorun oluştu")
	}
	var user models.User
	if err := db.Preload("AzureAccessModel").First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	if len(user.AzureAccessModel) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Azure Access not found for the user"})
		return
	}

	azureAccess := user.AzureAccessModel[0]

	config := auth.NewClientCredentialsConfig(azureAccess.ClientID,azureAccess.ClientSecret,azureAccess.TenantId)

	authorizer,err := config.Authorizer()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"Authentication Error ": err.Error()})
		return
	}
	vmClient := compute.NewVirtualMachinesClient(azureAccess.SubscriptionId)
	vmClient.Authorizer = authorizer

	filter:="resourceGroup eq 'bitirme'"
	vmList,err := vmClient.ListAll(ctx,"Stopped",filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, vmList)

}