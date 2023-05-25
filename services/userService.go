package services

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/repositories"
)

func CreateUser(user *models.User) (*models.User, error) {
	return repositories.CreateUser(user)
}

func GetAllUsers() ([]*models.User, error) {
	return repositories.GetAllUser()
}