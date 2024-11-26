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
// @Param req body request.GPIReceiveInput true "入参"
// @Success 200 {object} utils.Response{data=interface{},msg=string} "返参"
// @Security BearerAuth
// @Router /GPI/GPIReceive [post]
func (gpi GPIApi) GPIReceive(ctx *gin.Context) {
	var req request.GPIReceiveInput
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		utils.FailWithMessage("参数无效", ctx)
		return
	}
	switch req.ReceiveType {
	case "RabbitMQ.Send":
		receiveBytes, err := json.Marshal(req.ReceiveData)
		if err != nil {
			utils.FailWithMessage("Error Marshal ReceiveData", ctx)
			return
		}
		var input request.SendMqMsgInput
		err = json.Unmarshal(receiveBytes, &input)
		if err != nil {
			utils.FailWithMessage("Error Unmarshal SendMqMsgInput", ctx)
			return
		}
		var data rabbitmqextend.Data
		data.DataType = input.DataType
		bodyBytes, err := json.Marshal(input.JsonBody)
		if err != nil {
			utils.FailWithMessage("Error Marshal JsonBody", ctx)
			return
		}
		data.Body = json.RawMessage(bodyBytes)
		rabbitmqextend.InstancePublisher.PublishData(input.RoutingKey, data)
		utils.OkWithMessage("执行完成", ctx)
	case "WebSocket.Send":
		data, ok := req.ReceiveData.(string)
		if !ok {
			utils.FailWithMessage("ReceiveData is not a string", ctx)
			return
		}
		parts := strings.SplitN(data, ":", 2) //限制分割为2个字符串
		websocketextend.Instance.SendToClient(parts[0], []byte(parts[1]))
		utils.OkWithMessage("执行完成", ctx)
	case "IP.GetBlackList":
		var dto response.ShowBlackListDto
		dto.Items = middlewares.LoadBlacklist()
		utils.OkWithDetailed(dto, "查询成功", ctx)

	case "Sysauditlmslog.QueryLmsLog":
		receiveBytes, err := json.Marshal(req.ReceiveData)
		if err != nil {
			utils.FailWithMessage("Error Marshal ReceiveData", ctx)
			return
		}
		var input request.QueryLmsInput
		err = json.Unmarshal(receiveBytes, &input)
		if err != nil {
			utils.FailWithMessage("Error Unmarshal QueryLmsInput", ctx)
			return
		}
		list, err := ServicesGroup.sysauditlmslogService.QueryLmsLog(input)
		if err != nil {
			utils.FailWithMessage("获取失败!"+err.Error(), ctx)
			return
		}
		utils.OkWithDetailed(list, "获取成功", ctx)
	default:
		utils.FailWithMessage("ReceiveType不支持", ctx)
	}
}
