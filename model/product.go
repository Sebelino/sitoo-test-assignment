package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model  `json:"-"`
	Title       string  `json:"title" binding:"required" gorm:"size:32"`
	Sku         string  `json:"sku" binding:"required" gorm:"size:32;uniqueIndex;not null"`
	Description *string `json:"description" gorm:"size:1024"`
}

func (Product) TableName() string {
	return "product"
}

type ProductsEnvelope struct {
	TotalCount int       `json:"totalCount"`
	Items      []Product `json:"items"`
}
