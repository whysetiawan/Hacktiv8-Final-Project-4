package models

type ProductModel struct {
	BaseModel
	Title              string                    `json:"title"`
	Price              int                       `json:"price"`
	Stock              int                       `json:"stock"`
	CategoryID         uint                      `json:"category_id"`
	TransactionHistory []TransactionHistoryModel `json:"transaction_history" gorm:"foreignKey:ProductID;references:ID"`
}

func (ProductModel) TableName() string {
	return "public.Product"
}
