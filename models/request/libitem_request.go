package request

type ImportExcelByNameInput struct {
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
type ImportExcelByIndexInput struct {
	Path         string `default:"d:/4.xlsx" binding:"required" validate:"min=3,max=100"` //excel 服务器本地路径
	Title        int    `default:"-1"`                                                    //题名
	Author       int    `default:"-1"`                                                    //作者
	Tid          int    `default:"-1"`                                                    //tid
	CallNo       int    `default:"-1"`                                                    //索书号
	ISBN         int    `default:"-1"`                                                    //ISBN
	CatalogCode  int    `default:"-1"`                                                    //分类号
	Publisher    int    `default:"-1"`                                                    //出版社
	PubDate      int    `default:"-1"`                                                    //出版日期
	Price        int    `default:"-1"`                                                    //单价
	Pages        int    `default:"-1"`                                                    //页数
	Barcode      int    `default:"-1"`                                                    //条码号
	Locationname int    `default:"-1"`                                                    //当前馆藏地点
	Tenantid     int    `default:"1" binding:"required" validate:"min=1,max=99999"`       //租户
}
