package request

import "GINOWEN/utils"

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

type CreatWorkInput struct {
	TaskName        string   `form:"taskName" json:"taskName"`               // 任务名称
	WorkTime        string   `form:"workTime" json:"workTime"`               // 任务时间
	LayerIds        []string `form:"layerIds" json:"layerIds"`               // 层位ID集合
	DeviceType      string   `form:"deviceType" json:"deviceType"`           // 设备类型
	RobotId         string   `form:"robotId" json:"robotId"`                 // 机器人ID
	RobotName       string   `form:"robotName" json:"robotName"`             // 机器人名称
	RobotRouterId   string   `form:"robotRouterId" json:"robotRouterId"`     // 机器人路由ID
	RobotRouterName string   `form:"robotRouterName" json:"robotRouterName"` // 机器人路由名称
	WorkTimes       []string `form:"workTimes" json:"workTimes"`             // 工作时间集合
}
type UpdateWorkInput struct {
	WorkId          string   `form:"workId" json:"workId"`                   // 任务ID
	WorkName        string   `form:"workName" json:"workName"`               // 任务名称
	WorkTime        string   `form:"workTime" json:"workTime"`               // 任务时间 2006-01-02 15:04:05
	LayerIds        []string `form:"layerIds" json:"layerIds"`               // 层
	DeviceType      string   `form:"deviceType" json:"deviceType"`           // 设备类型
	RobotId         string   `form:"robotId" json:"robotId"`                 // 机器人ID
	RobotName       string   `form:"robotName" json:"robotName"`             // 机器人名称
	RobotRouterId   string   `form:"robotRouterId" json:"robotRouterId"`     // 机器人路由ID
	RobotRouterName string   `form:"robotRouterName" json:"robotRouterName"` // 机器人路由名称
}
type WorkListInput struct {
	utils.PagedResultRequest
	WorkId       string `form:"WorkId" json:"WorkId"`
	WorkName     string `form:"WorkName" json:"WorkName"`
	TriggerSatus int    `form:"TriggerSatus" json:"TriggerSatus"` //间隔盘点0  自定义盘点1  触发盘点2
	TaskStatus   int    `form:"TaskStatus" json:"TaskStatus"`     //任务初始化0 盘点中1  成功2 失败3
	DeviceType   string `form:"DeviceType" json:"DeviceType"`     //目前只有摄像头  盘点类型：0:全景巡盘球机(球机摄像头),1:书架定点摄像头, 2:视觉盘点机器人,示例值(0)
}
type DeleteWorkInput struct {
	Workid string `form:"workid" json:"workid"`
}
type DetailListInput struct {
	utils.PagedResultRequest
	WorkId         string `form:"WorkId" json:"WorkId"`
	LayerId        string `form:"LayerId" json:"LayerId"`
	InventoryState int    `form:"InventoryState" json:"InventoryState"`
	Barcode        string `form:"Barcode" json:"Barcode"`
	Title          string `form:"Title" json:"Title"`
	CallNo         string `form:"CallNo" json:"CallNo"`
	ISBN           string `form:"ISBN" json:"ISBN"`
	ShowOff        string `form:"ShowOff" json:"ShowOff"`
}
