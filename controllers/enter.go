package controllers

import "GINOWEN/services"

type Services struct {
	orderService      services.OrderService
	userService       services.UserService
	libitemService    services.LibItemService
	uploadfileService services.UploadfileService
}

var ServicesApp = new(Services)
