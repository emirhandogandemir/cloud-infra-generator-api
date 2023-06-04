package repositories

import (
	"fmt"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/db"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
)

func CreateAzureAccess(azureAccessModel *models.AzureAccessModel) (*models.AzureAccessModel, error) {
	db, err := db.Connect()

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	result := db.Create(&azureAccessModel)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create Azure access model: %w", result.Error)
	}

	return azureAccessModel, nil
}

func GetAllAzureAccess() ([]*models.AzureAccessModel, error) {
	db, err := db.Connect()

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	var azureAccessModel []*models.AzureAccessModel
	result := db.Find(&azureAccessModel)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve Azure access models: %w", result.Error)
	}

	return azureAccessModel, nil
}

func GetByUserIdAzure(userId uint) ([]*models.AzureAccessModel, error) {
	db, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	var azureAccessModel []*models.AzureAccessModel
	result := db.First(&azureAccessModel, userId)
	if result.Error != nil {
		return nil, fmt.Errorf("no Azure access model found for user ID: %d", userId)
	}

	return azureAccessModel, nil

}
