package manager

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/models/dto"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"GINOWEN/utils"
	"encoding/json"
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
	utils.SetCache(fmt.Sprintf("%dFlySetting", tenantid), req, 0) //缓存设置
	global.OWEN_DB.Model(&models.Systenantconfig{}).Save(sysconfig)
	return nil
}

// 获取盘点设置
func (FlyReadAppManager) GetFlyReadSetting(tenantid int) (resp dto.FlyReadSetting, err error) {

	var cachedSetting dto.FlyReadSetting
	found, err := utils.GetCache(fmt.Sprintf("%dFlySetting", tenantid), &cachedSetting)
	if err != nil {
		return resp, err
	}
	if found {
		return cachedSetting, nil
	}

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
func (b FlyReadAppManager) getHttpBySetting(flyseting dto.FlyReadSetting) string {

	url := fmt.Sprintf("http://%s:%s", flyseting.FlyReadIp, flyseting.FlyReadPort)
	if strings.Contains(flyseting.FlyReadIp, "http") || strings.Contains(flyseting.FlyReadIp, "https") {
		url = flyseting.FlyReadIp
	}
	return url
}
func (b FlyReadAppManager) getHttpByTenant(tenantid int) string {
	var flyseting dto.FlyReadSetting
	flyseting, err := b.GetFlyReadSetting(tenantid)
	if err != nil {
		return ""
	}
	return b.getHttpBySetting(flyseting)
}
func (b FlyReadAppManager) GetToken(tenantid int, IsForceRefresh bool) (resp string, err error) {
	if !IsForceRefresh {
		key := fmt.Sprintf("%d", tenantid)
		var token string
		found, err := utils.GetCache(key, &token)
		if err != nil {
			return resp, err
		}
		if found {
			return token, nil
		}
	}
	var flyseting dto.FlyReadSetting
	flyseting, err = b.GetFlyReadSetting(tenantid)
	if err != nil {
		return
	}
	url := b.getHttpBySetting(flyseting) + "/api/module/base/collection-auth/get-auth-token"
	fmt.Println("获取飞读Token:", url)

	// 请求体数据（会被序列化为 JSON）
	payload := map[string]interface{}{
		"appId":     flyseting.FlyAppid,
		"appSecret": flyseting.FlyReadAppSecret,
	}
	// 将 map 序列化为 JSON 字节
	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("JSON 序列化失败:", err)
		return
	}
	// 自定义请求头（可选）
	headers := map[string]string{
		"tenant-id": fmt.Sprintf("%d", tenantid),
		"Cookie":    fmt.Sprintf("tenant={%d}", tenantid),
	}
	// 发起 POST 请求
	var tokenResp string
	tokenResp, err = utils.GetFileContent("GetToken")
	if err != nil {
		fmt.Println("请求参数" + string(data))
		tokenResp, err = utils.Post(url, data, headers)
		if err != nil {
			fmt.Println("请求失败:", err)
			return
		}
		if tokenResp == "" {
			return "", fmt.Errorf("获取飞读Token失败，返回结果为空")
		}
	}
	var tokenData dto.GetTokenDto2
	err = json.Unmarshal([]byte(tokenResp), &tokenData)
	if err != nil {
		fmt.Println("JSON 反序列化失败:", err)
	}
	if tokenData.Code != 0 {
		return "", fmt.Errorf("获取Token失败，错误代码: %d, 错误信息: %s", tokenData.Code, tokenData.Msg)
	}
	utils.SetCache(fmt.Sprintf("%d", tenantid), tokenData.Data.AccessToken, 50*time.Minute) //缓存设置
	return tokenData.Data.AccessToken, nil
}

func (b FlyReadAppManager) UploadLibItem(lst []models.Libitem, tenantid int) (resp bool, msg string, err error) {

	var token string
	token, err = b.GetToken(tenantid, false)
	if err != nil {
		return false, "", fmt.Errorf("Token获取失败" + err.Error())
	}
	input := dto.UploadBookInfoInput{}
	input.Container = "lcsinv"
	input.Component = "shelf"
	input.Service = "service"
	input.Token = token
	for _, item := range lst {
		var bookitem dto.BookItem
		bookitem.Asset_id = item.Barcode
		bookitem.Barcode = ""
		if item.CallNo != nil {
			bookitem.Call_no = *item.CallNo
		}
		bookitem.Title = item.Title
		if item.Author != nil {
			bookitem.Author = *item.Author
		}
		if item.Publisher != nil {
			bookitem.Publisher = *item.Publisher
		}
		if item.PubDate != nil {
			bookitem.Pubdate = *item.PubDate
		}
		if item.CatalogCode != nil {
			bookitem.Clc = *item.CatalogCode
		}
		if item.LocationName != nil {
			bookitem.Lib_location = *item.LocationName
		}
		bookitem.Status = fmt.Sprintf("%d", item.ItemState)
		if item.CreationTime != nil {
			bookitem.Collection_time = item.CreationTime.Format("2006-01-02 15:04:05")
		} else {
			bookitem.Collection_time = ""
		}
		bookitem.Shelf_time = ""
		input.Obj.Books = append(input.Obj.Books, bookitem)
	}
	url := b.getHttpByTenant(tenantid) + "/lcsapi/lcsinv"
	fmt.Println("UploadLibItem:", url)

	// 将 map 序列化为 JSON 字节
	data, err := json.Marshal(input)
	if err != nil {
		fmt.Println("JSON 序列化失败:", err)
		return false, "", fmt.Errorf("JSON 序列化失败" + err.Error())
	}
	// 自定义请求头（可选）
	headers := map[string]string{
		"tenant-id":     fmt.Sprintf("%d", tenantid),
		"Cookie":        fmt.Sprintf("tenant={%d}", tenantid),
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	// 发起 POST 请求
	var UploadLibItemResp string
	UploadLibItemResp, err = utils.GetFileContent("UploadLibItem")
	if err != nil {
		fmt.Println("请求参数" + string(data))
		UploadLibItemResp, err = utils.Post(url, data, headers)
		if err != nil {
			fmt.Println("请求失败:", err)
			return false, "", fmt.Errorf("获取飞读Token失败，返回结果为空" + err.Error())
		}
		if UploadLibItemResp == "" {
			return false, "", fmt.Errorf("获取飞读Token失败，返回结果为空")
		}
	}
	var uploadLibItemResp dto.UploadBookInfoResp
	err = json.Unmarshal([]byte(UploadLibItemResp), &uploadLibItemResp)
	if err != nil {
		fmt.Println("JSON 反序列化失败:", err)
		return false, "", fmt.Errorf("JSON 反序列化失败" + err.Error())
	}
	return uploadLibItemResp.Success, uploadLibItemResp.Message, err
}
