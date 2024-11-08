package routers

import "GINOWEN/api"

type RouterApis struct {
	orderApi          api.OrderApi
	userApi           api.UserApi
	libItemApi        api.LibItemApi
	uploadfileApi     api.UploadfileApi
	sysauditlmslogApi api.SysauditlmslogApi
}

var Router = new(RouterApis)
