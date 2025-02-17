package request

type SendMqMsgInput struct {
	DataType   string      `json:"datatype" form:"datatype"`                                       // 消息类型
	RoutingKey string      `json:"routingkey" form:"routingkey" default:"default_routingKey.init"` // 路由键
	JsonBody   interface{} `json:"jsonbody" form:"jsonbody"`                                       //消息体
}
