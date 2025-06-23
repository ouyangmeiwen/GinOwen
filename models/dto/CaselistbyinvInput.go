package dto

type CaselistbyinvInput struct {
	DeviceType string `json:"deviceType"` //盘点类型：0:全景巡盘球机(球机摄像头),1:书架定点摄像头
}
type CaselistbyinvDto struct {
	FlyReadBaseResp
	Data []string `json:"data"`
}
