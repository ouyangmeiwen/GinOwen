package services

import (
	"GINOWEN/models"
	"GINOWEN/models/request"
	"GINOWEN/models/response"

	"gorm.io/gorm"
)

type UserService struct {
}

func (s *UserService) GetAllUsers() ([]response.UserResponse, error) {
	users, err := RepoApp.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, response.UserResponse{
			ID:   user.ID,
			Name: user.Name,
		})
	}
	return userResponses, nil
}

func (s *UserService) CreateUser(userRequest request.UserRequest) (response.UserResponse, error) {
	user := models.User{Name: userRequest.Name}
	if err := RepoApp.userRepo.CreateUser(&user); err != nil {
		return response.UserResponse{}, err
	}
	return response.UserResponse{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (s *UserService) UpdateUserAndCreateOrder(userID uint, order *models.User) error {
	return RepoApp.userRepo.ExecuteInTransaction(func(tx *gorm.DB) error {
		if err := RepoApp.userRepo.UpdateUser(userID, map[string]interface{}{"status": "updated"}); err != nil {
			return err
		}

		if err := tx.Create(order).Error; err != nil {
			return err
		}

		return nil
	})
}
