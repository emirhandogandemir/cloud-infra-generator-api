package controllers

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAwsAccess(c *gin.Context) {
	var awsAccess models.AwsAccessModel
	if err := c.ShouldBindJSON(&awsAccess); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	awssAccessCreated, err := services.CreateAwsAccess(&awsAccess)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": awssAccessCreated})
}

func GetAllAwsAccess(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}
