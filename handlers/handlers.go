package handlers

import (
	"fmt"
	"github.com/Sebelino/sitoo-test-assignment/database"
	"github.com/Sebelino/sitoo-test-assignment/model"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
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

type HttpError struct {
	ErrorCode int    `json:"errorCode"`
	ErrorText string `json:"errorText"`
}

func (e *ApiEnv) PostProduct(context *gin.Context) {
	product := model.Product{}
	if err := context.BindJSON(&product); err != nil {
		fmt.Println(err)
		httpError := HttpError{
			ErrorCode: 1,
			ErrorText: "Could not process response",
		}
		context.IndentedJSON(http.StatusBadRequest, httpError)
		return
	}
	product.Created = time.Now()
	if err := database.CreateProduct(e.Db, product); err != nil {
		me, _ := err.(*mysql.MySQLError)
		if me.Number == mysqlDuplicateEntry {
			httpError := HttpError{
				ErrorCode: dbToCustomErrorCode(me.Number),
				ErrorText: "The supplied product already exists",
			}
			context.IndentedJSON(http.StatusBadRequest, httpError)
			return
		}
		httpError := HttpError{
			ErrorCode: 2,
			ErrorText: "Could not insert product in database",
		}
		context.IndentedJSON(http.StatusInternalServerError, httpError)
		return
	}
	context.Status(http.StatusCreated)
}
