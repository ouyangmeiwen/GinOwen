package request

type UnLockIpInput struct {
	IP string `json:"ip"`
}

type AddBlackListInput struct {
	IP     string `json:"ip"`
	Unlock string `json:"unlock_time"` // 解锁时间，格式："yyyy-mm-dd hh:mm:ss"
}

type ShowBlackListInput struct {
}
