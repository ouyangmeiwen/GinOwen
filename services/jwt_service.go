package services

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"errors"
)

type JWTService struct {
}

func (JWTService) GetUserByUsername(req request.LoginRequest) (resp response.LoginResponse, err error) {
	var user models.OwenUser
	err = global.OWEN_DB.Model(&models.OwenUser{}).Where(" Username=?", req.Username).First(&user).Error
	if err != nil {
		return resp, errors.New("user not found")
	}
	resp.User = user

	var role models.OwenRole
	err = global.OWEN_DB.Model(&models.OwenRole{}).Where(" Id=?", user.RoleID).First(&role).Error
	if err == nil {
		resp.User.Role = role
	}
	return resp, err
}
