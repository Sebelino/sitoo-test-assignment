package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
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

func makeSku() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n := r1.Intn(10000)
	return fmt.Sprintf("SCK-%d", n)
}

func main() {
	fmt.Println("Starting server...")

	dsn := "root:@tcp(127.0.0.1:3306)/sitoo_test_assignment?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	sku := makeSku()
	db.Create(&Product{
		Title:   "Awesome socks",
		Sku:     sku,
		Created: time.Now(),
	})
	fmt.Println(db)
}
