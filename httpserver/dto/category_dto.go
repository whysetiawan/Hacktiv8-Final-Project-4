package dto

type CategoryDTO struct {
	Type string `json:"type" binding:"required"`
}
