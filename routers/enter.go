package routers

import "GINOWEN/controllers"

type RouterControllers struct {
	orderController          controllers.OrderController
	userController           controllers.UserController
	libItemController        controllers.LibItemController
	uploadfileControllers    controllers.UploadfileControllers
	sysauditlmslogController controllers.SysauditlmslogController
}

var Router = new(RouterControllers)
