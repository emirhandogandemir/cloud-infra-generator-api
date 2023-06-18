package repositories

import (
	"fmt"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/db"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
)

func CreateUser(user *models.User) (*models.User, error) {
	db, err := db.Connect()

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	result := db.Create(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create user: %w", result.Error)
	}

	return user, nil
}

func GetAllUser() ([]*models.User, error) {
	db, err := db.Connect()

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	var users []*models.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", result.Error)
	}

	return users, nil
}
func GetUserById(userId int) (*models.User, error) {
	db, err := db.Connect()

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	var user models.User
	result := db.Preload("AwsAccessModel").Preload("AzureAccessModel").First(&user, userId)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve user: %w", result.Error)
	}

	return &user, nil
}

func FindUserByUserName(username string)(*models.User,error){
	db,err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user,nil
}
