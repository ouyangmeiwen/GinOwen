package dto

type InventoryHisObj struct {
	Type string `json:"type"` //CURRENT-当前盘点 LASTEST-最近一次完成的盘点
}

type InventoryHisInput struct {
	Obj   InventoryHisObj `json:"obj"`
	Token string          `json:"token"`
}
type InventoryHisRespObj struct {
	CreateTime  string `json:"createTime"`  //"2024-05-03 17:44:58",//创建时间
	UpdateTime  string `json:"updateTime"`  //"2024-05-03 17:47:58",//更新时间
	WorkId      string `json:"workId"`      // "706441161815887872",//盘点id
	CheckType   string `json:"checkType"`   //"REALTIME",//任务类型
	CheckAll    bool   `json:"checkAll"`    //是否全馆
	DeviceType  string `json:"deviceType"`  //"CAMERA",//盘点设备类型
	Status      string `json:"status"`      //"EXPORTED",//盘点任务状态 INIT-创建 COLLECTING-采集 COMPUTED-盘点 EXPORTED-盘点完成 FAILED-失败
	TaskAllNum  int    `json:"taskAllNum"`  //盘点总个数
	FinishNum   int    `json:"finishNum"`   //完成个数
	BookCount   int    `json:"bookCount"`   //已盘点图书
	AssertCount int    `json:"assertCount"` //馆藏总数
}

type ChecklatestDto struct {
	FlyReadBaseResp
	Data []InventoryHisRespObj `json:"data"`
}

type InventoryHisDto struct {
	FlyReadBaseResp
	Obj InventoryHisRespObj `json:"Obj"`
}
