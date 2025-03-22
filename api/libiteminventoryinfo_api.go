package api

import (
	"GINOWEN/global"
	"GINOWEN/models/request"
	"GINOWEN/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LibiteminventoryinfoApi struct {
}

// QueryInventory
// @Tags     Libiteminventory
// @Summary  查询盘点位置
// @Produce   application/json
// @Param    data  query     request.GetInventoryInput 			true  "入参"
// @Success  200   {object}  utils.Response{data=[]response.GetInventoryDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/Libiteminventory/QueryInventory [get]
func (b *LibiteminventoryinfoApi) QueryInventory(c *gin.Context) {

	// 打印查询字符串
	queryStr := c.Request.URL.RawQuery
	fmt.Println("Query String:", queryStr)
	var req request.GetInventoryInput

	//post
	//err := c.ShouldBindJSON(&req)

	//get
	err := c.ShouldBindQuery(&req) //大小写敏感

	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	list, err := ServicesGroup.InventoryService.GetInventoryList(req)
	if err != nil {
		global.OWEN_LOG.Error("获取失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("获取失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(list, "获取成功", c)
}

// CreateInventory
// @Tags     Libiteminventory
// @Summary  新增盘点位置
// @Produce   application/json
// @Param    data  body     request.CreateInventoryInput 			true  "参数"
// @Success  200   {object}  utils.Response{data=[]response.CreateInventoryDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/Libiteminventory/CreateInventory [post]
func (b *LibiteminventoryinfoApi) CreateInventory(c *gin.Context) {
	var req request.CreateInventoryInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	list, err := ServicesGroup.InventoryService.CreateInventory(req)
	if err != nil {
		global.OWEN_LOG.Error("创建失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("创建失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(list, "创建成功", c)
}

// UpdateInventory
// @Tags     Libiteminventory
// @Summary  修改盘点位置
// @Produce   application/json
// @Param    data  body     request.UpdateInventoryInput 			true  "参数"
// @Success  200   {object}  utils.Response{data=[]response.UpdateInventoryDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /api/services/app/Libiteminventory/UpdateInventory [put]
func (b *LibiteminventoryinfoApi) UpdateInventory(c *gin.Context) {
	var req request.UpdateInventoryInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	list, err := ServicesGroup.InventoryService.UpdateInventory(req)
	if err != nil {
		global.OWEN_LOG.Error("更新失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("更新失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(list, "更新成功", c)
}

// DeleteInventory
// @Tags     Libiteminventory
// @Summary  删除盘点位置
// @Produce   application/json
// @Param    data  query     request.DeleteInventoryInput 			true  "参数"
// @Success  200   {object}  utils.Response{data=[]response.DeleteInventoryDto,msg=string}  "返回清单"
// @Security BearerAuth
// @Router   /api/services/app/Libiteminventory/DeleteInventory [delete]
func (b *LibiteminventoryinfoApi) DeleteInventory(c *gin.Context) {

	var req request.DeleteInventoryInput
	err := c.ShouldBindQuery(&req) //大小写敏感

	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	list, err := ServicesGroup.InventoryService.DeleteInventory(req)
	if err != nil {
		global.OWEN_LOG.Error("删除失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("删除失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(list, "删除成功", c)
}
