package services

import (
	"errors"
	"final-project-4/httpserver/dto"
	"final-project-4/httpserver/models"
	"final-project-4/httpserver/repositories"
)

type ProductService interface {
	Create(input dto.InputProduct) (interface{}, map[string]string, error)
	GetProducts() ([]dto.ResponseProducts, map[string]string, error)
	UpdateProduct(param *dto.Product) (*dto.Product, error)
	DeleteProduct(product *dto.Product) (*dto.Product, error)
}

type productService struct {
	productRepo repositories.ProductRepo
}

func NewProductService(pr repositories.ProductRepo) *productService {
	return &productService{pr}
}

func (ps *productService) Create(input dto.InputProduct) (interface{}, map[string]string, error) {
	var res dto.ResponseProductCreate

	if input.Title == "" {
		return nil, map[string]string{
			"en": "you must provide a title",
			"id": "Anda harus memberikan judul",
		}, errors.New("title cannot be empty")
	} else if input.Stock == 0 || input.Stock <= 5 {
		return nil, map[string]string{
			"en": "you must provide a stock or stock can't be less than 5",
			"id": "Anda harus menyediakan stok atau stok tidak boleh kurang dari 5",
		}, errors.New("stock cannot be empty and can't be less than 5")
	} else if input.Price == 0 || input.Price >= 50000000 || input.Price < 0 {
		return nil, map[string]string{
			"en": "You must provide a price and the price cannot be more than 50 million and cannot be less than 0",
			"id": "Anda harus menyediakan harga dan harga tidak boleh lebih dari 50 juta dan tidak boleh kurang dari 0",
		}, errors.New("price cannot be empty")
	}

	data, err := ps.productRepo.Save(models.ProductModel{
		Title:      input.Title,
		Price:      int(input.Price),
		Stock:      int(input.Stock),
		CategoryID: uint(input.CategoryID),
	})
	if err != nil {
		return nil, map[string]string{
			"en": "failed to save product",
			"id": "gagal menginput produk",
		}, err
	}

	res = dto.ResponseProductCreate{
		ID:         int64(data.BaseModel.ID),
		Title:      data.Title,
		Price:      int64(data.Price),
		Stock:      int64(data.Stock),
		CategoryID: int64(data.CategoryID),
		CreatedAt:  data.BaseModel.CreatedAt,
		UpdatedAt:  data.BaseModel.UpdatedAt,
	}

	return res, map[string]string{
		"en": "successfully to save product",
		"id": "berhasil menginput produk",
	}, nil
}

func (ps *productService) GetProducts() ([]dto.ResponseProducts, map[string]string, error) {
	// var res dto.ResponseProducts

	data, err := ps.productRepo.GetProducts()
	if err != nil {
		return data, map[string]string{
			"en": "failed to get data products",
			"id": "gagal mengambil data produk",
		}, err
	}

	return data, map[string]string{
		"en": "successfully get data products",
		"id": "berhasil menampilkan data produk",
	}, nil

}

func (ps *productService) UpdateProduct(param *dto.Product) (*dto.Product, error) {

	product := dto.Product{
		Title:      param.Title,
		Price:      param.Price,
		Stock:      param.Stock,
		CategoryID: param.CategoryID,
		ID:         param.ID,
	}

	result, err := ps.productRepo.UpdateProduct(&product)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (ps *productService) DeleteProduct(product *dto.Product) (*dto.Product, error) {
	product, err := ps.productRepo.DeleteProduct(&dto.Product{
		ID: product.ID,
	})
	if err != nil {
		return product, err
	}

	return product, nil

}
