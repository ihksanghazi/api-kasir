package web

import "github.com/google/uuid"

type CreateTransactionWebRequest struct {
	Products []CreateTransactionProductWebRequest `json:"products"`
}

type CreateTransactionProductWebRequest struct {
	ProductID uuid.UUID `json:"product_id"`
	Amount    int       `json:"amount"`
}
