package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `gorm:"notNull"`
	Price float64 `gorm:"notNull;default:0"`
	Stock int64   `gorm:"notNull;default:0"`
}
