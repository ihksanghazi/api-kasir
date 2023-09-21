package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/api-kasir/controllers"
	"github.com/ihksanghazi/api-kasir/database"
	"github.com/ihksanghazi/api-kasir/services"
)

func TransactionRouter(r *gin.RouterGroup) {

	var ctx context.Context
	service := services.NewTransactionService(database.DB, ctx)

	controller := controllers.NewTransactionController(service)

	r.POST("/", controller.CreateTransactionController)
	r.GET("/", controller.FindTransactionController)
	r.GET("/:id", controller.GetTransactionController)
	r.GET("/report", controller.ReportTransactionController)
}
