package dto

type BookItem struct {
	Asset_id        string `json:"asset_id"`        //资产号
	Barcode         string `json:"barcode"`         //条码
	Call_no         string `json:"call_no"`         //索书号
	Title           string `json:"title"`           //书名
	Isbn            string `json:"isbn"`            //ISBN
	Author          string `json:"author"`          //作者
	Publisher       string `json:"publisher"`       //出版社
	Pubdate         string `json:"pubdate"`         //出版时间
	Clc             string `json:"clc"`             //分类号
	Lib_location    string `json:"lib_location"`    //馆藏地
	Status          string `json:"status"`          //状态
	Collection_time string `json:"collection_time"` //入藏时间
	Shelf_time      string `json:"shelf_time"`      //上架时间
}
type BookObj struct {
	Books []BookItem `json:"books"` //书籍列表
}
type UploadBookInfoInput struct {
	Container string  `default:"lcsinv" json:"container"`        //容器名称
	Component string  `default:"shelf" json:"component"`         //组件名称
	Service   string  `default:"collection_sync" json:"service"` //服务名称
	Token     string  `json:"token"`                             //飞阅Token
	Obj       BookObj `json:"obj"`                               //书籍对象
}

type UploadBookInfoResp struct {
	FlyReadBaseResp
	obj string
}
