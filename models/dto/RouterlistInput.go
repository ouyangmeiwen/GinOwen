package dto

type RouterlistInput struct {
	Id string `json:"id"`
}

type RouterlistData struct {
	Id         string   `json:"id"`
	RouterName string   `json:"routerName"`
	CaseNoList []string `json:"caseNoList"`
}
type RouterlistDto struct {
	FlyReadBaseResp
	Data []RouterlistData `json:"data"`
}
