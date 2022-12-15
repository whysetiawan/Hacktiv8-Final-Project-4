package dto

import (
	"time"
)

type InputProduct struct {
	Title      string `json:"title"`
	Price      int64  `json:"price"`
	Stock      int64  `json:"stock"`
	CategoryID int64  `json:"category_id"`
}

type ResponseProductCreate struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Price      int64     `json:"price"`
	Stock      int64     `json:"stock"`
	CategoryID int64     `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ResponseProducts struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Price      int64     `json:"price"`
	Stock      int64     `json:"stock"`
	CategoryID int64     `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type Product struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Price      int64     `json:"price"`
	Stock      int64     `json:"stock"`
	CategoryID int64     `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func (Product) TableName() string {
	return "public.Product"
}
