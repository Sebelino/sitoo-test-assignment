package main

import (
	"fmt"
	"github.com/Sebelino/sitoo-test-assignment/database"
	"github.com/Sebelino/sitoo-test-assignment/routers"
	"time"
)

func main() {
	fmt.Println("Starting server...")
	time.Sleep(10 * time.Second)
	db := database.Setup()

	router := routers.Setup(db)
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
