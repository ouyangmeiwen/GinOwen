package api

import (
	"GINOWEN/services"
	"GINOWEN/services/manager"
)

type Services struct {
	libitemService        services.LibItemService
	uploadfileService     services.UploadfileService
	sysauditlmslogService services.SysauditlmslogService
	jwtService            services.JWTService
	TaskService           services.ScheduledTaskService
	InventoryService      services.LibiteminventoryinfoService
	librowService         services.LibrowService
	flyreadAppService     services.FlyReadAppService
	flyReadAppManager     manager.FlyReadAppManager
}

var ServicesGroup = new(Services)
