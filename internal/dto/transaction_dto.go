package dto

import "time"

type TransactionRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type TransactionResponse struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}
