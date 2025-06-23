package dto

type FlyReadLayer struct {
	Id             string
	Code           string
	Name           string
	Side           string
	LayerNo        int
	LayerFaultNums int
	CallNo         string
	PreCallNo      string
	ItemBarcode    string
}
type FlyReadShelf struct {
	Id               string
	Code             string
	Name             string
	ShelfNo          int
	Side             string
	StructId         string
	Layers           []FlyReadLayer
	ShelfFaultNum    int
	IsBindShelfPoint bool
}

type FlyReadRow struct {
	Id           string
	Code         string
	Name         string
	RowNo        int
	RowType      int
	RowUsageType int
	Shelfs       []FlyReadShelf
	RowFaultNum  int
}
type FlyRouter struct {
	RouterId   string
	RouterName string
	Items      []FlyReadRow
}

type GetEnableRowDto struct {
	Items []FlyReadRow
}
