package services

import (
	"GINOWEN/models"
	"GINOWEN/repositories"

	"gorm.io/gorm"
)

type UserRequest struct {
	Name string `json:"name" binding:"required"`
}

type UserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserService struct {
	Repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() ([]UserResponse, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, UserResponse{
			ID:   user.ID,
			Name: user.Name,
		})
	}
	return userResponses, nil
}

func (s *UserService) CreateUser(userRequest UserRequest) (UserResponse, error) {
	user := models.User{Name: userRequest.Name}
	if err := s.Repo.CreateUser(&user); err != nil {
		return UserResponse{}, err
	}
	return UserResponse{
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
