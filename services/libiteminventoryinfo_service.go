package services

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"GINOWEN/utils"
	"errors"
	"time"
)

type LibiteminventoryinfoService struct {
}

func (LibiteminventoryinfoService) GetInventoryList(req request.GetInventoryInput) (resp response.GetInventoryDto, err error) {
	var results []models.Libiteminventoryinfo
	query := global.OWEN_DB.Model(&models.Libiteminventoryinfo{})
	if len(req.Barcode) > 0 {
		query = query.Where(" ItemBarcode=?", req.Barcode)
	}
	if len(req.Title) > 0 {
		query = query.Where(" OCRItemTitle like ?", "%"+req.Title+"%")
	}
	if len(req.CallNo) > 0 {
		query = query.Where(" OCRItemCallNo = ?", req.CallNo)
	}
	if len(req.ISBN) > 0 {
		query = query.Where(" OCRItemISBN = ?", req.ISBN)
	}
	if len(req.LayerCode) > 0 {
		query = query.Where(" LayerCode = ?", req.LayerCode)
	}
	if len(req.LayerName) > 0 {
		query = query.Where(" LayerName = ?", req.LayerName)
	}
	if len(req.LocLayerCode) > 0 {
		query = query.Where(" LocLayerCode = ?", req.LocLayerCode)
	}
	if len(req.LocLayerName) > 0 {
		query = query.Where(" LocLayerName = ?", req.LocLayerName)
	}
	query = query.Order(" CreationTime desc")
	var total int64
	query.Count(&total)
	err = query.Offset(req.SkipCount).Limit(req.MaxResultCount).Find(&results).Error
	if err != nil {
		return resp, err
	} else {
		resp.Items = results
		resp.TotalCount = (total)
		return resp, err
	}
}

func (LibiteminventoryinfoService) CreateInventory(req request.CreateInventoryInput) (resp response.CreateInventoryDto, err error) {
	req.Item.ID = utils.GenerateGUID()
	req.Item.CreationTime = time.Now()
	err = global.OWEN_DB.Save(req.Item).Error
	if err != nil {
		return resp, errors.New("新增异常")
	}
	return resp, err
}

func (LibiteminventoryinfoService) UpdateInventory(req request.UpdateInventoryInput) (resp response.UpdateInventoryDto, err error) {

	if len(req.Item.ID) <= 0 {
		return resp, errors.New("修改的主键不能空！")
	}
	var first models.Libiteminventoryinfo
	err = global.OWEN_DB.Model(&models.Libiteminventoryinfo{}).Where("Id=?", req.Item.ID).First(&first).Error
	if err != nil {
		return resp, errors.New("主键查询异常")
	}
	if len(first.ID) <= 0 {
		return resp, errors.New("主键ID不存在")
	}
	err = global.OWEN_DB.Model(&models.Libiteminventoryinfo{}).Where("Id=?", req.Item.ID).Save(&req.Item).Error
	if err != nil {
		return resp, errors.New("更新异常")
	}
	return resp, err
}

func (LibiteminventoryinfoService) DeleteInventory(req request.DeleteInventoryInput) (resp response.DeleteInventoryDto, err error) {
	var first models.Libiteminventoryinfo
	err = global.OWEN_DB.Model(&models.Libiteminventoryinfo{}).Where("Id=?", req.Id).First(&first).Error
	if err != nil {
		return resp, errors.New("主键查询异常")
	}
	if len(first.ID) <= 0 {
		return resp, errors.New("主键ID不存在")
	}

	err = global.OWEN_DB.Model(&models.Libiteminventoryinfo{}).Where("Id=?", req.Id).Delete(&first).Error
	if err != nil {
		return resp, errors.New("删除异常")
	}
	return resp, err
}
