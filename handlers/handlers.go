package handlers

import (
	"fmt"
	"github.com/Sebelino/sitoo-test-assignment/database"
	"github.com/Sebelino/sitoo-test-assignment/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type ApiEnv struct {
	Db *gorm.DB
}

func (e *ApiEnv) GetProducts(c *gin.Context) {
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
	}, e.Db)

	response := model.ProductsEnvelope{
		TotalCount: len(products),
		Items:      products,
	}
	c.IndentedJSON(http.StatusOK, response)
}

func (e *ApiEnv) PostProduct(context *gin.Context) {
	fmt.Println("POSTing product")
	product := model.Product{}
	err := context.BindJSON(&product)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	product.Created = time.Now()
	database.CreateProduct(e.Db, product)
}
