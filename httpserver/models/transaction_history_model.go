package models

type TransactionHistoryModel struct {
	BaseModel
	Quantity   int  `json:"quantity"`
	TotalPrice int  `json:"total_price"`
	UserID     uint `json:"user_id"`
	ProductID  uint `json:"product_id"`
}

func (TransactionHistoryModel) TableName() string {
	return "public.TransactionHistory"
}
