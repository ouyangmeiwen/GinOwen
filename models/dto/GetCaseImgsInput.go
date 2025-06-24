package dto

type GetCaseImgsInput struct {
	CaseNo string `json:"caseNo"` // 层位编码
	Type   string `json:"type"`   // 图片类型：1采集图片 2 ocr处理图片
}

type GetCaseImgsItem struct {
	CaseNo    string   `json:"caseNo"`    // 层位编码
	ImageList []string `json:"imageList"` // 图片列表bs64
}

type GetCaseImgsDto struct {
	FlyReadBaseResp
	Data []GetCaseImgsItem `json:"data"` // 图片列表
}
