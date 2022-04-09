package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ProductId uint `gorm:"primarykey"`
	Title     string
	Sku       string
	Created   time.Time
}

func (Product) TableName() string {
	return "product"
}

func main() {
	fmt.Println("Starting server...")

	dsn := "root:@tcp(127.0.0.1:3306)/sitoo_test_assignment?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.Create(&Product{
		ProductId: 7,
		Title:     "Awesome socks",
		Sku:       "SCK-4511",
		Created:   time.Now(),
	})
	fmt.Println(db)
}
