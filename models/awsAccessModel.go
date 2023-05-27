package models

import "gorm.io/gorm"

type AwsAccessModel struct {
	gorm.Model
	UserID uint `json:"user_id"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}