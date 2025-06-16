package manager

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/models/dto"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"GINOWEN/utils"
	"fmt"
	"strings"
	"time"
)

type FlyReadAppManager struct {
}

func (f FlyReadAppManager) Hello(req request.HelloInput) (resp response.HelloResp) {

	var builder strings.Builder
	builder.WriteString(req.Name)
	fmt.Println(builder.String()) // 输出：数字是：123

	resp.Message = req.Name + ",Hello, FlyRead!"
	resp.Status = 200
	resp.Success = true
	return resp
}

// 盘点设置
func (FlyReadAppManager) SetFlyReadSetting(req dto.FlyReadSetting, tenantid int) (err error) {
	// 这里可以添加逻辑来处理设置
	var sysconfig models.Systenantconfig
	err = global.OWEN_DB.Model(&models.Systenantconfig{}).Where("TenantId = ?", tenantid).First(&sysconfig).Error
	if err != nil {
		return err
	}
	if sysconfig.ID == "" {
		//add
		var json string
		json, err = utils.ToJSON(req)
		if err != nil {
			return err
		}
		sysconfig.System = &json
		sysconfig.ID = utils.GenerateGUID()
		sysconfig.CreationTime = time.Now()
		sysconfig.TenantID = int64(tenantid)
	} else {
		//update
		var sc dto.SystemConfig
		err = utils.FromJSON(*sysconfig.System, &sc)
		if err != nil {
			return err
		}
		sc.FlyReadEnable = req.FlyReadEnable
		sc.FlyReadIp = req.FlyReadIp
		sc.FlyReadPort = req.FlyReadPort
		sc.FlyAppid = req.FlyAppid
		sc.FlyReadAppSecret = req.FlyReadAppSecret
		sc.FlyMode = req.FlyMode
		sc.FlyAddLoc = req.FlyAddLoc
		sc.FlyLocType = req.FlyLocType
		sc.FlyLostDay = req.FlyLostDay
		sc.FlyStartTime = req.FlyStartTime
		sc.FlyEndTime = req.FlyEndTime
		sc.FlyTenant = req.FlyTenant
		sc.FlyShowOff = req.FlyShowOff
		sc.FlyDeviceType = req.FlyDeviceType
		sc.FlyTempLoc = req.FlyTempLoc
		sc.RowShape = req.RowShape
		var json string
		json, err = utils.ToJSON(sc)
		if err != nil {
			return err
		}
		sysconfig.System = &json
		dtupdate := time.Now()
		sysconfig.LastModificationTime = &dtupdate
	}
	global.OWEN_DB.Model(&models.Systenantconfig{}).Save(sysconfig)
	return nil
}

// 获取盘点设置
func (FlyReadAppManager) GetFlyReadSetting(tenantid int) (resp dto.FlyReadSetting, err error) {
	var sysconfig models.Systenantconfig
	err = global.OWEN_DB.Model(&models.Systenantconfig{}).Where("TenantId = ?", tenantid).First(&sysconfig).Error
	if err != nil {
		return resp, err
	}
	if sysconfig.ID == "" {
		return resp, nil //没有配置
	}
	var sc dto.SystemConfig
	err = utils.FromJSON(*sysconfig.System, &sc)
	if err != nil {
		return resp, err
	}
	resp.FlyReadEnable = sc.FlyReadEnable
	resp.FlyReadIp = sc.FlyReadIp
	resp.FlyReadPort = sc.FlyReadPort
	resp.FlyAppid = sc.FlyAppid
	resp.FlyReadAppSecret = sc.FlyReadAppSecret
	resp.FlyMode = sc.FlyMode
	resp.FlyAddLoc = sc.FlyAddLoc
	resp.FlyLocType = sc.FlyLocType
	resp.FlyLostDay = sc.FlyLostDay
	resp.FlyStartTime = sc.FlyStartTime
	resp.FlyEndTime = sc.FlyEndTime
	resp.FlyTenant = sc.FlyTenant
	resp.FlyShowOff = sc.FlyShowOff
	resp.FlyDeviceType = sc.FlyDeviceType
	resp.FlyTempLoc = sc.FlyTempLoc
	resp.RowShape = sc.RowShape

	return resp, nil
}
