package dto

type InventoryCheckList struct {
	CaseNo   string `json:"caseNo"`   //"001B00000005001002",//层编号
	CaseName string `json:"caseName"` //"飞阅会议室 区域A 5排 B面第1列第2层"//层名称
}
type InventoryInputobj struct {
	DeviceType    string               `json:"deviceType"`    //目前只有摄像头  盘点类型：0:全景巡盘球机(球机摄像头),1:书架定点摄像头, 2:视觉盘点机器人,示例值(0)
	CheckAll      bool                 `json:"checkAll"`      //是否是全馆盘点 true-全馆盘点 false-区域盘点
	CheckType     string               `json:"checkType"`     // TIMING-定时任务 TRIGGER-触发任务 REALTIME-
	WorkType      string               `json:"workType"`      // CUSTOM_1 自定义全盘 CUSTOM_2 自定义部分盘点  TRIGGER-触发盘点  INTERVAL-间隔盘点 默认全盘
	WorkId        string               `json:"workId"`        //远望谷每次任务赋值
	WorkName      string               `json:"workName"`      //盘点任务名称
	CheckList     []InventoryCheckList `json:"checkList"`     //层集合
	RobotId       string               `json:"robotId"`       //机器人id,deviceType=2（视觉盘点机器人）时不能为空
	RobotRouterId string               `json:"robotRouterId"` //机器人路线id,deviceType=2（视觉盘点机器人）时不能为空
}
type InventoryInput struct {
	Obj   InventoryInputobj `json:"obj"` //
	Token string            `json:"token"`
}

type TaskCheckData struct {
	Msg       string   `json:"msg"`       //
	FailCases []string `json:"failCases"` //
}

// json 返回
type TaskCheckDto struct {
	FlyReadBaseResp
	Data TaskCheckData `json:"data"`
}

// 自定义对象返回非json返回
type InventoryDto struct {
	FlyReadBaseResp
	Errors []string
}
