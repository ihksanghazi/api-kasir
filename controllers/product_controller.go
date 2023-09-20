package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ihksanghazi/api-kasir/models/web"
	"github.com/ihksanghazi/api-kasir/utils"
)

type ProductController interface {
	CreateProductController(c *gin.Context)
}

type ProductControllerImpl struct {
}

func NewProductController() ProductController {
	return &ProductControllerImpl{}
}

func (p *ProductControllerImpl) CreateProductController(c *gin.Context) {
	var req web.CreateProductWebRequest
	if errBind := c.ShouldBindJSON(&req); errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": utils.ValidationError(errBind)})
		return
	}

	response := web.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   req,
	}

	c.JSON(200, response)
}
