package request

// LoginRequest 用于接收前端发送的登录请求数据
type LoginRequest struct {
	TenantId int    `json:"tenantid" binding:"required" default:"5325"`
	Username string `json:"username" binding:"required" default:"admin"`
	Password string `json:"password" binding:"required" default:"123456"`
}
type LoginOutRequest struct {
}

// RegisterRequest 用于接收前端注册请求的数据
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID   uint   `json:"role_id" binding:"required"`
}

type LocalLoginRequest struct {
	Token string `form:"token" json:"token" binding:"required"`
	// #get 请求一定要带上form:"token" json:"token" TODO
}

type LocalOutRequest struct {
	Token string `json:"token" binding:"required"`
}
