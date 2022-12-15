package repositories

import (
	"final-project-4/httpserver/models"

	"gorm.io/gorm"
)

type TransactionHistoryRepository interface {
	CreateTransaction(transaction *models.TransactionHistoryModel) (*models.TransactionHistoryModel, error)
	GetUserTransactions(user *models.UserModel) (*[]models.TransactionHistoryModel, error)
	GetMyTransactions(user *models.UserModel) (*[]models.TransactionHistoryModel, error)
}

type transactionHistoryRepository struct {
	db *gorm.DB
}

func NewTransactionHistoryRepository(db *gorm.DB) *transactionHistoryRepository {
	return &transactionHistoryRepository{
		db,
	}
}

func (r *transactionHistoryRepository) CreateTransaction(transaction *models.TransactionHistoryModel) (*models.TransactionHistoryModel, error) {
	err := r.db.Create(transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *transactionHistoryRepository) GetUserTransactions(user *models.UserModel) (*[]models.TransactionHistoryModel, error) {
	var transactions []models.TransactionHistoryModel
	err := r.db.Order("created_at desc").Find(&transactions).Error
	if err != nil {
		return &transactions, err
	}
	return &transactions, nil
}

func (r *transactionHistoryRepository) GetMyTransactions(user *models.UserModel) (*[]models.TransactionHistoryModel, error) {
	var transactions []models.TransactionHistoryModel
	err := r.db.Where("user_id = ?", user.ID).Order("created_at desc").Find(&transactions).Error
	if err != nil {
		return &transactions, err
	}
	return &transactions, nil
}
