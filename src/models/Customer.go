package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Fullname string `gorm:"notNull"`
	Email    string `gorm:"unique;size:255;notNull"`
	Password string `gorm:"notNull"`
}
