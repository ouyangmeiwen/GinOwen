package repositories

import (
	"GINOWEN/global"
	"GINOWEN/models"

	"gorm.io/gorm"
)

type GormUserRepository struct {
}

func (repo *GormUserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := global.GVA_DB.Find(&users).Error
	return users, err
}

func (repo *GormUserRepository) CreateUser(user *models.User) error {
	return global.GVA_DB.Create(user).Error
}

func (repo *GormUserRepository) UpdateUser(userID uint, updates map[string]interface{}) error {
	return global.GVA_DB.Model(&models.User{}).Where("id = ?", userID).Updates(updates).Error
}

func (repo *GormUserRepository) ExecuteInTransaction(txFunc func(*gorm.DB) error) error {
	tx := global.GVA_DB.Begin()
	if err := txFunc(tx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
