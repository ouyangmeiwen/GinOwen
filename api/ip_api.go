package api

import (
	"GINOWEN/global"
	"GINOWEN/middlewares"
	"GINOWEN/models/request"
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

	// 检查 IP 是否在黑名单中
	if _, exists := middlewares.Blacklist[request.IP]; exists {
		// 从黑名单中移除该 IP
		delete(middlewares.Blacklist, request.IP)
		// 保存更新后的黑名单
		middlewares.SaveBlacklist()
		utils.OkWithMessage(fmt.Sprintf("IP %s has been removed from blacklist", request.IP), c)
	} else {
		// 如果 IP 不在黑名单中，返回 404 错误
		utils.FailWithMessage(fmt.Sprintf("IP %s is not in the blacklist", request.IP), c)
	}
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
	unlockTime, err := utils.FormatLocalTime(request.Unlock)
	if err != nil {
		utils.FailWithMessage("Invalid time", c)
		return
	}
	// 将 IP 添加到黑名单
	global.OWEN_LOCK.Lock()
	middlewares.Blacklist[request.IP] = unlockTime
	global.OWEN_LOCK.Unlock()
	// 保存黑名单
	middlewares.SaveBlacklist()
	utils.OkWithMessage(fmt.Sprintf("IP %s has been added to the blacklist until %s", request.IP, unlockTime), c)
}
