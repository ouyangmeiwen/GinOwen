package api

import (
	"GINOWEN/auth"
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/models/request"
	"GINOWEN/utils"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type JWTAPI struct {
}

// JWT 登录接口
// @Summary User Login
// @Summary 用户登录，获取JWT Token
// @Description 用户输入用户名和密码，验证通过后返回JWT Token。
// @Tags Auth
// @Accept json
// @Produce json
// @Param loginRequest body request.LoginRequest true "用户登录请求"
// @Success 200 {string} string "JWT token"
// @Failure 400 {string} string "Invalid input"
// @Failure 401 {string} string "Invalid username or password"
// @Router /auth/Login [post]
func (JWTAPI) Login(ctx *gin.Context) {
	var loginReq request.LoginRequest
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		utils.FailWithMessage("Invalid input", ctx)
		return
	}

	// 假设我们从数据库获取用户数据
	user, err := ServicesGroup.jwtService.GetUserByUsername(loginReq)
	if err != nil {
		utils.FailWithMessage("Invalid username or password", ctx)
		return
	}

	// 检查密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(user.User.Password), []byte(loginReq.Password))
	if err != nil {

		utils.FailWithMessage("Invalid username or password", ctx)
		return
	}

	// 用户身份验证通过，生成 JWT
	token, err := auth.GenerateJWT(user.User.ID, user.User.RoleID)
	if err != nil {
		utils.FailWithMessage("Could not generate token", ctx)
		return
	}
	TokenExpire := global.OWEN_CONFIG.System.TokenExpire
	utils.SetToken(ctx, token, TokenExpire)
	var back = context.Background()
	if len(global.OWEN_CONFIG.Redis.Addr) > 0 {
		rolestr, err := utils.ToJSON(user.User.Role)
		if err == nil {
			global.OWEN_REDIS.Set(back, token, rolestr, time.Duration(TokenExpire)*time.Second)
		}
	}
	utils.OkWithDetailed(token, "success", ctx)
}

func (JWTAPI) LoginOut(ctx *gin.Context) {
	utils.ClearToken(ctx)
	utils.OkWithDetailed("", "success", ctx)
}

// Register 注册接口
// @Summary 用户注册
// @Description 用户提交注册信息，注册成功后返回成功消息。
// @Tags Auth
// @Accept json
// @Produce json
// @Param registerRequest body request.RegisterRequest true "用户注册请求"
// @Success 200 {string} string "User registered successfully"
// @Failure 400 {string} string "Invalid input"
// @Failure 409 {string} string "Username already exists"
// @Router /auth/Register [post]
func (JWTAPI) Register(ctx *gin.Context) {
	var registerReq request.RegisterRequest
	if err := ctx.ShouldBindJSON(&registerReq); err != nil {
		utils.FailWithMessage("Invalid input", ctx)
		return
	}

	// 检查用户名是否已存在
	var loginReq request.LoginRequest
	loginReq.Username = registerReq.Username
	_, err := ServicesGroup.jwtService.GetUserByUsername(loginReq)
	if err == nil {
		utils.FailWithMessage("Username already exists", ctx)
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.FailWithMessage("Error encrypting password", ctx)
		return
	}

	// 创建用户
	newUser := models.OwenUser{
		Username: registerReq.Username,
		Password: string(hashedPassword), // 存储加密后的密码
		RoleID:   registerReq.RoleID,     // 可以从请求中获取角色ID
	}

	// 存储用户到数据库
	if err := global.OWEN_DB.Create(&newUser); err != nil {
		utils.FailWithMessage("Error creating user", ctx)
		return
	}

	// 注册成功，返回成功消息
	utils.OkWithMessage("User registered successfully", ctx)
}
