package main

import (
	"fmt"
	"github.com/Sebelino/sitoo-test-assignment/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"strconv"
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

func insert(connection *gorm.DB) {
	sku := makeSku()
	connection.Create(&Product{
		Title:   "Awesome socks",
		Sku:     sku,
		Created: time.Now(),
	})
}

func getProductsHandler(connection *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		startString := c.DefaultQuery("start", "0")
		numString := c.DefaultQuery("num", "10")
		sku := c.DefaultQuery("sku", "")
		barcode := c.DefaultQuery("barcode", "")

		start, errStart := strconv.Atoi(startString)
		if errStart != nil {
			panic(fmt.Sprintf("Could not parse query parameter \"start\" into integer: %s", startString))
		}
		num, errNum := strconv.Atoi(numString)
		if errNum != nil {
			panic(fmt.Sprintf("Could not parse query parameter \"num\" into integer: %s", numString))
		}

		var filter = make(map[string]string)
		if sku != "" {
			filter["sku"] = sku
		}
		if barcode != "" {
			filter["barcode"] = barcode
		}
		var products []Product
		if num > 0 {
			connection.Where(filter).Offset(start).Limit(num).Find(&products)
		}
		response := ProductsEnvelope{
			TotalCount: len(products),
			Items:      products,
		}
		c.IndentedJSON(http.StatusOK, response)
	}
}

func setupRouter(connection *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/api/products", getProductsHandler(connection))
	return router
}

func main() {
	fmt.Println("Starting server...")
	connection := database.MakeConnection()
	insert(connection)

	router := setupRouter(connection)
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
