package dto

type UploadLayersCase struct {
	Case_no       string `json:"case_no"`       //001A00000005001001 层号
	Case_name     string `json:"case_name"`     //第1层
	Shelf_no      string `json:"shelf_no"`      //001A00000005 架号
	Column_no     string `json:"column_no"`     //001A00000005001 列号
	Case_seq      int    `json:"case_seq"`      //1
	First_call_no string `json:"first_call_no"` //层首书
	Last_call_no  string `json:"last_call_no"`  //层尾书
}

type UploadLayersColumnInfo struct {
	Column_no     string `json:"column_no"`   //001A00000005001
	Column_name   string `json:"column_name"` //第1列
	Shelf_no      string `json:"shelf_no"`    //001A00000005
	Case_num      int    `json:"case_num"`    //2
	Column_seq    int    `json:"column_seq"`  //
	First_call_no string `json:"first_call_no"`
	Last_call_no  string `json:"last_call_no"`
}
type UploadLayersColumns struct {
	Column_info UploadLayersColumnInfo `json:"column_info"`
	Cases       []UploadLayersCase     `json:"cases"`
}
type UploadLayersShelf_info struct {
	Shelf_no      string `json:"shelf_no"`      //001A00000005
	Shelf_name    string `json:"shelf_name"`    //飞阅会议室 区域A 5排 A面
	Branch_no     string `json:"branch_no"`     //001
	Region_no     string `json:"region_no"`     //00001
	Case_num      int    `json:"case_num"`      //2  层数
	Column_num    int    `json:"column_num"`    // 2 列数
	Side          string `json:"side"`          //A
	First_call_no string `json:"first_call_no"` //
	Last_call_no  string `json:"last_call_no"`
}
type UploadLayersShelf struct {
	Shelf_info UploadLayersShelf_info `json:"shelf_info"`
	Columns    []UploadLayersColumns  `json:"columns"`
}
type UploadLayersObj struct {
	Shelfs []UploadLayersShelf `json:"shelfs"`
}
type UploadLayersInput struct {
	Container string          `json:"container"` //lcsinv
	Component string          `json:"component"` //shelf
	Service   string          `json:"service"`   //shelf_sync
	Token     string          `json:"token"`
	Obj       UploadLayersObj `json:"obj"`
}
type UploadLayersDto struct {
	FlyReadBaseResp
}
