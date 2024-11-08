package api

import (
	"GINOWEN/models/request"
	"GINOWEN/utils" // Import the response package

	"github.com/gin-gonic/gin"
)

// UserApi 用户控制器
type UserApi struct {
}

// GetUsers 获取所有用户
// @Summary 获取所有用户
// @Description 获取所有用户信息
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.Response{data=[]response.UserDto,msg=string} "获取到的所有用户列表"
// @Failure 500 {object} utils.Response{msg=string} "内部服务器错误"
// @Router /api/services/app/user/GetUsers [get]
func (c *UserApi) GetUsers(ctx *gin.Context) {
	users, err := ServicesApp.userService.GetAllUsers()
	if err != nil {
		utils.FailWithMessage("Failed to retrieve", ctx)
		return
	}
	utils.OkWithDetailed(users, "get successfully", ctx)
}

// CreateUser 创建用户
// @Summary 创建新用户
// @Description 创建一个新的用户并返回用户信息
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body request.UserRequest true "用户信息"
// @Success 200 {object} utils.Response{data=response.UserDto,msg=string} "成功创建的用户信息"
// @Failure 400 {object} utils.Response{msg=string} "无效的输入"
// @Failure 500 {object} utils.Response{msg=string} "创建用户失败"
// @Router /api/services/app/user/CreateUser [post]
func (c *UserApi) CreateUser(ctx *gin.Context) {
	var req request.UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.FailWithMessage("Invalid request data", ctx)
		return
	}
	user, err := ServicesApp.userService.CreateUser(req)
	if err != nil {
		utils.FailWithMessage("Failed to create", ctx)
		return
	}
	utils.OkWithDetailed(user, "created successfully", ctx)
}
