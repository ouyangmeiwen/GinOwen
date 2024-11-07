package controllers

import (
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"GINOWEN/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
type UserController struct {
	UserService *services.UserService
}

// NewUserController 创建一个新的 UserController 实例
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

// GetUsers 获取所有用户
// @Summary 获取所有用户
// @Description 获取所有用户信息
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} response.CommonResponse "获取到的所有用户列表"
// @Failure 500 {object} response.CommonResponse "内部服务器错误"
// @Router /users [get]
func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := c.UserService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.CommonResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve users",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.CommonResponse{
		Code:    http.StatusOK,
		Message: "Users retrieved successfully",
		Data:    users,
	})
}

// CreateUser 创建用户
// @Summary 创建新用户
// @Description 创建一个新的用户并返回用户信息
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body request.UserRequest true "用户信息"
// @Success 200 {object} response.CommonResponse "成功创建的用户信息"
// @Failure 400 {object} response.CommonResponse "无效的输入"
// @Failure 500 {object} response.CommonResponse "创建用户失败"
// @Router /users [post]
func (c *UserController) CreateUser(ctx *gin.Context) {
	var req request.UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.CommonResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid input data",
			Data:    nil,
		})
		return
	}

	user, err := c.UserService.CreateUser(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.CommonResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create user",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Code:    http.StatusOK,
		Message: "User created successfully",
		Data:    user,
	})
}
