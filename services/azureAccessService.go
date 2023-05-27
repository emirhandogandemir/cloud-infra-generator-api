package services

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/repositories"
)

func CreateAzureAccess(awsAccessModel *models.AzureAccessModel) (*models.AzureAccessModel, error) {
	return repositories.CreateAzureAccess(awsAccessModel)

}

func GetAllAzureAccess() ([]*models.AzureAccessModel, error) {
	return repositories.GetAllAzureAccess()
}
