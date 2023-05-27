package services

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/repositories"
)

func CreateAzureAccess(azureAccessModel *models.AzureAccessModel) (*models.AzureAccessModel, error) {
	return repositories.CreateAzureAccess(azureAccessModel)
}

func GetAllAzureAccess() ([]*models.AzureAccessModel, error) {
	return repositories.GetAllAzureAccess()
}