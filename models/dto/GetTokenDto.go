package dto

type FlyReadBase struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type GetTokenDto2Data struct {
	AccessToken string `json:"accessToken"`
}
type GetTokenDto2 struct {
	FlyReadBase
	Data GetTokenDto2Data `json:"data"`
}
