package api

import "GINOWEN/services"

type Services struct {
	orderService          services.OrderService
	userService           services.UserService
	libitemService        services.LibItemService
	uploadfileService     services.UploadfileService
	sysauditlmslogService services.SysauditlmslogService
}

var ServicesApp = new(Services)
