package services

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"errors"
)

type SysauditlmslogService struct {
}

// 查询审计日志LMS
func (SysauditlmslogService) QueryLmsLog(req request.QueryLmsInput) (resp response.QueryLmsDto, err error) {

	var lmslogs_lst []models.Sysauditlmslog
	logquery := global.OWEN_DB.Model(&models.Sysauditlmslog{})
	if req.IsException != nil && *req.IsException {
		logquery = logquery.Where(" (Exception !='' and Exception is not null)")
	}
	if req.ExecutionDuration > 0 {
		logquery = logquery.Where(" ExecutionDuration>? ", req.ExecutionDuration)
	}
	if req.MethodName != nil && len(*req.MethodName) > 0 {
		logquery = logquery.Where(" MethodName like ? ", "%"+*req.MethodName+"%")
	}
	if req.Parameters != nil && len(*req.Parameters) > 0 {
		logquery = logquery.Where(" Parameters like ? ", "%"+*req.Parameters+"%")
	}
	if req.ServiceName != nil && len(*req.ServiceName) > 0 {
		logquery = logquery.Where(" ServiceName like ? ", "%"+*req.ServiceName+"%")
	}
	if req.TenantId > 0 {
		logquery = logquery.Where(" TenantId = ? ", req.TenantId)
	}
	logquery = logquery.Order(" ExecutionTime desc")

	var cnt int64
	logquery.Count(&cnt)
	err = logquery.Offset(req.SkipCount).Limit(req.MaxResultCount).Find(&lmslogs_lst).Error
	if err != nil {
		return resp, err
	}
	resp.TotalCount = cnt
	for _, item := range lmslogs_lst {
		var cp response.QueryLmsItem
		cp.ID = item.ID
		cp.BrowserInfo = item.BrowserInfo
		cp.ClientIPAddress = item.ClientIPAddress
		cp.ClientName = item.ClientName
		cp.CustomData = item.CustomData
		cp.Exception = item.Exception
		cp.ExecutionDuration = item.ExecutionDuration
		cp.ExecutionTime = item.ExecutionTime
		cp.ImpersonatorTenantID = item.ImpersonatorTenantID
		cp.ImpersonatorUserID = item.ImpersonatorUserID
		cp.MethodName = item.MethodName
		cp.Parameters = item.Parameters
		cp.ServiceName = item.ServiceName
		cp.TenantID = item.TenantID
		cp.UserID = item.UserID
		cp.ReturnValue = item.ReturnValue
		resp.Items = append(resp.Items, cp)
	}
	return resp, err
}

// 新增审计日志
func (SysauditlmslogService) CreateLmsLog(req request.CreateLmsLogInput) (resp response.CreateLmsLogDto, err error) {

	var lmslog models.Sysauditlmslog
	lmslog.ID = 0
	lmslog.BrowserInfo = req.BrowserInfo
	lmslog.ClientIPAddress = req.ClientIPAddress
	lmslog.ClientName = req.BrowserInfo
	lmslog.CustomData = req.CustomData
	lmslog.Exception = req.Exception
	lmslog.ExecutionDuration = req.ExecutionDuration
	lmslog.ExecutionTime = req.ExecutionTime
	lmslog.ImpersonatorTenantID = req.ImpersonatorTenantID
	lmslog.ImpersonatorUserID = req.ImpersonatorUserID
	lmslog.MethodName = req.MethodName
	lmslog.Parameters = req.Parameters
	lmslog.ServiceName = req.ServiceName
	lmslog.TenantID = req.TenantID
	lmslog.UserID = req.UserID
	lmslog.ReturnValue = req.ReturnValue
	err = global.OWEN_DB.Model(&models.Sysauditlmslog{}).Create(&lmslog).Error
	if err != nil {
		return resp, err
	}
	return resp, err
}

// 修改审计日志
func (SysauditlmslogService) UpdateLmsLog(req request.UpdateLmsLogInput) (resp response.UpdateLmsLogDto, err error) {

	var lmslog models.Sysauditlmslog
	err = global.OWEN_DB.Model(&models.Sysauditlmslog{}).Where("Id=?", req.ID).First(&lmslog).Error
	if err != nil {
		return resp, err
	}
	if lmslog.ID <= 0 {
		return resp, errors.New("ID查询失败,无法完成更新")
	}
	lmslog.BrowserInfo = req.BrowserInfo
	lmslog.ClientIPAddress = req.ClientIPAddress
	lmslog.ClientName = req.BrowserInfo
	lmslog.CustomData = req.CustomData
	lmslog.Exception = req.Exception
	lmslog.ExecutionDuration = req.ExecutionDuration
	lmslog.ExecutionTime = req.ExecutionTime
	lmslog.ImpersonatorTenantID = req.ImpersonatorTenantID
	lmslog.ImpersonatorUserID = req.ImpersonatorUserID
	lmslog.MethodName = req.MethodName
	lmslog.Parameters = req.Parameters
	lmslog.ServiceName = req.ServiceName
	lmslog.TenantID = req.TenantID
	lmslog.UserID = req.UserID
	lmslog.ReturnValue = req.ReturnValue
	err = global.OWEN_DB.Model(&models.Sysauditlmslog{}).Where("Id=?", req.ID).Save(&lmslog).Error
	if err != nil {
		return resp, err
	}
	return resp, err
}

// 删除指定审计日志
func (SysauditlmslogService) DeleteLmsLog(req request.DeleteLmsLogInput) (resp response.DeleteLmsLogDto, err error) {

	var lmslog models.Sysauditlmslog
	err = global.OWEN_DB.Model(&models.Sysauditlmslog{}).Where("Id=?", req.ID).First(&lmslog).Error
	if err != nil {
		return resp, err
	}
	if lmslog.ID <= 0 {
		return resp, err
	}
	err = global.OWEN_DB.Model(&models.Sysauditlmslog{}).Where("Id=?", req.ID).Unscoped().Delete(&models.Sysauditlmslog{}).Error
	if err != nil {
		return resp, err
	}
	return resp, err
}
