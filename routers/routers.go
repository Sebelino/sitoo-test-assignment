package routers

import (
	"github.com/Sebelino/sitoo-test-assignment/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(connection *gorm.DB) *gin.Engine {
	api := handlers.ApiEnv{
		Connection: connection,
	}
	router := gin.Default()
	router.GET("/api/products", api.GetProducts)
	return router
}
