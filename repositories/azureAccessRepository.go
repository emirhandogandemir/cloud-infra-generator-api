package repositories

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/db"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
)

func CreateAzureAccess(azureAccessModel *models.AzureAccessModel) (*models.AzureAccessModel, error) {
	db, err := db.Connect()

	if err != nil {
		return nil, err
	}

	result := db.Create(&azureAccessModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return azureAccessModel, nil
}

func GetAllAzureAccess() ([]*models.AzureAccessModel, error) {
	db, err := db.Connect()

	if err != nil {
		return nil, err
	}

	var azureAccessModel []*models.AzureAccessModel
	result := db.Find(&azureAccessModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return azureAccessModel, nil
}