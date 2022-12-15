package models

type CategoryModel struct {
	BaseModel
	UserID            uint           `json:"user_id"`
	Type              string         `json:"type"`
	SoldProductAmount int            `json:"sold_product_amount"`
	Products          []ProductModel `json:"products" gorm:"foreignKey:CategoryID;references:ID"`
}

func (CategoryModel) TableName() string {
	return "public.Category"
}
