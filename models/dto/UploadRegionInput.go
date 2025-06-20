package dto

type RegionObj struct {
	Branch_no   string `json:"branch_no"`   // 关联分馆编号
	Region_no   string `json:"region_no"`   //区域编号
	Region_name string `json:"region_name"` //区域名称
}

type UploadRegionInput struct {
	Container string    `json:"container"` //lcsinv
	Component string    `json:"component"` //shelf
	Service   string    `json:"service"`   //region_sync
	Token     string    `json:"token"`
	Obj       RegionObj `json:"obj"`
}
type UploadRegionResp struct {
	FlyReadBaseResp
	Obj string `json:"obj"`
}
