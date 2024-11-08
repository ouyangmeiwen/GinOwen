package controllers

import (
	"GINOWEN/global"
	"GINOWEN/models/request"
	"GINOWEN/utils"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UploadfileControllers struct {
}

// uploadFileHandler 文件上传接口
// @Tags     File
// @Summary 上传文件
// @Description 接收一个文件和对象信息
// @Accept multipart/form-data
// @Produce application/json
// @Param name formData string true "文件名称"
// @Param description formData string false "文件描述"
// @Param cover formData bool false "文件描述"
// @Param file formData file true "上传的文件"
// @Success 200 {object} utils.Response{data=response.UploadResponse,msg=string} "上传成功"
// @Failure 400 {object} utils.Response{msg=string} "无效的请求"
// @Failure 500 {object} utils.Response{msg=string} "服务器内部错误"
// @Router /api/services/app/file/UploadFile [post]
func (b *UploadfileControllers) UploadFile(c *gin.Context) {
	var req request.UploadRequest
	err := c.ShouldBind(&req)
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		utils.FailWithMessage("上传失败!"+err.Error(), c)
		return
	}

	// 验证文件类型
	if !utils.IsValidFileType(file) {
		utils.FailWithMessage("不支持的文件类型", c)
		return
	}

	// 定义文件保存路径
	dst := "./uploads/" + file.Filename

	// 检查文件是否存在
	if _, err := os.Stat(dst); err == nil {
		if req.Cover {
			if err := os.Remove(dst); err != nil {
				utils.FailWithMessage("文件已存在，删除失败，请更换名字!", c)
				return
			}
		} else {
			// 方式 2: 修改文件名（添加时间戳）
			timestamp := time.Now().Format("20060102150405")
			dst = fmt.Sprintf("./uploads/%s_%s", timestamp, file.Filename)
		}
	} else if !os.IsNotExist(err) {
		// 如果检查文件出现其他错误
		utils.FailWithMessage("文件状态检查失败:!"+err.Error(), c)
		return
	}
	// 保存文件到指定路径
	if err := c.SaveUploadedFile(file, dst); err != nil {
		utils.FailWithMessage("保存失败!"+err.Error(), c)
		return
	}
	// 获取绝对路径
	absDst, err := filepath.Abs(dst)
	if err != nil {
		utils.FailWithMessage("获取文件绝对路径失败:!"+err.Error(), c)
		return
	}
	resp, err := ServicesApp.uploadfileService.UploadFileHandler(req)
	resp.FilePath = absDst
	if err != nil {
		global.OWEN_LOG.Error("上传失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("上传失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "上传成功", c)
}
