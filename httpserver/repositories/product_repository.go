package repositories

import (
	"final-project-4/httpserver/dto"
	"final-project-4/httpserver/models"

	"gorm.io/gorm"
)

type ProductRepo interface {
	Save(product models.ProductModel) (models.ProductModel, error)
	GetProducts() ([]dto.ResponseProducts, error)
	GetProduct(product *models.ProductModel) (*models.ProductModel, error)
	UpdateProduct(product *dto.Product) (*dto.Product, error)
	DeleteProduct(product *dto.Product) (*dto.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *productRepo {
	return &productRepo{db: db}
}

func (pr *productRepo) Save(product models.ProductModel) (models.ProductModel, error) {
	err := pr.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (pr *productRepo) GetProducts() ([]dto.ResponseProducts, error) {
	var product []dto.ResponseProducts
	// err := r.db.Preload(clause.Associations).Find(&product).Limit(10).Error

	err := pr.db.Raw(`select id, title, price, stock, category_id, created_at from public."Product"`).Find(&product).Error
	// err := r.db.Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (pr *productRepo) UpdateProduct(product *dto.Product) (*dto.Product, error) {
	// var product dto.ResponseProducts

	err := pr.db.Model(&product).Updates(product).Where("product_id = ?", product.ID)

	if err != nil {
		return product, err.Error
	}
	return product, nil
}

func (pr *productRepo) DeleteProduct(product *dto.Product) (*dto.Product, error) {
	err := pr.db.Model(product).Delete(product).Error

	if err != nil {
		return product, err
	}
	return product, nil

}

func (pr *productRepo) GetProduct(product *models.ProductModel) (*models.ProductModel, error) {
	err := pr.db.Find(product).Error
	if err != nil {
		return product, err
	}
	return product, nil

}
