package services

import (
	"context"
	"errors"

	"github.com/ihksanghazi/api-kasir/models/domain"
	"github.com/ihksanghazi/api-kasir/models/web"
	"gorm.io/gorm"
)

type TransactionService interface {
	CreateTransactionService(req web.CreateTransactionWebRequest) (web.CreateTransactionWebResponse, error)
}

type TransactionServiceImpl struct {
	db  *gorm.DB
	ctx context.Context
}

func NewTransactionService(db *gorm.DB, ctx context.Context) TransactionService {
	return &TransactionServiceImpl{
		db:  db,
		ctx: ctx,
	}
}

func (t *TransactionServiceImpl) CreateTransactionService(req web.CreateTransactionWebRequest) (web.CreateTransactionWebResponse, error) {
	var response web.CreateTransactionWebResponse
	errTransaction := t.db.Transaction(func(tx *gorm.DB) error {
		var totalTransaction float64
		// create transaction
		var transaction domain.Transaction
		transaction.Total = totalTransaction
		if err := tx.Model(transaction).WithContext(t.ctx).Create(&transaction).Error; err != nil {
			return err
		}
		//transaction details
		for _, item := range req.Products {
			// cek amount product
			var product domain.Product
			if err := tx.Model(product).WithContext(t.ctx).Where("id = ?", item.ProductID).First(&product).Error; err != nil {
				return err
			}
			if product.Stock-item.Amount < 0 {
				return errors.New(product.ProductName + "is not available")
			}
			// kalkulation
			totalTransaction += float64(item.Amount) * product.SellingPrice
			// create
			var transactionDetail domain.TransactionDetail
			transactionDetail.TransactionID = transaction.ID
			transactionDetail.ProductID = item.ProductID
			transactionDetail.Amount = item.Amount
			if err := tx.Model(transactionDetail).WithContext(t.ctx).Create(&transactionDetail).Error; err != nil {
				return err
			}
			// update product
			if err := tx.Model(product).WithContext(t.ctx).Where("id = ?", item.ProductID).Update("stock", gorm.Expr("stock - ?", item.Amount)).Error; err != nil {
				return err
			}
		}
		// update transaction
		if err := tx.Model(transaction).WithContext(t.ctx).Update("total", totalTransaction).Error; err != nil {
			return err
		}

		if err := tx.Model(transaction).WithContext(t.ctx).Where("id = ?", transaction.ID).Preload("TransactionDetails.Products").First(&response).Error; err != nil {
			return err
		}
		return nil
	})

	return response, errTransaction
}
