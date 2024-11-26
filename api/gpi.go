package api

import (
	"GINOWEN/extend/rabbitmqextend"
	"GINOWEN/extend/websocketextend"
	"GINOWEN/middlewares"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"GINOWEN/utils"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
)

type GPIApi struct {
}

// GPI接口 GPI接口
// @Summary GPI接口
// @Description GPI接口
// @Tags GPI
// @Accept json
// @Produce json
// @Param req body request.GPIReceiveInput true "消息入参 receivetype标记管道,receivedata标记管道参数{}"
// @Success 200 {object} utils.Response{data=interface{},msg=string} "返参"
// @Security BearerAuth
// @Router /GPI/GPIReceive [post]
func (gpi GPIApi) GPIReceive(c *gin.Context) {
	var req request.GPIReceiveInput
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage("参数无效", c)
		return
	}
	switch req.ReceiveType {
	case "RabbitMQ.Send":
		gpi.handle_RabbitMQ_Send(c, req)
	case "WebSocket.Send":
		gpi.handle_WebSocket_Send(c, req)
	case "IP.GetBlackList":
		gpi.handle_IP_GetBlackList(c)
	case "Sysauditlmslog.QueryLmsLog":
		gpi.handle_Sysauditlmslog_QueryLmsLog(c, req)
	default:
		utils.FailWithMessage("ReceiveType不支持", c)
	}
}

func (GPIApi) handle_RabbitMQ_Send(c *gin.Context, req request.GPIReceiveInput) {
	receiveBytes, err := json.Marshal(req.ReceiveData)
	if err != nil {
		utils.FailWithMessage("Error Marshal ReceiveData", c)
		return
	}
	var input request.SendMqMsgInput
	err = json.Unmarshal(receiveBytes, &input)
	if err != nil {
		utils.FailWithMessage("Error Unmarshal SendMqMsgInput", c)
		return
	}
	var data rabbitmqextend.Data
	data.DataType = input.DataType
	bodyBytes, err := json.Marshal(input.JsonBody)
	if err != nil {
		utils.FailWithMessage("Error Marshal JsonBody", c)
		return
	}
	data.Body = json.RawMessage(bodyBytes)
	rabbitmqextend.InstancePublisher.PublishData(input.RoutingKey, data)
	utils.OkWithMessage("执行完成", c)
}

func (GPIApi) handle_WebSocket_Send(c *gin.Context, req request.GPIReceiveInput) {
	data, ok := req.ReceiveData.(string)
	if !ok {
		utils.FailWithMessage("ReceiveData is not a string", c)
		return
	}
	parts := strings.SplitN(data, ":", 2) //限制分割为2个字符串
	websocketextend.Instance.SendToClient(parts[0], []byte(parts[1]))
	utils.OkWithMessage("执行完成", c)
}

func (GPIApi) handle_IP_GetBlackList(c *gin.Context) {
	var dto response.ShowBlackListDto
	dto.Items = middlewares.LoadBlacklist()
	utils.OkWithDetailed(dto, "查询成功", c)
}
func (GPIApi) handle_Sysauditlmslog_QueryLmsLog(c *gin.Context, req request.GPIReceiveInput) {
	receiveBytes, err := json.Marshal(req.ReceiveData)
	if err != nil {
		utils.FailWithMessage("Error Marshal ReceiveData", c)
		return
	}
	var input request.QueryLmsInput
	err = json.Unmarshal(receiveBytes, &input)
	if err != nil {
		utils.FailWithMessage("Error Unmarshal QueryLmsInput", c)
		return
	}
	list, err := ServicesGroup.sysauditlmslogService.QueryLmsLog(input)
	if err != nil {
		utils.FailWithMessage("获取失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(list, "获取成功", c)
}
