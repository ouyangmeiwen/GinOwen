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

type InventoryHisInput struct {
	IsHistory bool `json:"isHistory" form:"isHistory"` // 是否历史盘点
}

type InventoryListInput struct {
	PageIndex int    `form:"pageIndex" json:"pageIndex"`                  // 页码
	PageSize  int    `form:"pageSize" json:"pageSize"`                    // 页大小
	DtStart   string `form:"dtStart" json:"dtStart" default:"2025-01-01"` // 开始时间
	DtEnd     string `form:"dtEnd" json:"dtEnd" default:"2025-01-10"`     // 结束时间
}
type SetBussinessInput struct {
	Addloc     string `form:"addloc" json:"addloc"`                // 0:则不上架 1:自动上架  默认自动上架
	Newlocmode string `form:"newlocmode" json:"newlocmode"`        // 1无定位 2自动定位 3配置定位 4RFID首书定位  5 RFID强制定位
	Lostday    string `form:"lostday" json:"lostday"`              // 疑似丢失判定天数
	Mode       string `binding:"required" form:"mode" json:"mode"` // 0:OCR;1:OCR+条码 默认ocr
	ShowOff    string `form:"showOff" json:"showOff"`              // 是否显示离架数据
	Shape      string `form:"shape" json:"shape"`                  // 书架形状 U型还是同序 0或者空则默认 U型，1则表示同序
}
type GetInventorySetInput struct {
}

type GetEnableRowInput struct {
	IsQueryAll bool   `form:"isQueryAll" json:"isQueryAll"` // 是否查询全部行
	Devicetype string `form:"devicetype" json:"devicetype"` // 设备类型
}
type GetRobotlistInput struct {
}
type GetCaseCodeImageInput struct {
	Layer_codes []string `form:"layer_codes" json:"layer_codes"` // 层位编码集合
}
type GetOcrImgsInput struct {
	LayerCode   string `form:"layerCode" json:"layerCode"`     // 层位编码
	ItemBarcode string `form:"itemBarcode" json:"itemBarcode"` // 物品条码
}
type InventorySetInput struct {
	SysStartTime       string `form:"SysStartTime" json:"SysStartTime"`                // 系统开始时间编码
	SysEndTime         string `form:"SysEndTime" json:"SysEndTime"`                    // 系统结束时间编码
	IsEnable           bool   `form:"IsEnable" json:"IsEnable"`                        // 是否启用
	InventoryStartDate string `form:"InventoryStartDate" json:"InventoryStartDate"`    // 盘点开始时间
	InventoryEndDate   string `form:"InventoryEndDate" json:"InventoryEndDate"`        // 盘点结束时间
	Interval           int    `form:"Interval" json:"Interval"`                        // 盘点间隔时间
	DeviceType         string `binding:"required" form:"DeviceType" json:"DeviceType"` // 设备类型 目前只有摄像头  盘点类型：0:全景巡盘球机(球机摄像头),1:书架定点摄像头
}
