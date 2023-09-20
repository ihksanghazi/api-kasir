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
	FindProductService(search string, page int, limit int) (result []web.FindProductWebResponse, totalPage int64, err error)
	UpdateProductService(id string, req web.UpdateProductWebRequest) (web.UpdateProductWebRequest, error)
	DeleteProductService(id string) error
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

func (p *ProductServiceImpl) FindProductService(search string, page int, limit int) (result []web.FindProductWebResponse, totalPage int64, err error) {
	var model domain.Product
	var response []web.FindProductWebResponse
	//pagination
	var total int64
	offset := (page - 1) * limit
	// getall user by page
	Err := p.db.Model(model).WithContext(p.ctx).Where("product_name ILIKE ?", "%"+search+"%").Count(&total).Offset(offset).Limit(limit).Find(&response).Error

	TotalPage := (total + int64(limit) - 1) / int64(limit)

	return response, TotalPage, Err
}

func (p *ProductServiceImpl) UpdateProductService(id string, req web.UpdateProductWebRequest) (web.UpdateProductWebRequest, error) {
	// parsing to model
	var model domain.Product
	model.ProductName = req.ProductName
	model.PurchasePrice = req.PurchasePrice
	model.SellingPrice = req.SellingPrice
	model.Stock = req.Stock

	err := p.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(model).WithContext(p.ctx).Where("id = ?", id).Updates(&model).Find(&req).Error; err != nil {
			return err
		}
		return nil
	})

	return req, err
}

func (p *ProductServiceImpl) DeleteProductService(id string) error {
	var model domain.Product
	err := p.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(model).WithContext(p.ctx).Where("id = ?", id).Delete(&model).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
