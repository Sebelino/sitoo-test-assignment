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
	ProductId   uint      `json:"productId" gorm:"primarykey"`
	Title       string    `json:"title"`
	Sku         string    `json:"sku"`
	Created     time.Time `json:"created"`
	Description *string   `json:"description"`
}

type ProductsEnvelope struct {
	TotalCount int       `json:"totalCount"`
	Items      []Product `json:"items"`
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

func getProductsHandler(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		start := c.DefaultQuery("start", "0")
		num := c.DefaultQuery("num", "10")
		sku := c.DefaultQuery("sku", "")
		barcode := c.DefaultQuery("barcode", "")
		fmt.Printf("Query parameter: start=%s\n", start)
		fmt.Printf("Query parameter: num=%s\n", num)
		fmt.Printf("Query parameter: sku=%s\n", sku)
		fmt.Printf("Query parameter: barcode=%s\n", barcode)
		var products []Product
		db.Find(&products)
		response := ProductsEnvelope{
			TotalCount: len(products),
			Items:      products,
		}
		c.IndentedJSON(http.StatusOK, response)
	}
}

func main() {
	fmt.Println("Starting server...")
	connection := makeConnection()
	insert(connection)

	router := gin.Default()
	router.GET("/api/products", getProductsHandler(connection))
	router.Run(":8080")
}
