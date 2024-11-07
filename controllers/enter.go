package controllers

import "GINOWEN/services"

type Services struct {
	orderService services.OrderService
	userService  services.UserService
}

var ServicesApp = new(Services)
