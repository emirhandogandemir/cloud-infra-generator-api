package repositories

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) CreateUser (user *models.User)error{
	return r.DB.Create(user).Error
}

func(r *UserRepository) GetUserById(userId int64)(*models.User,error){
	user:= &models.User{}
	result := r.DB.First(user,userId)
	if result.Error != nil{
		return nil,result.Error
	}
	return user,nil
}