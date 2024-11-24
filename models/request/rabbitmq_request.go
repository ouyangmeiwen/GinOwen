package request

type SendMqMsgInput struct {
	DataType string      `json:"datatype" form:"datatype"` // 消息类型
	JsonBody interface{} `json:"jsonbody" form:"jsonbody"` //消息体
}
