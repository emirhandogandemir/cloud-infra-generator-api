package controllers

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	user := models.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}

	c.JSON(http.StatusOK, user)
}
