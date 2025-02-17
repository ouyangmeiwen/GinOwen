package api

import (
	"GINOWEN/global"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"GINOWEN/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ScheduledTaskApi struct {
}

// AddTask
// @Tags     Task
// @Summary  添加任务
// @Produce   application/json
// @Param    data  query     request.AddScheduledTaskInput 			true  "参数"
// @Success  200   {object}  utils.Response{data=response.AddScheduledTaskDto,msg=string}  "返回"
// @Security BearerAuth
// @Router   /Task/AddTask [get]
func (ScheduledTaskApi) AddTask(c *gin.Context) {
	var req request.AddScheduledTaskInput
	err := c.ShouldBindQuery(&req) //大小写敏感
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	if req.IntervalSeconds == nil || *req.IntervalSeconds <= 0 {
		zero := 0
		req.IntervalSeconds = &zero
	}
	if time.Now().After(req.ScheduleTime) {
		utils.FailWithMessage("任务执行时间必须在今天之后！", c)
		return
	}

	var resp response.AddScheduledTaskDto
	resp, err = ServicesGroup.TaskService.AddScheduledTask(req)
	if err != nil {
		global.OWEN_LOG.Error("新增任务失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("新增任务失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(resp, "新增任务成功", c)
}
