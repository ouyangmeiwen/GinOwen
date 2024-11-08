package controllers

import (
	"GINOWEN/models"
	"GINOWEN/utils"

	"github.com/gin-gonic/gin"
)

// OrderController 订单控制器
type OrderController struct {
}

// CreateOrder 创建订单
// @Summary 创建订单
// @Description 创建新的订单并返回订单信息
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body models.TestOrder true "订单信息"
// @Success 201 {object} utils.Response{data=models.TestOrder,msg=string} "成功创建的订单信息"
// @Failure 400 {object} utils.Response{msg=string} "无效的请求"
// @Failure 500 {object} utils.Response{msg=string} "服务器内部错误"
// @Router /api/services/app/order/CreateOrder [post]
func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.TestOrder
	if err := ctx.ShouldBindJSON(&order); err != nil {
		utils.FailWithMessage("Invalid request data", ctx)
		return
	}
	if err := ServicesApp.orderService.CreateOrder(&order); err != nil {
		utils.FailWithMessage("Failed to create", ctx)
		return
	}
	utils.OkWithDetailed(order, "created successfully", ctx)
}

// GetOrders 获取所有订单
// @Summary 获取所有订单
// @Description 获取所有订单信息
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.Response{data=[]models.TestOrder,msg=string} "订单列表"
// @Failure 500 {object} utils.Response{msg=string} "服务器内部错误"
// @Router /api/services/app/order/GetOrders [get]
func (c *OrderController) GetOrders(ctx *gin.Context) {
	orders, err := ServicesApp.orderService.GetAllOrders()
	if err != nil {
		utils.FailWithMessage("Failed to retrieve", ctx)
		return
	}
	utils.OkWithDetailed(orders, "get successfully", ctx)
}
