package dto

type GetShelfNumObj struct {
	Query_type string `json:"query_type"` //1
}

type GetShelfNumInput struct {
	Container string         `json:"container"` //lcsinv
	Component string         `json:"component"` //shelf_book
	Service   string         `json:"service"`   //query_shelf_num
	Token     string         `json:"token"`     //token
	Obj       GetShelfNumObj `json:"obj"`       //obj
}

type QueryShelfNumDto struct {
	FlyReadBaseResp
	Obj GetShelfNumDto `json:"obj"`
}

type GetShelfNumDto struct {
	Seg_num          int `json:"seg_num"`          // 书籍数量(总书脊识别量)
	Ocr_num          int `json:"ocr_num"`          // ocr识别数量（总ocr识别量）  目前不用
	Match_num        int `json:"match_num"`        // 匹配数量(总匹配量)
	Confidence_num   int `json:"confidence_num"`   // 总命中数量
	Unconfidence_num int `json:"unconfidence_num"` // 总未命中数量
}
