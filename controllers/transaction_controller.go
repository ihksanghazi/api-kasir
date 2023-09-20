package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/api-kasir/models/web"
	"github.com/ihksanghazi/api-kasir/utils"
)

type TransactionController interface {
	CreateTransaction(c *gin.Context)
}

type TransactionControllerImpl struct{}

func NewTransactionController() TransactionController {
	return &TransactionControllerImpl{}
}

func (t *TransactionControllerImpl) CreateTransaction(c *gin.Context) {
	var req web.CreateTransactionWebRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ValidationError(err))
		return
	}

	c.JSON(200, req)
}
