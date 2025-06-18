package dto

type FlyReadBaseResp struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	Err     string `json:"err"`
	Message string `json:"message"`
}

type GetTokenDto2Data struct {
	AccessToken string `json:"accessToken"`
}
type GetTokenDto2 struct {
	FlyReadBaseResp
	Data GetTokenDto2Data `json:"data"`
}
