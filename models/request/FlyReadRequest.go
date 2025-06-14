package request

type HelloInput struct {
	Name string `form:"name" json:"name"` // 用户名
}
