package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string    `gorm:"size:255;not null;unique" json:"username,required"`
	Email      string    `gorm:"size:100;not null;unique" json:"email,required"`
	Password string `gorm:"not null" json:"password,required"`
	AwsAccessModel []AwsAccessModel `json:"aws_accesses" gorm:"foreignkey:UserID"`
}