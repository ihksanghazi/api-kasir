package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/api-kasir/database"
	"github.com/ihksanghazi/api-kasir/routers"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file : ", err.Error())
		return
	}

	//connect database
	database.Connect()

	// migrate
	// database.DB.AutoMigrate(domain.Product{}, domain.Transaction{}, domain.TransactionDetail{})

	routers.ProductRoute(r.Group("/api/product"))
	routers.TransactionRouter(r.Group("/api/transaction"))

	r.Run(":5000")
}
