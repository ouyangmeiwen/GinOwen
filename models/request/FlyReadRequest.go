package request

type HelloInput struct {
	Name string `form:"name" json:"name"` // 用户名
}
type GetFlyTokenInput struct {
	IsForceRefresh bool `form:"isForceRefresh" json:"isForceRefresh"` // 是否强制刷新
}

type UploadLibItemInput struct {
	Barcodes []string `form:"barcodes" json:"barcodes"` // 条码集合
}

type UploadTenantInput struct {
	Tenantid int `form:"tenantid" json:"tenantid"` // 条码集合
}

type UploadStructInput struct {
	Structid string `form:"structid" json:"structid"` //结构ID或者结构编码
}

type UploadLibItemLocInput struct {
	Layercode []string `form:"layercode" json:"layercode"` // 层架集合
}
type UploadRowInput struct {
	RowNos []int `form:"rownos" json:"rownos"` // 架号
}
