package model

import "time"

type Product struct {
	ProductId   uint      `json:"productId" gorm:"primarykey"`
	Title       string    `json:"title"`
	Sku         string    `json:"sku"`
	Created     time.Time `json:"created"`
	Description *string   `json:"description"`
}

func (Product) TableName() string {
	return "product"
}

type ProductsEnvelope struct {
	TotalCount int       `json:"totalCount"`
	Items      []Product `json:"items"`
}
