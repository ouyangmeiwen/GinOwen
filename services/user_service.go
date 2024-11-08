package services

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
)

type UserService struct {
}

func (s *UserService) GetAllUsers() ([]response.UserDto, error) {
	var users []models.TestUser
	err := global.OWEN_DB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	var userResponses []response.UserDto
	for _, user := range users {
		userResponses = append(userResponses, response.UserDto{
			ID:   user.ID,
			Name: user.Name,
		})
	}
	return userResponses, nil
}

func (s *UserService) CreateUser(userRequest request.UserRequest) (response.UserDto, error) {
	user := models.TestUser{Name: userRequest.Name}
	if err := global.OWEN_DB.Create(user).Error; err != nil {
		return response.UserDto{}, err
	}
	return response.UserDto{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (s *UserService) UpdateUserAndCreateOrder(userID uint, order *models.TestUser) error {
	// 开启事务
	tx := global.OWEN_DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	// 使用事务进行更新操作
	err := tx.Model(&models.TestUser{}).Where("id = ?", userID).Updates(map[string]interface{}{"status": "updated"}).Error
	if err != nil {
		// 更新失败，回滚事务
		tx.Rollback()
		return err
	}
	// 使用事务进行创建操作
	if err := tx.Create(order).Error; err != nil {
		// 创建失败，回滚事务
		tx.Rollback()
		return err
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil // 成功
}
