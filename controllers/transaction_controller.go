package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/api-kasir/models/web"
	"github.com/ihksanghazi/api-kasir/services"
	"github.com/ihksanghazi/api-kasir/utils"
	"gorm.io/gorm"
)

type TransactionController interface {
	CreateTransactionController(c *gin.Context)
}

type TransactionControllerImpl struct {
	service services.TransactionService
}

func NewTransactionController(service services.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		service: service,
	}
}

func (t *TransactionControllerImpl) CreateTransactionController(c *gin.Context) {
	var req web.CreateTransactionWebRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ValidationError(err))
		return
	}

	result, err := t.service.CreateTransactionService(req)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, err.Error())
			return
		default:
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	response := web.Response{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   result,
	}

	c.JSON(http.StatusCreated, response)
}
