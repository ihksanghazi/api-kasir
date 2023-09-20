package services

import (
	"context"
	"errors"

	"github.com/ihksanghazi/api-kasir/models/domain"
	"github.com/ihksanghazi/api-kasir/models/web"
	"gorm.io/gorm"
)

type ProductService interface {
	CreateProductService(req web.CreateProductWebRequest) (web.CreateProductWebRequest, error)
}

type ProductServiceImpl struct {
	db  *gorm.DB
	ctx context.Context
}

func NewProductService(db *gorm.DB, ctx context.Context) ProductService {
	return &ProductServiceImpl{
		db:  db,
		ctx: ctx,
	}
}

func (p *ProductServiceImpl) CreateProductService(req web.CreateProductWebRequest) (web.CreateProductWebRequest, error) {
	// parsing to model
	var model domain.Product
	model.ProductName = req.ProductName
	model.PurchasePrice = req.PurchasePrice
	model.SellingPrice = req.SellingPrice
	model.Stock = req.Stock
	err := p.db.Transaction(func(tx *gorm.DB) error {
		// find product
		if err := tx.Model(&model).WithContext(p.ctx).Where("product_name = ?", req.ProductName).First(&model).Error; err == nil {
			return errors.New("product already exists")
		} else if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		// create product
		if err := tx.Model(&model).WithContext(p.ctx).Create(&model).Error; err != nil {
			return err
		}
		return nil
	})

	return req, err
}
