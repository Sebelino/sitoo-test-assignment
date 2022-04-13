package database

import (
	"fmt"
	"github.com/Sebelino/sitoo-test-assignment/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Setup() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/sitoo_test_assignment?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	return db
}

type ProductFilter struct {
	Start   int
	Num     int
	Sku     string
	Barcode string
}

func GetProducts(filter ProductFilter, db *gorm.DB) []model.Product {
	var queryFilter = make(map[string]string)
	if filter.Sku != "" {
		queryFilter["sku"] = filter.Sku
	}
	if filter.Barcode != "" {
		queryFilter["barcode"] = filter.Barcode
	}
	var products []model.Product
	if filter.Num > 0 {
		db.Where(queryFilter).Offset(filter.Start).Limit(filter.Num).Find(&products)
	}
	return products
}

func CreateProduct(db *gorm.DB, product model.Product) {
	fmt.Printf("Insert %v\n", product)
	db.Create(&product)
}
