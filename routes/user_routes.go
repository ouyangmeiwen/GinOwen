package routes

import (
	"GINOWEN/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterUserRoutes(r *gin.Engine, userController *controllers.UserController) {
	api := r.Group("/api")
	{
		api.GET("/users", userController.GetUsers)
		api.POST("/users", userController.CreateUser)
	}
}
