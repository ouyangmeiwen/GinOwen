package request

type UploadRequest struct {
	Name        string `form:"name" json:"name" binding:"required"`        // 文件名称
	Description string `form:"description" json:"description" binding:"-"` // 文件描述（可选）
	Cover       bool   `form:"cover" json:"cover"`                         // 文件名称
}
