package request

// LoginRequest 用于接收前端发送的登录请求数据
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginOutRequest struct {
}

// RegisterRequest 用于接收前端注册请求的数据
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID   uint   `json:"role_id" binding:"required"`
}
