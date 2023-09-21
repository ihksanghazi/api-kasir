package web

import (
	"time"

	"github.com/google/uuid"
)

type CreateTransactionWebRequest struct {
	Products []CreateTransactionProductWebRequest `json:"products"`
}

type CreateTransactionProductWebRequest struct {
	ProductID uuid.UUID `json:"product_id"`
	Amount    int       `json:"amount"`
}

type CreateTransactionWebResponse struct {
	ID        uuid.UUID `json:"id"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Association
	TransactionDetails []CreateTransactionDetailWebResponse `gorm:"foreignKey:TransactionID" json:"transaction_details"`
}

func (c *CreateTransactionWebResponse) TableName() string {
	return "transactions"
}

type CreateTransactionDetailWebResponse struct {
	ID            uuid.UUID `json:"-"`
	TransactionID uuid.UUID `json:"-"`
	ProductID     uuid.UUID `json:"-"`
	Amount        int       `json:"amount"`
	// Association
	Products CreateTransactionDetailProductWebResponse `gorm:"foreignKey:ProductID" json:"products"`
}

func (c *CreateTransactionDetailWebResponse) TableName() string {
	return "transaction_details"
}

type CreateTransactionDetailProductWebResponse struct {
	ID           uuid.UUID `json:"-"`
	ProductName  string    `json:"product_name"`
	SellingPrice float64   `json:"selling_price"`
}

func (c *CreateTransactionDetailProductWebResponse) TableName() string {
	return "products"
}
