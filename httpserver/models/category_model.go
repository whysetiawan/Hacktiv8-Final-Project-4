package models

type CategoryModel struct {
	BaseModel
	Type              string         `json:"type"`
	SoldProductAmount int            `json:"sold_product_amount"`
	Products          []ProductModel `json:"products" gorm:"foreignKey:CategoryID;references:ID"`
}

func (CategoryModel) TableName() string {
	return "public.Category"
}
