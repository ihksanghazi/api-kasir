package controllers

import (
	"net/http"
	"strconv"
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
	GetTransactionController(c *gin.Context)
	ReportTransactionController(c *gin.Context)
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
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "5")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid start date format"})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid end date format"})
		return
	}

	Page, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	Limit, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	result, totalPage, err := t.service.FindTransactionService(startDate, endDate, Page, Limit)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	response := web.Pagination{
		Code:        200,
		Status:      "OK",
		CurrentPage: Page,
		TotalPage:   totalPage,
		Data:        result,
	}

	c.JSON(200, response)
}

func (t *TransactionControllerImpl) GetTransactionController(c *gin.Context) {
	id := c.Param("id")

	result, err := t.service.GetTransactionController(id)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	response := web.Response{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	c.JSON(200, response)
}

func (t *TransactionControllerImpl) ReportTransactionController(c *gin.Context) {
	startDateStr := c.DefaultQuery("startDate", time.Now().Format("2006-01-02"))
	endDateStr := c.DefaultQuery("endDate", time.Now().Add(24*time.Hour).Format("2006-01-02"))

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid start date format"})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid end date format"})
		return
	}

	c.JSON(200, gin.H{
		"start_date": startDate,
		"endDate":    endDate,
	})
}
