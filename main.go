package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
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

func makeConnection() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/sitoo_test_assignment?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	return connection
}

func insert(connection *gorm.DB) {
	sku := makeSku()
	connection.Create(&Product{
		Title:   "Awesome socks",
		Sku:     sku,
		Created: time.Now(),
	})
}

func getProducts(c *gin.Context) {
	sampleProducts := []Product{
		{ProductId: 777, Title: "Sample title"},
	}
	c.IndentedJSON(http.StatusOK, sampleProducts)
}

func main() {
	fmt.Println("Starting server...")
	connection := makeConnection()
	insert(connection)

	router := gin.Default()
	router.GET("/api/products", getProducts)
	router.Run("localhost:8080")
}
