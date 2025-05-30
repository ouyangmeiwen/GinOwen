package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}
type GPIReceiveInput struct {
	ReceiveType string      `json:"receivetype" form:"receivetype"` //消息类型
	ReceiveData interface{} `json:"receivedata" form:"receivedata"` //消息体
}
