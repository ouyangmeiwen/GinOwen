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
type InventoryInput struct {
	IsAll         bool     `form:"IsAll" json:"IsAll"`
	Workid        string   `form:"workid" json:"workid"`
	Triggers      int      `form:"triggers" json:"triggers"`
	Workname      string   `form:"workname" json:"workname"`
	Layerids      []string `form:"layerids" json:"layerids"`
	Devicetype    string   `form:"devicetype" json:"devicetype"`
	RobotId       string   `form:"robotId" json:"robotId"`
	RobotRouterId string   `form:"robotRouterId" json:"robotRouterId"`
}

type GetRobotRouterlistInput struct {
	Robotid string `form:"robotid" json:"robotid"`
}
