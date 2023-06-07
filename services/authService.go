package services

import (
	"errors"
	"fmt"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/repositories"
)

func SignUp(username string, email string, password string) (*models.User, error) {
	user := &models.User{
		Username: username,
		Email:    email,
		Password: password,
	}
	return repositories.CreateUser(user)
}

func SignIn(username string, password string) (*models.User, error) {
	user, err := repositories.FindUserByUserName(username)
	if err != nil {
		return nil, fmt.Errorf("username göre userı getirirken hata oluştu")
	}
	// Here you should compare the hashed password with the one in the database
	if user.Password != password {
		return nil, errors.New("parola yanlış girildi")
	}
	return user, nil
}