package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/api-kasir/models/web"
	"github.com/ihksanghazi/api-kasir/services"
	"github.com/ihksanghazi/api-kasir/utils"
	"gorm.io/gorm"
)

type ProductController interface {
	CreateProductController(c *gin.Context)
	FindProductController(c *gin.Context)
	UpdateProductController(c *gin.Context)
	DeleteProductController(c *gin.Context)
}

type ProductControllerImpl struct {
	service services.ProductService
}

func NewProductController(service services.ProductService) ProductController {
	return &ProductControllerImpl{
		service: service,
	}
}

func (p *ProductControllerImpl) CreateProductController(c *gin.Context) {
	var req web.CreateProductWebRequest
	if errBind := c.ShouldBindJSON(&req); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": utils.ValidationError(errBind)})
		return
	}

	result, err := p.service.CreateProductService(req)
	if err != nil {
		switch err.Error() {
		case "product already exists":
			c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	response := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}

	c.JSON(http.StatusCreated, response)
}

func (p *ProductControllerImpl) FindProductController(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "5")
	search := c.DefaultQuery("search", "")

	Page, _ := strconv.Atoi(page)
	Limit, _ := strconv.Atoi(limit)

	result, totalPage, err := p.service.FindProductService(search, Page, Limit)
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

	pagination := web.Pagination{
		Code:        200,
		Status:      "OK",
		CurrentPage: Page,
		TotalPage:   totalPage,
		Data:        result,
	}

	c.JSON(200, pagination)
}

func (p *ProductControllerImpl) UpdateProductController(c *gin.Context) {
	id := c.Param("id")

	var req web.UpdateProductWebRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := p.service.UpdateProductService(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response := web.Response{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	c.JSON(200, response)
}

func (p *ProductControllerImpl) DeleteProductController(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{"message": id})
}
