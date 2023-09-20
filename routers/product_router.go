package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/api-kasir/controllers"
	"github.com/ihksanghazi/api-kasir/database"
	"github.com/ihksanghazi/api-kasir/services"
)

func ProductRoute(r *gin.RouterGroup) {
	var ctx context.Context
	service := services.NewProductService(database.DB, ctx)
	controller := controllers.NewProductController(service)
	r.POST("/", controller.CreateProductController)
	r.GET("/", controller.FindProductController)
	r.PUT("/:id", controller.UpdateProductController)
	r.DELETE("/:id", controller.DeleteProductController)
}
