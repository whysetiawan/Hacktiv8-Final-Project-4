package services

import (
	"final-project-4/httpserver/dto"
	"final-project-4/httpserver/models"
	"final-project-4/httpserver/repositories"
)

type CategoryService interface {
	CreateCategory(dto *dto.CategoryDTO, userID uint) (*models.CategoryModel, error)
	GetCategories(userID uint) (*[]models.CategoryModel, error)
	UpdateCategory(dto *dto.CategoryDTO, categoryID uint, userID uint) (*models.CategoryModel, error)
	DeleteCategory(categoryID uint, userID uint) (*models.CategoryModel, error)
}

type categoryService struct {
	repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) *categoryService {
	return &categoryService{
		categoryRepository,
	}
}

func (s *categoryService) CreateCategory(dto *dto.CategoryDTO, userID uint) (*models.CategoryModel, error) {
	category := models.CategoryModel{
		UserID: userID,
		Type:   dto.Type,
	}

	result, err := s.CategoryRepository.CreateCategory(&category)
	if err != nil {
		return &category, err
	}
	return result, nil
}

func (s *categoryService) GetCategories(userID uint) (*[]models.CategoryModel, error) {
	category := models.CategoryModel{
		UserID: userID,
	}

	result, err := s.CategoryRepository.GetCategories(&category)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *categoryService) UpdateCategory(dto *dto.CategoryDTO, categoryID, userID uint) (*models.CategoryModel, error) {
	category := models.CategoryModel{
		UserID: userID,
		BaseModel: models.BaseModel{
			ID: categoryID,
		},
		Type: dto.Type,
	}

	result, err := s.CategoryRepository.UpdateCategory(&category)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *categoryService) DeleteCategory(categoryID uint, userID uint) (*models.CategoryModel, error) {
	category := models.CategoryModel{
		UserID: userID,
		BaseModel: models.BaseModel{
			ID: categoryID,
		},
	}

	result, err := s.CategoryRepository.DeleteCategory(&category)

	if err != nil {
		return result, err
	}

	return result, nil
}
