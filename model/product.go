package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model  `json:"-"`
	ID          uint             `json:"productId" gorm:"primarykey"`
	Title       string           `json:"title" binding:"required" gorm:"size:32"`
	Sku         string           `json:"sku" binding:"required" gorm:"size:32;uniqueIndex;not null"`
	Description *string          `json:"description" gorm:"size:1024"`
	Barcodes    []ProductBarcode `json:"barcodes" gorm:"foreignKey:ProductId"`
}

type ProductBarcode struct {
	ProductId uint   `json:"productId" gorm:"primaryKey"`
	Barcode   string `json:"barcode" gorm:"size:32;uniqueIndex"`
}

func (c *ProductBarcode) UnmarshalJSON(b []byte) error {
	c.Barcode = string(b)
	return nil
}

func (Product) TableName() string {
	return "product"
}

type ProductsEnvelope struct {
	TotalCount int       `json:"totalCount"`
	Items      []Product `json:"items"`
}
