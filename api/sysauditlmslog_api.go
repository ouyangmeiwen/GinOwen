package api

import (
	"GINOWEN/global"
	"GINOWEN/models/request"
	"GINOWEN/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysauditlmslogApi struct {
}

// QueryLmsLog
// @Tags     Sysauditlmslog
// @Summary  查询LMS日志
// @Produce   application/json
// @Param    data  query     request.QueryLmsInput 			true  "参数"
// @Success  200   {object}  utils.Response{data=[]response.QueryLmsDto,msg=string}  "返回清单"
// @Security BearerAuth
// @Router   /api/services/app/Sysauditlmslog/QueryLmsLog [get]
func (b *SysauditlmslogApi) QueryLmsLog(c *gin.Context) {

	// 打印查询字符串
	queryStr := c.Request.URL.RawQuery
	fmt.Println("Query String:", queryStr)

	//post
	var req request.QueryLmsInput
	//err := c.ShouldBindJSON(&req)
	//get
	err := c.ShouldBindQuery(&req) //大小写敏感

	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	list, err := ServicesGroup.sysauditlmslogService.QueryLmsLog(req)
	if err != nil {
		global.OWEN_LOG.Error("获取失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("获取失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(list, "获取成功", c)
}

// {
// 	"BrowserInfo": "string",
// 	"ClientIpAddress": "string",
// 	"ClientName": "string",
// 	"CustomData": "string",
// 	"Exception": "string",
// 	"ExecutionDuration": 0,
// 	"ExecutionTime": "2024-10-02T15:04:05+00:00",
// 	"Id": 0,
// 	"ImpersonatorTenantId": 0,
// 	"ImpersonatorUserId": 0,
// 	"MethodName": "string",
// 	"Parameters": "string",
// 	"ReturnValue": "string",
// 	"ServiceName": "string",
// 	"TenantId": 0,
// 	"UserId": 0
//   }

// CreateLmsLog
// @Tags     Sysauditlmslog
// @Summary  新增LMS审计日志
// @Produce   application/json
// @Param    data  body     request.CreateLmsLogInput 			true  "参数"
// @Success  200   {object}  utils.Response{data=[]response.CreateLmsLogDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/Sysauditlmslog/CreateLmsLog [post]
func (b *SysauditlmslogApi) CreateLmsLog(c *gin.Context) {
	var req request.CreateLmsLogInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	list, err := ServicesGroup.sysauditlmslogService.CreateLmsLog(req)
	if err != nil {
		global.OWEN_LOG.Error("创建失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("创建失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(list, "创建成功", c)
}

// {
// 	"BrowserInfo": "string",
// 	"ClientIpAddress": "string",
// 	"ClientName": "string",
// 	"CustomData": "string",
// 	"Exception": "string",
// 	"ExecutionDuration": 0,
// 	 "ExecutionTime": "2024-10-02T15:04:05+00:00",
// 	"Id": 76541,
// 	"ImpersonatorTenantId": 0,
// 	"ImpersonatorUserId": 0,
// 	"MethodName": "string",
// 	"Parameters": "string",
// 	"ReturnValue": "string",
// 	"ServiceName": "哈哈哈",
// 	"TenantId": 0,
// 	"UserId": 0
//   }

// UpdateLmsLog
// @Tags     Sysauditlmslog
// @Summary  修改LMS审计日志
// @Produce   application/json
// @Param    data  body     request.UpdateLmsLogInput 			true  "参数"
// @Success  200   {object}  utils.Response{data=[]response.UpdateLmsLogDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/Sysauditlmslog/UpdateLmsLog [put]
func (b *SysauditlmslogApi) UpdateLmsLog(c *gin.Context) {
	var req request.UpdateLmsLogInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	list, err := ServicesGroup.sysauditlmslogService.UpdateLmsLog(req)
	if err != nil {
		global.OWEN_LOG.Error("更新失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("更新失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(list, "更新成功", c)
}

// DeleteLmsLog
// @Tags     Sysauditlmslog
// @Summary  查询LMS日志
// @Produce   application/json
// @Param    data  query     request.DeleteLmsLogInput 			true  "参数"
// @Success  200   {object}  utils.Response{data=[]response.DeleteLmsLogDto,msg=string}  "返回清单"
// @Security BearerAuth
// @Router   /api/services/app/Sysauditlmslog/DeleteLmsLog [delete]
func (b *SysauditlmslogApi) DeleteLmsLog(c *gin.Context) {

	var req request.DeleteLmsLogInput
	err := c.ShouldBindQuery(&req) //大小写敏感

	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	list, err := ServicesGroup.sysauditlmslogService.DeleteLmsLog(req)
	if err != nil {
		global.OWEN_LOG.Error("删除失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("删除失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(list, "删除成功", c)
}
