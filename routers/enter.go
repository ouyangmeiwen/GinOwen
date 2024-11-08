package routers

import "GINOWEN/controllers"

type Controllers struct {
	orderController   controllers.OrderController
	userController    controllers.UserController
	libItemController controllers.LibItemController
}

var ApiApp = new(Controllers)
