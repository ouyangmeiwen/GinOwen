package routes

import (
	"GINOWEN/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterUserRoutes(r *gin.Engine, userController *controllers.UserController) {
	api := r.Group("api/user")
	{
		api.GET("GetUsers", userController.GetUsers)
		api.POST("CreateUser", userController.CreateUser)
	}
}
