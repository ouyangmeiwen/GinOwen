package dto

type GetOcrImgsData struct {
	CaseNo    string   `json:"caseNo"`    // 层位编码
	ImageList []string `json:"imageList"` // 图片列表bs64
}

type GetOcrImgsDto struct {
	FlyReadBaseResp
	Data GetOcrImgsData `json:"data"` // 图片列表
}
type GetOcrImgsInput struct {
	CaseNo  string `json:"caseNo"`  // 层位编码
	AssetId string `json:"assetId"` // 资产ID
}
