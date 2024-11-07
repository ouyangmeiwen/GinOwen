package routes

import "GINOWEN/controllers"

type Controllers struct {
	orderController controllers.OrderController
	userController  controllers.UserController
}

var ApiApp = new(Controllers)
