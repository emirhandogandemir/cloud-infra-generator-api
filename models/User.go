package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username   string    `gorm:"size:255;not null;unique" json:"username,required"`
	Email      string    `gorm:"size:100;not null;unique" json:"email,required"`
	Password string `gorm:"not null" json:"password,required"`
}