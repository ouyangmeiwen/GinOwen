package services

import (
	"GINOWEN/models"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"GINOWEN/repositories"

	"gorm.io/gorm"
)

type UserService struct {
	Repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() ([]response.UserResponse, error) {
	users, err := s.Repo.GetAllUsers()
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
	if err := s.Repo.CreateUser(&user); err != nil {
		return response.UserResponse{}, err
	}
	return response.UserResponse{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (s *UserService) UpdateUserAndCreateOrder(userID uint, order *models.User) error {
	return s.Repo.ExecuteInTransaction(func(tx *gorm.DB) error {
		if err := s.Repo.UpdateUser(userID, map[string]interface{}{"status": "updated"}); err != nil {
			return err
		}

		if err := tx.Create(order).Error; err != nil {
			return err
		}

		return nil
	})
}
