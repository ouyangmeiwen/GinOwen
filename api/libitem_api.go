package api

import (
	"GINOWEN/global"
	"GINOWEN/models/request"
	"GINOWEN/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LibItemApi struct {
}

// ImportExcelByName 导入图书信息
// @Summary 导入图书信息
// @Description 通过提供的Excel文件导入图书信息
// @Tags libitems
// @Accept  json
// @Produce  json
// @Param data body request.ImportExcelByNameInput true "导入参数"
// @Success 200 {object} utils.Response{data=response.ImportExcelDto,msg=string} "导入成功，返回导入的数据"
// @Failure 400 {object} utils.Response{msg=string} "无效的请求"
// @Failure 500 {object} utils.Response{msg=string} "服务器内部错误"
// @Router /api/services/app/libitem/ImportExcelByName [post]
func (b *LibItemApi) ImportExcelByName(ctx *gin.Context) {
	var req request.ImportExcelByNameInput
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage("参数无效", ctx)
		return
	}
	list, err := ServicesGroup.libitemService.ImportExcelByName(req)
	if err != nil {
		global.OWEN_LOG.Error("导入失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("导入失败!"+err.Error(), ctx)
		return
	}
	utils.OkWithDetailed(list, "导入成功", ctx)
}

// ImportExcel 导入图书信息
// @Summary 导入图书信息
// @Description 通过提供的Excel文件导入图书信息
// @Tags libitems
// @Accept  json
// @Produce  json
// @Param data body request.ImportExcelByIndexInput true "导入参数"
// @Success 200 {object} utils.Response{data=response.ImportExcelDto,msg=string} "导入成功，返回导入的数据"
// @Failure 400 {object} utils.Response{msg=string} "无效的请求"
// @Failure 500 {object} utils.Response{msg=string} "服务器内部错误"
// @Router /api/services/app/libitem/ImportExcelByIndex [post]
func (b *LibItemApi) ImportExcelByIndex(ctx *gin.Context) {
	var req request.ImportExcelByIndexInput
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage("参数无效", ctx)
		return
	}
	list, err := ServicesGroup.libitemService.ImportExcelByIndex(req)
	if err != nil {
		global.OWEN_LOG.Error("导入失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("导入失败!"+err.Error(), ctx)
		return
	}
	utils.OkWithDetailed(list, "导入成功", ctx)
}
