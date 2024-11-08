package api

import "GINOWEN/services"

type Services struct {
	libitemService        services.LibItemService
	uploadfileService     services.UploadfileService
	sysauditlmslogService services.SysauditlmslogService
}

var ServicesApp = new(Services)
