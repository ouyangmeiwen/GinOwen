package dto

type Brandobj struct {
	Branch_no   string `json:"branch_no"`   //分馆编号
	Branch_name string `json:"branch_name"` //分馆名字
}

type UploadBranchInput struct {
	Container string   `json:"container"` //lcsinv
	Component string   `json:"component"` //shelf
	Service   string   `json:"service"`   //branch_sync
	Token     string   `json:"token"`
	Obj       Brandobj `json:"obj"`
}
type UploadBranchResp struct {
	FlyReadBaseResp
	Obj string `json:"obj"`
}
