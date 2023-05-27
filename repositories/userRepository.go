package repositories

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/db"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"

)

func CreateUser(user *models.User) (*models.User, error) {
	db, err := db.Connect()

	if err != nil {
		return nil, err
	}

	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func GetAllUser() ([]*models.User, error) {
	db, err := db.Connect()

	if err != nil {
		return nil, err
	}

	var users []*models.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

