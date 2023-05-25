package models

type AwsAccessModel struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}
