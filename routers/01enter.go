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
	inventoryApi      api.LibiteminventoryinfoApi
	librowApi         api.LibrowApi
	flyreadApi        api.FlyReadController
}

var ApiGroup = new(RouterApis)
