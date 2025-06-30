package dto

type GetNotHitObj struct {
	Query_type  string `json:"query_type"`  // 2
	Query_limit int    `json:"query_limit"` // 10
}

type GetNotHitListInput struct {
	Container string       `json:"container"` //lcsinv
	Component string       `json:"component"` //shelf_book
	Service   string       `json:"service"`   //query_shelf_num
	Token     string       `json:"token"`     //token
	Obj       GetNotHitObj `json:"obj"`       //obj
}

type GetNotHitListDto struct {
	Case_name        string `json:"case_name"`        //为层位名字
	Case_no          string `json:"case_no"`          //为层位号
	Inventory_time   string `json:"inventory_time"`   //盘点时间
	UnConfidence_num string `json:"unConfidence_num"` //unConfidence_num
}

type QueryNotHitRankDto struct {
	FlyReadBaseResp
	Obj []GetNotHitListDto `json:"unConfidence_num"` //obj
}
