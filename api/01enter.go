package api

import "GINOWEN/services"

type Services struct {
	libitemService        services.LibItemService
	uploadfileService     services.UploadfileService
	sysauditlmslogService services.SysauditlmslogService
	jwtService            services.JWTService
	TaskService           services.ScheduledTaskService
	InventoryService      services.LibiteminventoryinfoService
	librowService         services.LibrowService
}

var ServicesGroup = new(Services)
