package request

type SendMqMsgInput struct {
	DataType int    `json:"datatype" form:"datatype"` // 消息类型 0则不转换json 其他类型则必须保证jsonbody 是个合法的json 对象
	JsonBody string `json:"jsonbody" form:"jsonbody"` //消息体
}
