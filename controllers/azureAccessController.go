package controllers

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAzureAccess(c *gin.Context) {
	var azureAccess models.AzureAccessModel
	if err := c.ShouldBindJSON(&azureAccess); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	azureAccessCreated, err := services.CreateAzureAccess(&azureAccess)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": azureAccessCreated})
}

func GetAllAzureAccess(c *gin.Context) {
	azureAccessModels, err := services.GetAllAzureAccess()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": azureAccessModels})
}