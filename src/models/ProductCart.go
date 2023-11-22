package models

import "gorm.io/gorm"

type ProductCart struct {
	gorm.Model

	Name        string
	Price       float64
	Stock       int64
	Id_customer uint
}
