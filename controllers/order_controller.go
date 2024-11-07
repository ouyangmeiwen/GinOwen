package controllers

import (
	"GINOWEN/models"
	"GINOWEN/models/response"
	"net/http"

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
// @Param order body models.Order true "订单信息"
// @Success 201 {object} response.CommonResponse "成功创建的订单信息"
// @Failure 400 {object} response.CommonResponse "无效的请求"
// @Failure 500 {object} response.CommonResponse "服务器内部错误"
// @Router /api/orders/CreateOrder [post]
func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, response.CommonResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request data",
			Data:    nil,
		})
		return
	}
	if err := ServicesApp.orderService.CreateOrder(&order); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.CommonResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create order",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusCreated, response.CommonResponse{
		Code:    http.StatusCreated,
		Message: "Order created successfully",
		Data:    order,
	})
}

// GetOrders 获取所有订单
// @Summary 获取所有订单
// @Description 获取所有订单信息
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {object} response.CommonResponse "订单列表"
// @Failure 500 {object} response.CommonResponse "服务器内部错误"
// @Router /api/orders/GetOrders [get]
func (c *OrderController) GetOrders(ctx *gin.Context) {
	orders, err := ServicesApp.orderService.GetAllOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.CommonResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve orders",
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, response.CommonResponse{
		Code:    http.StatusOK,
		Message: "Orders retrieved successfully",
		Data:    orders,
	})
}
