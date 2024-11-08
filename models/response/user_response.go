// models/response/user.go
package response

// UserDto 用户返回的DTO
type UserDto struct {
    ID   uint   `json:"id"`   // 用户ID
    Name string `json:"name"` // 用户名
}
