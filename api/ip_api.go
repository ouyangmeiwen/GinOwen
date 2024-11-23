package api

import (
	"GINOWEN/global"
	"GINOWEN/middlewares"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"GINOWEN/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type IPApi struct {
}

// UnLockIp 解除IP限制
// @Summary 解除IP限制
// @Description 解除IP限制
// @Tags IP
// @Accept json
// @Produce json
// @Param UnLockIpInput body request.UnLockIpInput true "解锁参数"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "IP not found in the blacklist"
// @Router /IP/UnLockIp [post]
func (IPApi) UnLockIp(c *gin.Context) {
	var request request.UnLockIpInput
	// 绑定请求数据
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.FailWithMessage("Invalid request", c)
		return
	}
	// 锁定黑名单以确保并发安全
	global.OWEN_LOCK.Lock()
	defer global.OWEN_LOCK.Unlock()
	middlewares.RemoveFromBlacklist(request.IP)
	utils.OkWithMessage(fmt.Sprintf("IP %s has been removed from blacklist", request.IP), c)
}

// AddBlackList 增加IP限制
// @Summary 增加IP限制
// @Description 增加IP限制
// @Tags IP
// @Accept json
// @Produce json
// @Param AddBlackListInput body request.AddBlackListInput true "增加黑名单参数"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid request"
// @Failure 400 {string} string "Invalid unlock time format"
// @Router /IP/AddBlackList [post]
func (IPApi) AddBlackList(c *gin.Context) {
	var request request.AddBlackListInput
	// 绑定请求数据
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.FailWithMessage("Invalid request", c)
		return
	}
	if request.IP == "::1" || request.IP == "localhost" || request.IP == "127.0.0.1" {
		utils.FailWithMessage("IP Invalid", c)
		return
	}
	unlockTime, err := utils.FormatLocalTime(request.Unlock)
	if err != nil {
		utils.FailWithMessage("Invalid time", c)
		return
	}
	// 将 IP 添加到黑名单
	global.OWEN_LOCK.Lock()
	defer global.OWEN_LOCK.Unlock()
	// 保存黑名单
	middlewares.SaveToBlacklist(request.IP, unlockTime)

	utils.OkWithMessage(fmt.Sprintf("IP %s has been added to the blacklist until %s", request.IP, unlockTime), c)
}

// GetBlackList 获取IP黑名单列表
// @Summary 获取IP黑名单列表
// @Description 获取当前IP黑名单的详细信息
// @Tags IP
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response{data=response.ShowBlackListDto,msg=string} "返回清单"
// @Failure 400 {string} string "Invalid request"
// @Router /IP/GetBlackList [get]
func (IPApi) GetBlackList(c *gin.Context) {
	var dto response.ShowBlackListDto
	dto.Items = middlewares.LoadBlacklist()
	utils.OkWithDetailed(dto, "获取成功", c)

}
