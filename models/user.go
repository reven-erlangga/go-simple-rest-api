package models

import "gorm.io/gorm"

type User struct {
	Email string `gorm:"unique"`
	Password string 
	gorm.Model
}