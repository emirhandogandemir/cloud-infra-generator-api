package controllers

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createtedUser, err := services.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": createtedUser})
}

func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}
func GetUserById(c *gin.Context){
	userId:= c.Param("id")
	id,err:= strconv.Atoi(userId)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":"Invalid user ID"})
	}
	user,err:= services.GetUserById(id)
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"message":"User not found"})
	}
	c.JSON(http.StatusOK,user)

}
