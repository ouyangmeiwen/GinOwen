package routers

import "GINOWEN/api"

type RouterApis struct {
	libItemApi        api.LibItemApi
	uploadfileApi     api.UploadfileApi
	sysauditlmslogApi api.SysauditlmslogApi
	loginApi          api.LoginAPI
	blacklistApi      api.BlackListApi
	gpiApi            api.GPIApi
}

var ApiGroup = new(RouterApis)
