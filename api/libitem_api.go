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

// ImportExcel 导入图书信息
// @Summary 导入图书信息
// @Description 通过提供的Excel文件导入图书信息
// @Tags libitems
// @Accept  json
// @Produce  json
// @Param data body request.ImportExcelInput true "导入参数"
// @Success 200 {object} utils.Response{data=response.ImportExcelDto,msg=string} "导入成功，返回导入的数据"
// @Failure 400 {object} utils.Response{msg=string} "无效的请求"
// @Failure 500 {object} utils.Response{msg=string} "服务器内部错误"
// @Router /api/services/app/libitem/ImportExcel [post]
func (b *LibItemApi) ImportExcel(ctx *gin.Context) {
	var req request.ImportExcelInput
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage("参数无效", ctx)
		return
	}
	list, err := ServicesGroup.libitemService.ImportExcel(req)
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
// @Param data body request.ImportExcel2Input true "导入参数"
// @Success 200 {object} utils.Response{data=response.ImportExcelDto,msg=string} "导入成功，返回导入的数据"
// @Failure 400 {object} utils.Response{msg=string} "无效的请求"
// @Failure 500 {object} utils.Response{msg=string} "服务器内部错误"
// @Router /api/services/app/libitem/ImportExcel2 [post]
func (b *LibItemApi) ImportExcel2(ctx *gin.Context) {
	var req request.ImportExcel2Input
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage("参数无效", ctx)
		return
	}
	list, err := ServicesGroup.libitemService.ImportExcel2(req)
	if err != nil {
		global.OWEN_LOG.Error("导入失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("导入失败!"+err.Error(), ctx)
		return
	}
	utils.OkWithDetailed(list, "导入成功", ctx)
}
