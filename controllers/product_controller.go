package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/api-kasir/models/web"
	"github.com/ihksanghazi/api-kasir/services"
	"github.com/ihksanghazi/api-kasir/utils"
)

type ProductController interface {
	CreateProductController(c *gin.Context)
	FindProductController(c *gin.Context)
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

	c.JSON(200, response)
}

func (p *ProductControllerImpl) FindProductController(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "5")
	search := c.DefaultQuery("search", "")

	c.JSON(200, gin.H{
		"page":   page,
		"limit":  limit,
		"search": search,
	})
}
