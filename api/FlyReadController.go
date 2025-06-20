package api

import (
	"GINOWEN/global"
	"GINOWEN/models/dto"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"GINOWEN/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FlyReadController struct {
}

// Hello
// @Tags     FlyRead
// @Summary  查询书架
// @Produce  application/json
// @Param    data  query     request.HelloInput 			true  "参数"
// @Success  200   {object}  utils.Response{data=[]response.HelloResp,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/Hello [get]
func (b *FlyReadController) Hello(c *gin.Context) {
	// 打印查询字符串
	queryStr := c.Request.URL.RawQuery
	fmt.Println("Query String:", queryStr)
	//post
	var req request.HelloInput
	//err := c.ShouldBindJSON(&req)
	//get
	err := c.ShouldBindQuery(&req) //大小写敏感

	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	list, err := ServicesGroup.flyreadAppService.Hello(req)
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(list, "成功", c)
}

// SetFlyReadSetting
// @Tags     FlyRead
// @Summary  设置飞阅参数 注意FlyReadIp 如果带上http参数则不启用端口参数
// @Produce  application/json
// @Param    data  body       dto.FlyReadSetting 			true  "参数"
// @Success  200   {object}  utils.Response{data=interface{},msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/SetFlyReadSetting [post]
func (b *FlyReadController) SetFlyReadSetting(c *gin.Context) {
	//post
	var req dto.FlyReadSetting
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	err = ServicesGroup.flyReadAppManager.SetFlyReadSetting(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithMessage("成功", c)
}

// GetFlyReadSetting
// @Tags     FlyRead
// @Summary  获取当前机构的飞阅参数
// @Produce  application/json
// @Success  200   {object}  utils.Response{data=dto.FlyReadSetting,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/GetFlyReadSetting [get]
func (b *FlyReadController) GetFlyReadSetting(c *gin.Context) {
	// 打印查询字符串
	queryStr := c.Request.URL.RawQuery
	fmt.Println("Query String:", queryStr)
	result, err := ServicesGroup.flyReadAppManager.GetFlyReadSetting(utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("获取失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("获取失败!"+err.Error(), c)
		return
	}
	if result.FlyLocType == "2" {
		switch result.FlyTempLoc {
		case "1":
			result.FlyNewLocMode = "2"
		case "2":
			result.FlyNewLocMode = "3"
		default:
			result.FlyNewLocMode = "1"
		}
	} else {
		if result.FlyLocType == "0" {
			result.FlyNewLocMode = "5"
		} else {
			result.FlyNewLocMode = "4"
		}
	}
	if result.RowShape == "" {
		result.RowShape = "1"
	}
	utils.OkWithDetailed(result, "成功", c)
}

// GetFlyReadToken
// @Tags     FlyRead
// @Summary  查询token
// @Produce  application/json
// @Param    data  query     request.GetFlyTokenInput 			true  "参数"
// @Success  200   {object}  utils.Response{data=response.GetFlyTokenInputResp,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/GetFlyReadToken [get]
func (b *FlyReadController) GetFlyReadToken(c *gin.Context) {
	var req request.GetFlyTokenInput
	err := c.ShouldBindQuery(&req) //大小写敏感
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	token, err := ServicesGroup.flyReadAppManager.GetToken(utils.GetCurrentTenantTd(c), req.IsForceRefresh)
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	var response response.GetFlyTokenInputResp
	response.AccessToken = token
	utils.OkWithDetailed(response, "成功", c)
}

// UploadLibItem
// @Tags     FlyRead
// @Summary  图书推送
// @Produce  application/json
// @Param    data  body       request.UploadLibItemInput 			true  "参数"
// @Success  200   {object}  utils.Response{data=response.UploadLibItemResp,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/UploadLibItem [post]
func (b *FlyReadController) UploadLibItem(c *gin.Context) {
	var req request.UploadLibItemInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.UploadLibItemResp
	resp, err = ServicesGroup.flyreadAppService.UploadLibItem(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// UploadTenant
// @Tags     FlyRead
// @Summary  分馆推送
// @Produce  application/json
// @Param    data  body       request.UploadTenantInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.UploadTenantDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/UploadTenant [post]
func (b *FlyReadController) UploadTenant(c *gin.Context) {
	var req request.UploadTenantInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	if req.Tenantid <= 0 {
		req.Tenantid = utils.GetCurrentTenantTd(c)
	}
	var resp response.UploadTenantDto
	resp, err = ServicesGroup.flyreadAppService.UploadTenant(req)
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}
