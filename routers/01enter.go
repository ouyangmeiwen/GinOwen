package routers

import "GINOWEN/api"

type RouterApis struct {
	libItemApi        api.LibItemApi
	uploadfileApi     api.UploadfileApi
	sysauditlmslogApi api.SysauditlmslogApi
	loginApi          api.LoginAPI
	blacklistApi      api.BlackListApi
	gpiApi            api.GPIApi
	taskApi           api.ScheduledTaskApi
}

var ApiGroup = new(RouterApis)
