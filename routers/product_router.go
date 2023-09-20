package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/api-kasir/controllers"
)

func ProductRoute(r *gin.RouterGroup) {
	productController := controllers.NewProductController()
	r.POST("/", productController.CreateProductController)
}
