package routers

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterUserRoutes(r *gin.Engine) {
	api := r.Group("api/user")
	{
		api.GET("GetUsers", ApiApp.userController.GetUsers)
		api.POST("CreateUser", ApiApp.userController.CreateUser)
	}
}
