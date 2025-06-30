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

// UploadStruct
// @Tags     FlyRead
// @Summary  结构推送
// @Produce  application/json
// @Param    data  body       request.UploadStructInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.UploadStructDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/UploadStruct [post]
func (b *FlyReadController) UploadStruct(c *gin.Context) {
	var req request.UploadStructInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.UploadStructDto
	resp, err = ServicesGroup.flyreadAppService.UploadStruct(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// UploadLibItemLoc
// @Tags     FlyRead
// @Summary  推送图书定位
// @Produce  application/json
// @Param    data  body       request.UploadLibItemLocInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.UploadLibItemLocDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/UploadLibItemLoc [post]
func (b *FlyReadController) UploadLibItemLoc(c *gin.Context) {
	var req request.UploadLibItemLocInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.UploadLibItemLocDto
	resp, err = ServicesGroup.flyreadAppService.UploadLibItemLoc(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// UploadRow
// @Tags     FlyRead
// @Summary  推送图书定位
// @Produce  application/json
// @Param    data  body       request.UploadRowInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.UploadRowDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/UploadRow [post]
func (b *FlyReadController) UploadRow(c *gin.Context) {
	var req request.UploadRowInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.UploadRowDto
	resp, err = ServicesGroup.flyreadAppService.UploadRow(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// Inventory
// @Tags     FlyRead
// @Summary  盘点指令下发
// @Produce  application/json
// @Param    data  body       request.InventoryInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.InventoryDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/Inventory [post]
func (b *FlyReadController) Inventory(c *gin.Context) {
	var req request.InventoryInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.InventoryDto
	resp, err = ServicesGroup.flyreadAppService.Inventory(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// GetRobotRouterlist
// @Tags     FlyRead
// @Summary  获取机器人规划路线列表
// @Produce  application/json
// @Param    data  query       request.GetRobotRouterlistInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.GetRobotRouterlistDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/GetRobotRouterlist [get]
func (b *FlyReadController) GetRobotRouterlist(c *gin.Context) {

	var req request.GetRobotRouterlistInput
	err := c.ShouldBindQuery(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.GetRobotRouterlistDto
	resp.Items, err = ServicesGroup.flyReadAppManager.GetRobotRouterlist(req.Robotid, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// InventoryHis
// @Tags     FlyRead
// @Summary  盘点任务查询
// @Produce  application/json
// @Param    data  body       request.InventoryHisInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.InventoryHisDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/InventoryHis [post]
func (b *FlyReadController) InventoryHis(c *gin.Context) {
	var req request.InventoryHisInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.InventoryHisDto
	resp, err = ServicesGroup.flyreadAppService.InventoryHis(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// InventoryList
// @Tags     FlyRead
// @Summary  盘点摄像头历史任务查询
// @Produce  application/json
// @Param    data  body       request.InventoryListInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.InventoryListDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/InventoryList [post]
func (b *FlyReadController) InventoryList(c *gin.Context) {
	var req request.InventoryListInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.InventoryListDto
	resp, err = ServicesGroup.flyreadAppService.InventoryList(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// SetBussiness
// @Tags     FlyRead
// @Summary  设置业务规则
// @Produce  application/json
// @Param    data  body       request.SetBussinessInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.SetBussinessDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/SetBussiness [post]
func (b *FlyReadController) SetBussiness(c *gin.Context) {
	var req request.SetBussinessInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.SetBussinessDto
	resp, err = ServicesGroup.flyreadAppService.SetBussiness(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// GetInventorySet
// @Tags     FlyRead
// @Summary  获取盘点设置
// @Produce  application/json
// @Param    data  query       request.GetInventorySetInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.GetInventorySetDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/GetInventorySet [get]
func (b *FlyReadController) GetInventorySet(c *gin.Context) {

	var req request.GetInventorySetInput
	err := c.ShouldBindQuery(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.GetInventorySetDto
	resp, err = ServicesGroup.flyreadAppService.GetInventorySet(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// GetEnableRow
// @Tags     FlyRead
// @Summary  获取可用的层架关系
// @Produce  application/json
// @Param    data  query       request.GetEnableRowInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.GetEnableRowDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/GetEnableRow [get]
func (b *FlyReadController) GetEnableRow(c *gin.Context) {

	var req request.GetEnableRowInput
	err := c.ShouldBindQuery(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.GetEnableRowDto
	resp, err = ServicesGroup.flyreadAppService.GetEnableRow(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// GetRobotlist
// @Tags     FlyRead
// @Summary  获取可用的层架关系
// @Produce  application/json
// @Param    data  query       request.GetRobotlistInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.GetRobotlistDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/GetRobotlist [get]
func (b *FlyReadController) GetRobotlist(c *gin.Context) {

	var req request.GetRobotlistInput
	err := c.ShouldBindQuery(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.GetRobotlistDto
	resp, err = ServicesGroup.flyreadAppService.GetRobotlist(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// GetCaseCodeImage
// @Tags     FlyRead
// @Summary  获取可用的层架关系
// @Produce  application/json
// @Param    data  body     request.GetCaseCodeImageInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.GetCaseImgsDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/GetCaseCodeImage [post]
func (b *FlyReadController) GetCaseCodeImage(c *gin.Context) {

	var req request.GetCaseCodeImageInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.GetCaseImgsDto
	var dto dto.GetCaseImgsDto
	dto, err = ServicesGroup.flyReadAppManager.GetCaseImgs(req.Layer_codes, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	resp.GetCaseImgsDto = dto
	utils.OkWithDetailed(resp, "成功", c)
}

// GetOcrImgs
// @Tags     FlyRead
// @Summary  获取层图片
// @Produce  application/json
// @Param    data  query       request.GetOcrImgsInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.GetOcrImgsDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/GetOcrImgs [get]
func (b *FlyReadController) GetOcrImgs(c *gin.Context) {

	var req request.GetOcrImgsInput
	err := c.ShouldBindQuery(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.GetOcrImgsDto
	var dto_response dto.GetOcrImgsDto
	dto_response, err = ServicesGroup.flyReadAppManager.GetOcrImgs(req.LayerCode, req.ItemBarcode, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	resp.GetOcrImgsDto = dto_response
	utils.OkWithDetailed(resp, "成功", c)
}

// InventorySet
// @Tags     FlyRead
// @Summary  盘点设置
// @Produce  application/json
// @Param    data  body     request.InventorySetInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.InventorySetDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/InventorySet [post]
func (b *FlyReadController) InventorySet(c *gin.Context) {

	var req request.InventorySetInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.InventorySetDto
	resp, err = ServicesGroup.flyreadAppService.InventorySet(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// CreatWork
// @Tags     FlyRead
// @Summary  创建任务
// @Produce  application/json
// @Param    data  body     request.CreatWorkInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.CreatWorkDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/CreatWork [post]
func (b *FlyReadController) CreatWork(c *gin.Context) {

	var req request.CreatWorkInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.CreatWorkDto
	resp, err = ServicesGroup.flyreadAppService.CreatWork(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// UpdateWork
// @Tags     FlyRead
// @Summary  更新任务
// @Produce  application/json
// @Param    data  body     request.UpdateWorkInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.UpdateWorkDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/UpdateWork [post]
func (b *FlyReadController) UpdateWork(c *gin.Context) {

	var req request.UpdateWorkInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.UpdateWorkDto
	resp, err = ServicesGroup.flyreadAppService.UpdateWork(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// WorkList
// @Tags     FlyRead
// @Summary  任务列表
// @Produce  application/json
// @Param    data  body     request.WorkListInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.PageWorkListDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/WorkList [post]
func (b *FlyReadController) WorkList(c *gin.Context) {

	var req request.WorkListInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.PageWorkListDto
	resp, err = ServicesGroup.flyreadAppService.WorkList(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// DeleteWork
// @Tags     FlyRead
// @Summary  获取层图片
// @Produce  application/json
// @Param    data  query       request.DeleteWorkInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.DeleteWorkDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/DeleteWork [get]
func (b *FlyReadController) DeleteWork(c *gin.Context) {

	var req request.DeleteWorkInput
	err := c.ShouldBindQuery(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.DeleteWorkDto
	resp, err = ServicesGroup.flyreadAppService.DeleteWork(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// DetailList
// @Tags     FlyRead
// @Summary  获取任务详情
// @Produce  application/json
// @Param    data  body     request.DetailListInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.PageDetailListDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/DetailList [post]
func (b *FlyReadController) DetailList(c *gin.Context) {

	var req request.DetailListInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.PageDetailListDto
	resp, err = ServicesGroup.flyreadAppService.DetailList(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}

// DetailStatusList
// @Tags     FlyRead
// @Summary  获取任务列表层详情
// @Produce  application/json
// @Param    data  body     request.DetailStatusListInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=response.PageDetailStatusListDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/DetailStatusList [post]
func (b *FlyReadController) DetailStatusList(c *gin.Context) {

	var req request.DetailStatusListInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var resp response.PageDetailStatusListDto
	resp, err = ServicesGroup.flyreadAppService.DetailStatusList(req, utils.GetCurrentTenantTd(c))
	if err != nil {
		global.OWEN_LOG.Error("失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "成功", c)
}
