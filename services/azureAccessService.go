package services

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/repositories"
)

<<<<<<< HEAD
func CreateAzureAccess(awsAccessModel *models.AzureAccessModel) (*models.AzureAccessModel, error) {
	return repositories.CreateAzureAccess(awsAccessModel)
=======
func CreateAzureAccess(azureAccessModel *models.AzureAccessModel) (*models.AzureAccessModel, error) {
	return repositories.CreateAzureAccess(azureAccessModel)
>>>>>>> bdf4f6a4c5fdc179e298812ae1126a653e7b482e
}

func GetAllAzureAccess() ([]*models.AzureAccessModel, error) {
	return repositories.GetAllAzureAccess()
<<<<<<< HEAD
}
=======
}
>>>>>>> bdf4f6a4c5fdc179e298812ae1126a653e7b482e
