package routers

import (
	"github.com/Sebelino/sitoo-test-assignment/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	api := handlers.ApiEnv{
		Db: db,
	}
	router := gin.Default()
	router.GET("/api/products", api.GetProducts)
	router.POST("/api/products", api.PostProduct)
	return router
}
