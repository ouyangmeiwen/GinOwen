package routers

import "GINOWEN/api"

type RouterApis struct {
	libItemApi        api.LibItemApi
	uploadfileApi     api.UploadfileApi
	sysauditlmslogApi api.SysauditlmslogApi
}

var Router = new(RouterApis)
