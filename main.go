package main

import (
	"fmt"
	"github.com/Sebelino/sitoo-test-assignment/database"
	"github.com/Sebelino/sitoo-test-assignment/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"strconv"
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

		products := database.GetProducts(database.ProductFilter{
			Start:   start,
			Num:     num,
			Sku:     sku,
			Barcode: barcode,
		}, connection)

		response := model.ProductsEnvelope{
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
	connection := database.Setup()
	insert(connection)

	router := setupRouter(connection)
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
