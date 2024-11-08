package routers

import "GINOWEN/controllers"

type RouterControllers struct {
	orderController       controllers.OrderController
	userController        controllers.UserController
	libItemController     controllers.LibItemController
	uploadfileControllers controllers.UploadfileControllers
}

var Router = new(RouterControllers)
