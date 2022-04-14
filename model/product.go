package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"strconv"
)

type Product struct {
	gorm.Model  `json:"-"`
	ID          uint             `json:"productId" gorm:"primarykey"`
	Title       string           `json:"title" binding:"required" gorm:"size:32"`
	Sku         string           `json:"sku" binding:"required" gorm:"size:32;uniqueIndex;not null"`
	Description *string          `json:"description" gorm:"size:1024"`
	Price       decimal.Decimal  `json:"price" gorm:"type:decimal(12,2)"`
	Barcodes    []ProductBarcode `json:"barcodes" gorm:"foreignKey:ProductId"`
}

func (Product) TableName() string {
	return "product"
}

type ProductBarcode struct {
	ProductId uint   `json:"productId" gorm:"primaryKey"`
	Barcode   string `json:"barcode" gorm:"size:32;primaryKey;uniqueIndex"`
}

func (c *ProductBarcode) UnmarshalJSON(b []byte) error {
	str, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	c.Barcode = str
	return nil
}

func (c *ProductBarcode) MarshalJSON() ([]byte, error) {
	quoted := strconv.Quote(c.Barcode)
	return []byte(quoted), nil
}

type ProductsEnvelope struct {
	TotalCount int       `json:"totalCount"`
	Items      []Product `json:"items"`
}
