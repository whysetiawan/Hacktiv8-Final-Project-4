package models

type UserModel struct {
	BaseModel
	FullName           string                    `gorm:"uniqueIndex" json:"full_name"`
	Email              string                    `gorm:"uniqueIndex" json:"email"`
	Password           string                    `json:"password"`
	Role               string                    `json:"role"`
	Balance            int64                     `json:"balance"`
	TransactionHistory []TransactionHistoryModel `json:"transaction_history" gorm:"foreignKey:UserID;references:ID"`
}

func (UserModel) TableName() string {
	return "public.User"
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
