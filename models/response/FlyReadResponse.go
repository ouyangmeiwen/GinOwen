package response

type HelloResp struct {
	Message string `json:"message"` // 响应消息
	Status  int    `json:"status"`  // 响应状态码
	Success bool   `json:"success"` // 是否成功
}
