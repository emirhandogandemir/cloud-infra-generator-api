package services

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/repositories"
)

func CreateAwsAccess(awsAccessModel *models.AwsAccessModel) (*models.AwsAccessModel, error) {
	return repositories.CreateAwsAccess(awsAccessModel)
}

func GetAllAwsAccess() ([]*models.AwsAccessModel, error) {
	return repositories.GetAllAwsAccess()
}