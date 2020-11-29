package model

import (
	"gorm.io/gorm"
)

// User :  User model
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(64)"`
	Username  string `gorm:"type:varchar(64);unique"`
	Email     string `gorm:"type:varchar(256);unique"`
	Password  string `gorm:"type:varchar(256)"`
	CompanyID uint64
	Company   Company `gorm:"foreignKey:CompanyID"`
}

// Company :  Company model
type Company struct {
	gorm.Model
	Name    string
	City    string `gorm:"type:varchar(64)"`
	State   string `gorm:"type:varchar(64)"`
	Country string `gorm:"type:varchar(64)"`
}
