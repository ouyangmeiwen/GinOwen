package repositories

import (
	"GINOWEN/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(userID uint, updates map[string]interface{}) error
	ExecuteInTransaction(txFunc func(*gorm.DB) error) error
}

type GormUserRepository struct {
	DB *gorm.DB
}

func (repo *GormUserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := repo.DB.Find(&users).Error
	return users, err
}

func (repo *GormUserRepository) CreateUser(user *models.User) error {
	return repo.DB.Create(user).Error
}

func (repo *GormUserRepository) UpdateUser(userID uint, updates map[string]interface{}) error {
	return repo.DB.Model(&models.User{}).Where("id = ?", userID).Updates(updates).Error
}

func (repo *GormUserRepository) ExecuteInTransaction(txFunc func(*gorm.DB) error) error {
	tx := repo.DB.Begin()
	if err := txFunc(tx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
