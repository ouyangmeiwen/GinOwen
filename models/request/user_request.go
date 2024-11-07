package request

type UserRequest struct {
	Name string `json:"name" binding:"required"`
}
