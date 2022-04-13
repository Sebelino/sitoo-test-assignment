package main

import (
	"fmt"
	"github.com/Sebelino/sitoo-test-assignment/database"
	"github.com/Sebelino/sitoo-test-assignment/model"
	"github.com/Sebelino/sitoo-test-assignment/routers"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func makeSku() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n := r1.Intn(10000)
	return fmt.Sprintf("SCK-%d", n)
}

func insert(db *gorm.DB) {
	sku := makeSku()
	db.Create(&model.Product{
		Title:   "Awesome socks",
		Sku:     sku,
		Created: time.Now(),
	})
}

func main() {
	fmt.Println("Starting server...")
	db := database.Setup()
	insert(db)

	router := routers.Setup(db)
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
