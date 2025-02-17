package request

import (
	"GINOWEN/models"
	"GINOWEN/utils"
)

// get 方式 大小写是敏感的
type QueryLmsInput struct {
	utils.PagedResultRequest
	IsException       *bool   `json:"IsException" form:"IsException"`             //是否异常
	ExecutionDuration int     `json:"ExecutionDuration" form:"ExecutionDuration"` //最大执行时间
	MethodName        *string `json:"MethodName" form:"MethodName"`               //方法
	Parameters        *string `json:"Parameters" form:"Parameters"`               //参数
	ServiceName       *string `json:"ServiceName" form:"ServiceName"`             //服务名
	TenantId          int     `json:"TenantId" form:"TenantId"`                   //租户
}

type CreateLmsLogInput struct {
	models.Sysauditlmslog
}
type UpdateLmsLogInput struct {
	models.Sysauditlmslog
}
type DeleteLmsLogInput struct {
	ID int64 `json:"Id" form:"Id"` //根据主键删除
}
