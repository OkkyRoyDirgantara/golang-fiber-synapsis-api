package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Fullname string
	Email    string `gorm:"index:,unique;size:255"`
	Password string
}
