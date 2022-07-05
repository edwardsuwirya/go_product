package model

import "gorm.io/gorm"

type Product struct {
	ProductId   string
	ProductName string
	gorm.Model
}
