package services

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"fmt"
)

type FlyReadAppService struct {
}

func (f *FlyReadAppService) Hello(req request.HelloInput) (resp response.HelloResp, err error) {
	return ManagerGroup.frymanager.Hello(req), nil
}

func (f *FlyReadAppService) UploadLibItem(req request.UploadLibItemInput, tenantid int) (resp response.UploadLibItemResp, err error) {
	if len(req.Barcodes) == 0 {
		return resp, fmt.Errorf("条码集合不能空")
	}

	var items []models.Libitem
	err = global.OWEN_DB.Model(&models.Libitem{}).Where(" Barcode in ?", req.Barcodes).Find(&items).Error
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

func (f *FlyReadAppService) UploadTenant(input request.UploadTenantInput) (resp response.UploadTenantDto, err error) {

	var tenant models.Abptenant
	err = global.OWEN_DB.Model(&models.Abptenant{}).Where("id=?", input.Tenantid).First(&tenant).Error
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
