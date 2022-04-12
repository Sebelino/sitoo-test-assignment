package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MakeConnection() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/sitoo_test_assignment?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	return connection
}
