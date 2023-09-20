package web

type CreateProductWebRequest struct {
	ProductName   string  `json:"product_name" binding:"required"`
	PurchasePrice float64 `json:"purchase_price" binding:"required"`
	SellingPrice  float64 `json:"selling_price" binding:"required"`
	Stock         int     `json:"stock" binding:"required"`
}
