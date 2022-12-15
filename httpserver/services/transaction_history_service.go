package services

import (
	"errors"
	"final-project-4/httpserver/dto"
	"final-project-4/httpserver/models"
	"final-project-4/httpserver/repositories"
)

type TransactionHistoryService interface {
	CreateTransaction(dto *dto.CreateTransactionDto, user *models.UserModel) (*models.TransactionHistoryModel, error)
	GetMyTransactions(user *models.UserModel) (*[]models.TransactionHistoryModel, error)
	GetUserTransactions(user *models.UserModel) (*[]models.TransactionHistoryModel, error)
}

type transactionHistoryService struct {
	repositories.TransactionHistoryRepository
	repositories.ProductRepo
	repositories.UserRepository
	repositories.CategoryRepository
}

func NewTransactionHistoryService(
	transactionRepository repositories.TransactionHistoryRepository,
	productRepository repositories.ProductRepo,
	userRepository repositories.UserRepository,
	categoryRepository repositories.CategoryRepository,
) *transactionHistoryService {
	return &transactionHistoryService{
		transactionRepository,
		productRepository,
		userRepository,
		categoryRepository,
	}
}

func (s *transactionHistoryService) CreateTransaction(trxDto *dto.CreateTransactionDto, user *models.UserModel) (*models.TransactionHistoryModel, error) {
	transaction := &models.TransactionHistoryModel{
		ProductID: uint(trxDto.ProductID),
		Quantity:  int(trxDto.Quantity),
		UserID:    user.ID,
	}

	product := &models.ProductModel{
		BaseModel: models.BaseModel{
			ID: uint(trxDto.ProductID),
		},
	}

	// Check if product exist or has stocks, otherwise it will return error
	product, err := s.ProductRepo.GetProduct(product)
	if err != nil {
		return transaction, err
	}

	category := &models.CategoryModel{
		BaseModel: models.BaseModel{
			ID: product.CategoryID,
		},
	}

	category, err = s.CategoryRepository.GetCategory(category)

	if err != nil {
		return transaction, err
	}

	if product.Stock < transaction.Quantity {
		return transaction, errors.New("stock isn't enough")
	}

	transaction.TotalPrice = product.Price * int(trxDto.Quantity)

	// Check if user has enough balance to purchase
	user, err = s.UserRepository.GetUser(user)

	if err != nil {
		return transaction, err
	}

	if user.Balance < int64(transaction.TotalPrice) {
		return transaction, errors.New("insuffienct balance")
	}
	user.Balance = user.Balance - int64(transaction.TotalPrice)
	product.Stock = product.Stock - transaction.Quantity
	category.SoldProductAmount = category.SoldProductAmount + transaction.Quantity

	transaction, err = s.TransactionHistoryRepository.CreateTransaction(transaction)

	if err != nil {
		return transaction, err
	}
	s.ProductRepo.UpdateProduct(&dto.Product{
		ID:         int64(product.ID),
		Title:      product.Title,
		Price:      int64(product.Price),
		Stock:      int64(product.Stock),
		CategoryID: int64(product.CategoryID),
		CreatedAt:  product.CreatedAt,
	})
	s.UserRepository.UpdateUser(user)
	s.CategoryRepository.UpdateCategory(category)

	return transaction, nil

}

func (s *transactionHistoryService) GetUserTransactions(user *models.UserModel) (*[]models.TransactionHistoryModel, error) {
	return s.TransactionHistoryRepository.GetUserTransactions(user)
}

func (s *transactionHistoryService) GetMyTransactions(user *models.UserModel) (*[]models.TransactionHistoryModel, error) {
	return s.TransactionHistoryRepository.GetMyTransactions(user)
}
