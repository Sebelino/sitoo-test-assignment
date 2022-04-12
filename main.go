package main

import (
	"fmt"
	"github.com/Sebelino/sitoo-test-assignment/database"
	"github.com/Sebelino/sitoo-test-assignment/handlers"
	"github.com/Sebelino/sitoo-test-assignment/model"
	"github.com/gin-gonic/gin"
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

func insert(connection *gorm.DB) {
	sku := makeSku()
	connection.Create(&model.Product{
		Title:   "Awesome socks",
		Sku:     sku,
		Created: time.Now(),
	})
}

func setupRouter(connection *gorm.DB) *gin.Engine {
	api := handlers.ApiEnv{
		Connection: connection,
	}
	router := gin.Default()
	router.GET("/api/products", api.GetProducts)
	return router
}

func main() {
	fmt.Println("Starting server...")
	connection := database.Setup()
	insert(connection)

	router := setupRouter(connection)
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
