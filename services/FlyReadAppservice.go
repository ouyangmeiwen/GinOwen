package services

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"fmt"
	"strconv"
)

type FlyReadAppService struct {
}

func (f *FlyReadAppService) Hello(req request.HelloInput) (resp response.HelloResp, err error) {
	return ManagerGroup.frymanager.Hello(req), nil
}

// 上传图书
func (f *FlyReadAppService) UploadLibItem(req request.UploadLibItemInput, tenantid int) (resp response.UploadLibItemResp, err error) {
	if len(req.Barcodes) == 0 {
		return resp, fmt.Errorf("条码集合不能空")
	}

	var items []models.Libitem
	err = global.OWEN_DB.Model(&models.Libitem{}).Where("TenantId=? and IsDeleted=0", tenantid).Where(" Barcode in ?", req.Barcodes).Find(&items).Error
	if err != nil {
		return resp, err
	}
	if len(items) == 0 {
		return resp, fmt.Errorf("条码查询图书为空不允许推送！")
	}
	var result bool
	var msg string
	result, msg, err = ManagerGroup.frymanager.UploadLibItem(items, tenantid)
	if err != nil {
		return resp, err
	}
	fmt.Println(result)
	fmt.Println(msg)
	resp.Success = result
	return resp, nil

}

// 上报租户
func (f *FlyReadAppService) UploadTenant(input request.UploadTenantInput) (resp response.UploadTenantDto, err error) {

	var tenant models.Abptenant
	err = global.OWEN_DB.Model(&models.Abptenant{}).Where("IsDeleted=0").Where("id=?", input.Tenantid).First(&tenant).Error
	if err != nil {
		return resp, err
	}
	if tenant.ID == 0 {
		return resp, fmt.Errorf("error tenantid")
	}
	var bol bool
	bol, err = ManagerGroup.frymanager.UploadTenant(tenant)
	if err != nil {
		return resp, err
	}
	resp.Success = bol
	return resp, nil
}

// 上报结构
func (f *FlyReadAppService) UploadStruct(input request.UploadStructInput, tenantid int) (resp response.UploadStructDto, err error) {
	if len(input.Structid) < 6 {
		return resp, fmt.Errorf("Structid长度不够")
	}
	buildNo, _ := strconv.Atoi(input.Structid[0:2])
	floorNo, _ := strconv.Atoi(input.Structid[2:4])
	roomNo, _ := strconv.Atoi(input.Structid[4:6])

	var struct_one models.Libstruct
	err = global.OWEN_DB.Model(&models.Libstruct{}).Where("TenantId=? and IsDeleted=0", tenantid).Where("Id=?  or (BuildNo=? and FloorNo=? and RoomNo=?)", input.Structid, buildNo, floorNo, roomNo).First(&struct_one).Error
	if err != nil {
		return resp, err
	}
	if struct_one.ID == "" {
		return resp, fmt.Errorf("id not found")
	}
	var bol bool
	bol, err = ManagerGroup.frymanager.UploadStruct(struct_one, tenantid)
	if err != nil {
		return resp, err
	}
	resp.Success = bol
	return resp, nil
}

// 推送图书定位
func (f *FlyReadAppService) UploadLibItemLoc(input request.UploadLibItemLocInput, tenantid int) (resp response.UploadLibItemLocDto, err error) {
	if len(input.Layercode) <= 0 {
		return resp, fmt.Errorf("layercode集合不能空")
	}
	var lst []models.Libitemlocinfo
	err = global.OWEN_DB.Model(&models.Libitemlocinfo{}).Where("IsDeleted=0 and LayerCode in ? and TenantId=?", input.Layercode, tenantid).Find(&lst).Error
	if err != nil {
		return resp, err
	}
	if len(lst) <= 0 {
		return resp, fmt.Errorf("layercode  is invalid")
	}
	var bol bool
	bol, err = ManagerGroup.frymanager.UploadLibItemLoc(lst, tenantid)
	if err != nil {
		return resp, err
	}
	resp.Success = bol
	return resp, nil
}

// 层架推送
func (f *FlyReadAppService) UploadRow(input request.UploadRowInput, tenantid int) (resp response.UploadRowDto, err error) {
	islimitquery := false
	if len(input.RowNos) > 0 {
		islimitquery = true
	}
	var rows []models.Librow
	row_query := global.OWEN_DB.Model(&models.Librow{}).Where("IsDeleted=0 and TenantId=?", tenantid)
	if islimitquery {
		row_query = row_query.Where("RowNo in ?", input.RowNos)
	}
	err = row_query.Find(&rows).Error
	if err != nil {
		return resp, err
	}
	if len(rows) <= 0 {
		return resp, fmt.Errorf("no data row")
	}
	var rowids []string
	for _, row := range rows {
		rowids = append(rowids, row.ID)
	}
	var shelfs []models.Libshelf
	shelfs_map := make(map[string][]models.Libshelf) //按照架分号的列
	shelf_query := global.OWEN_DB.Model(&models.Libshelf{}).Where("IsDeleted=0 and IsEnable=1 and TenantId=?", tenantid)
	if islimitquery {
		shelf_query = shelf_query.Where("RowIdentity in ?", rowids)
	}
	err = shelf_query.Find(&shelfs).Error
	if err != nil {
		return resp, err
	}
	if len(shelfs) <= 0 {
		return resp, fmt.Errorf("no data shelfs")
	}
	var shelfids []string
	struct_map := make(map[string]struct{})
	for _, shelf := range shelfs {
		shelfids = append(shelfids, shelf.ID)

		struct_map[*shelf.StructID] = struct{}{}

		shelfs_map[shelf.RowIdentity] = append(shelfs_map[shelf.RowIdentity], shelf)
	}
	var structids []string
	for k := range struct_map {
		structids = append(structids, k)
	}
	var stucts []models.Libstruct
	err = global.OWEN_DB.Model(&models.Libstruct{}).Select("ID", "IsDeleted", "BuildNo", "FloorNo", "RoomNo", "BuildingName", "FloorName", "RoomName", "TenantID").Where("IsDeleted=0 and TenantId=?", tenantid).Where("Id in ?", structids).Find(&stucts).Error
	if err != nil {
		return resp, err
	}
	if len(stucts) <= 0 {
		return resp, fmt.Errorf("no data stuct")
	}
	var layers []models.Liblayer
	layers_map := make(map[string][]models.Liblayer) //按照列分号的层
	layer_query := global.OWEN_DB.Model(&models.Liblayer{}).Where("IsDeleted=0 and TenantId=?", tenantid)
	if islimitquery {
		layer_query = layer_query.Where("ShelfId in ?", shelfids)
	}
	err = layer_query.Find(&layers).Error
	if err != nil {
		return resp, err
	}
	if len(layers) <= 0 {
		return resp, fmt.Errorf("no data layer")
	}
	for _, ly := range layers {
		layers_map[ly.ShelfID] = append(layers_map[ly.ShelfID], ly)
	}

	for _, row := range rows {
		row_shelfs := shelfs_map[row.ID]
		var row_layers []models.Liblayer
		for _, shelf := range row_shelfs {
			shelf_layers := layers_map[shelf.ID]
			row_layers = append(row_layers, shelf_layers...)
		}
		_, err = ManagerGroup.frymanager.UploadRow(row, row_shelfs, row_layers, stucts, tenantid)
		if err != nil {
			return resp, err
		}
	}
	resp.Success = true
	return resp, nil
}
