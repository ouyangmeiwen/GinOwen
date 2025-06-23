package dto

type RobotlistInput struct {
}

type RobotlistData struct {
	Id        string `json:"id"`
	RobotName string `json:"robotName"`
}
type RobotlistDto struct {
	FlyReadBaseResp
	Data []RobotlistData `json:"data"`
}
