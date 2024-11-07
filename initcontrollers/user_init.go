// initrouter/user.go
package initcontrollers

import (
	"GINOWEN/controllers"
	"GINOWEN/repositories"
	"GINOWEN/services"

	"gorm.io/gorm"
)

// InitUserController 初始化用户控制器
func InitUserController(db *gorm.DB) *controllers.UserController {
	userRepo := &repositories.GormUserRepository{DB: db}
	userService := services.NewUserService(userRepo)
	return controllers.NewUserController(userService)
}
