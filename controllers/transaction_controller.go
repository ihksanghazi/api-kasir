package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/api-kasir/models/web"
	"github.com/ihksanghazi/api-kasir/services"
	"github.com/ihksanghazi/api-kasir/utils"
	"gorm.io/gorm"
)

type TransactionController interface {
	CreateTransactionController(c *gin.Context)
	FindTransactionController(c *gin.Context)
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

func (t *TransactionControllerImpl) FindTransactionController(c *gin.Context) {
	startDateStr := c.DefaultQuery("startDate", time.Now().Format("2006-01-02"))
	endDateStr := c.DefaultQuery("endDate", time.Now().Add(24*time.Hour).Format("2006-01-02"))

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid start date format"})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid end date format"})
		return
	}

	c.JSON(200, gin.H{
		"start_date": startDate,
		"end_date":   endDate,
	})
}
