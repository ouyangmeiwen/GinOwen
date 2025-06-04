package request

// get 方式 大小写是敏感的
type QueryRowInput struct {
	RowNo      *int    `form:"row_no" json:"row_no"`           // 书架号
	StructCode *string `form:"struct_code" json:"struct_code"` //馆藏地编码
}
