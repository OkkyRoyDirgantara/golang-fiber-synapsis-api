package models

import "gorm.io/gorm"

type ProductCart struct {
	gorm.Model
	Quantity   int      `gorm:"notNull;default:0"`
	CustomerId uint     `gorm:"notNull"`
	Customer   Customer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId  uint     `gorm:"notNull"`
	Product    Product  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
