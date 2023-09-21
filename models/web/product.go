package web

import (
	"time"

	"github.com/google/uuid"
)

type CreateProductWebRequest struct {
	ProductName   string  `json:"product_name" binding:"required"`
	PurchasePrice float64 `json:"purchase_price" binding:"required"`
	SellingPrice  float64 `json:"selling_price" binding:"required"`
	Stock         int     `json:"stock" binding:"required"`
}
type UpdateProductWebRequest struct {
	ProductName   string  `json:"product_name"`
	PurchasePrice float64 `json:"purchase_price"`
	SellingPrice  float64 `json:"selling_price"`
	Stock         int     `json:"stock"`
}

type FindProductWebResponse struct {
	ID            uuid.UUID `json:"id"`
	ProductName   string    `json:"product_name"`
	PurchasePrice float64   `json:"purchase_price"`
	SellingPrice  float64   `json:"selling_price"`
	Stock         int       `json:"stock"`
}

type GetProductWebResponse struct {
	ID            uuid.UUID `json:"id"`
	ProductName   string    `json:"product_name"`
	PurchasePrice float64   `json:"purchase_price"`
	SellingPrice  float64   `json:"selling_price"`
	Stock         int       `json:"stock"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	// Association
	Transactions []GetProductTransactionWebResponse `gorm:"many2many:transaction_details;foreignKey:ID;joinForeignKey:ProductID;References:ID;joinReferences:TransactionID"`
}

func (g *GetProductWebResponse) TableName() string {
	return "products"
}

type GetProductTransactionWebResponse struct {
	ID        uuid.UUID `json:"id"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (g *GetProductTransactionWebResponse) TableName() string {
	return "transactions"
}
