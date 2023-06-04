package repositories

import (
	"fmt"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/db"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
)

func CreateAwsAccess(awsAccessModel *models.AwsAccessModel) (*models.AwsAccessModel, error) {
	db, err := db.Connect()

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err.Error)
	}

	result := db.Create(&awsAccessModel)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve AWS access models: %w", result.Error)
	}

	return awsAccessModel, nil
}

func GetAllAwsAccess() ([]*models.AwsAccessModel, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	var awsAccessModel []*models.AwsAccessModel
	result := db.Find(&awsAccessModel)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve AWS access models: %w", result.Error)
	}
	return awsAccessModel, nil
}

func GetByUserIdAws(userId uint) ([]*models.AwsAccessModel, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	var awsAccessModel []*models.AwsAccessModel
	result := db.First(&awsAccessModel, userId)
	if result.Error != nil {
		return nil, fmt.Errorf("no AWS access model found for user ID: %d", userId)
	}

	return awsAccessModel, nil

}
