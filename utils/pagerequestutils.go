package utils

type PagedResultRequest struct {
	SkipCount      int `json:"SkipCount" form:"SkipCount"  default:"0"  validate:"min=-1,max=99999"`            //跳过多少数量
	MaxResultCount int `json:"MaxResultCount" form:"MaxResultCount"  default:"100"  validate:"min=1,max=99999"` //获取多少条结果
}
