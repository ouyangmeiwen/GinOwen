package routers

import "GINOWEN/api"

type RouterApis struct {
	libItemApi        api.LibItemApi
	uploadfileApi     api.UploadfileApi
	sysauditlmslogApi api.SysauditlmslogApi
	jwtApi            api.JWTAPI
	ipApi             api.IPApi
}

var ApiGroup = new(RouterApis)
