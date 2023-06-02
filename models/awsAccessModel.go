package models

import "gorm.io/gorm"

type AwsAccessModel struct {
	gorm.Model
	UserID uint `json:"user_id"`
	Name string `json:"name"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}