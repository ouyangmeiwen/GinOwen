package dto

type UploadBookLocInput struct {
	CaseNo     string   `json:"caseNo"`
	AssetIdSet []string `json:"assetIdSet"`
}
type UploadBookInfoLocDto struct {
	FlyReadBaseResp
}
