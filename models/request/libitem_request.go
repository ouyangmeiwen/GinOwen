package request

type ImportExcelInput struct {
	Path         string `default:"d:/4.xlsx" binding:"required" validate:"min=3,max=100"` //excel 服务器本地路径
	Title        string `default:"题名"`                                                    //题名
	Author       string `default:"作者"`                                                    //作者
	Tid          string `default:"tid"`                                                   //tid
	CallNo       string `default:"索书号"`                                                   //索书号
	ISBN         string `default:"ISBN"`                                                  //ISBN
	CatalogCode  string `default:"分类号"`                                                   //分类号
	Publisher    string `default:"出版社"`                                                   //出版社
	PubDate      string `default:"出版日期"`                                                  //出版日期
	Price        string `default:"单价"`                                                    //单价
	Pages        string `default:"页数"`                                                    //页数
	Barcode      string `default:"条码号"`                                                   //条码号
	Locationname string `default:"当前馆藏地点"`                                                //当前馆藏地点
	Tenantid     int    `default:"1" binding:"required" validate:"min=1,max=99999"`       //租户
}
