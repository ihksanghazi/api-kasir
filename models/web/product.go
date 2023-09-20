package web

import "github.com/google/uuid"

type CreateProductWebRequest struct {
	ProductName   string  `json:"product_name" binding:"required"`
	PurchasePrice float64 `json:"purchase_price" binding:"required"`
	SellingPrice  float64 `json:"selling_price" binding:"required"`
	Stock         int     `json:"stock" binding:"required"`
}

type FindProductWebResponse struct {
	ID            uuid.UUID `json:"id"`
	ProductName   string    `json:"product_name"`
	PurchasePrice float64   `json:"purchase_price"`
	SellingPrice  float64   `json:"selling_price"`
	Stock         int       `json:"stock"`
}
