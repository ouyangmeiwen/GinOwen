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
	err = global.OWEN_DB.Model(&models.Libitemlocinfo{}).Where("IsDeleted=0 and LayerCode in ?", input.Layercode).Find(&lst).Error
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
