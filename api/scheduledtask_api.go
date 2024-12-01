package api

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ScheduledTaskApi struct {
}

// AddTask
// @Tags     Task
// @Summary  添加任务
// @Produce   application/json
// @Param    data  query     models.ScheduledTask 			true  "参数"
// @Success  200   {object}  utils.Response{msg=string}  "返回"
// @Security BearerAuth
// @Router   /Task/AddTask [get]
func AddTask(c *gin.Context) {
	var req models.ScheduledTask
	err := c.ShouldBindQuery(&req) //大小写敏感
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	err = ServicesGroup.TaskService.AddScheduledTask(req)
	if err != nil {
		global.OWEN_LOG.Error("新增任务失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("新增任务失败!"+err.Error(), c)
		return
	}
	utils.OkWithMessage("新增任务成功", c)
}
