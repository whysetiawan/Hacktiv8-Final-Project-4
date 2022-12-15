package repositories

import (
	"final-project-4/httpserver/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	Register(user *models.UserModel) (*models.UserModel, error)
	Login(user *models.UserModel) (*models.UserModel, error)
	GetUsers() (*[]models.UserModel, error)
	GetUser(user *models.UserModel) (*models.UserModel, error)
	UpdateUser(user *models.UserModel) (*models.UserModel, error)
	DeleteUser(user *models.UserModel) (*models.UserModel, error)
	TopUpBalance(userID uint, balance uint) (*models.UserModel, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Register(user *models.UserModel) (*models.UserModel, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetUser(user *models.UserModel) (*models.UserModel, error) {
	err := r.db.Find(user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Login(user *models.UserModel) (*models.UserModel, error) {
	err := r.db.Where("email = ?", user.Email).First(user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetUsers() (*[]models.UserModel, error) {
	var users []models.UserModel
	err := r.db.Preload(clause.Associations).Find(&users).Limit(10).Error

	if err != nil {
		return &users, err
	}

	return &users, nil
}

func (r *userRepository) UpdateUser(user *models.UserModel) (*models.UserModel, error) {
	err := r.db.Model(user).Omit("TransactionHistory").Updates(user).Error

	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *userRepository) DeleteUser(user *models.UserModel) (*models.UserModel, error) {
	err := r.db.Preload(clause.Associations).Delete(user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) TopUpBalance(userID uint, balance uint) (*models.UserModel, error) {
	user := models.UserModel{
		BaseModel: models.BaseModel{
			ID: userID,
		},
	}
	err := r.db.First(&user).Error
	if err != nil {
		return &user, err
	}

	newBalance := user.Balance + int64(balance)
	err = r.db.Model(&user).Update("balance", newBalance).Error
	if err != nil {
		return &user, err
	}
	return &user, nil
}
