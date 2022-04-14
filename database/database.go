package database

import (
	"github.com/Sebelino/sitoo-test-assignment/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Setup() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/sitoo_test_assignment?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	if err := db.AutoMigrate(&model.Product{}); err != nil {
		log.Fatalf("Failed to perform database migration of products: %v", err)
	}
	if err := db.AutoMigrate(&model.ProductBarcode{}); err != nil {
		log.Fatalf("Failed to perform database migration of product barcodes: %v", err)
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
		db.Preload("Barcodes").Where(queryFilter).Offset(filter.Start).Limit(filter.Num).Find(&products)
	}
	return products
}

func CreateProduct(db *gorm.DB, product *model.Product) error {
	if err := db.Create(product).Error; err != nil {
		return err
	}
	return nil
}
