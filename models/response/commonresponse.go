package response

// CommonResponse 定义统一返回结构体
type CommonResponse struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回的数据，可以是任何类型
}
