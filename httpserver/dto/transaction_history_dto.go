package dto

type CreateTransactionDto struct {
	ProductID int  `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
