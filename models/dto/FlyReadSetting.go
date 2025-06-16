package dto

type FlyReadSetting struct {
	FlyReadEnable    string //是否开启飞阅接口1 则开启
	FlyReadIp        string //飞阅接口IP
	FlyReadPort      string //飞阅接口端口
	FlyAppid         string //飞阅id
	FlyReadAppSecret string //飞阅秘钥
	FlyMode          string //飞阅OCR模式 0:OCR;1:OCR+条码 默认ocr
	FlyAddLoc        string // 自动上架 默认自动上架 0:则不上架 1:自动上架  默认自动上架
	FlyLocType       string //定位模式0 强制定位 1首书定位 2无定位 默认强制定位
	FlyLostDay       string //疑似丢失判定天数
	FlyStartTime     string //每天开始时间
	FlyEndTime       string //每天结束时间
	FlyTenant        string //飞阅租户
	FlyShowOff       string //是否显示离架数据
	FlyDeviceType    string //任务类型
	FlyTempLoc       string //0无定位模式下  是否启用临时定位 1层提供定位  2面提供定位
	FlyNewLocMode    string //获取定位方式 新 1无定位 2自动定位 3配置定位 4RFID首书定位  5 RFID强制定位
	RowShape         string //书架形状 U型还是同序 0或者空则默认 U型，1则表示同序
}
