package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/api-kasir/controllers"
)

func TransactionRouter(r *gin.RouterGroup) {

	controller := controllers.NewTransactionController()

	r.POST("/", controller.CreateTransaction)
}
