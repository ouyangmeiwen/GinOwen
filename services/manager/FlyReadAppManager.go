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
	"sort"
	"strings"
	"time"
)

type FlyReadAppManager struct {
}

// helloworld
func (b *FlyReadAppManager) Hello(req request.HelloInput) (resp response.HelloResp) {

	var builder strings.Builder
	builder.WriteString(req.Name)
	fmt.Println(builder.String()) // 输出：数字是：123

	resp.Message = req.Name + ",Hello, FlyRead!"
	resp.Status = 200
	resp.Success = true
	return resp
}

// 盘点设置
func (b *FlyReadAppManager) SetFlyReadSetting(req dto.FlyReadSetting, tenantid int) (err error) {
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
func (b *FlyReadAppManager) GetFlyReadSetting(tenantid int) (resp dto.FlyReadSetting, err error) {

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

// 获取http地址
func (b *FlyReadAppManager) getHttpBySetting(flyseting dto.FlyReadSetting) string {

	url := fmt.Sprintf("http://%s:%s", flyseting.FlyReadIp, flyseting.FlyReadPort)
	if strings.Contains(flyseting.FlyReadIp, "http") || strings.Contains(flyseting.FlyReadIp, "https") {
		url = flyseting.FlyReadIp
	}
	return url
}

// 获取http地址
func (b *FlyReadAppManager) getHttpByTenant(tenantid int) string {
	var flyseting dto.FlyReadSetting
	flyseting, err := b.GetFlyReadSetting(tenantid)
	if err != nil {
		return ""
	}
	return b.getHttpBySetting(flyseting)
}

// 获取token
func (b *FlyReadAppManager) GetToken(tenantid int, IsForceRefresh bool) (resp string, err error) {
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
	var httpResp string
	httpResp, err = utils.GetFileContent("GetToken")
	if err != nil {
		fmt.Println("请求参数" + string(data))
		httpResp, err = utils.Post(url, data, headers)
		if err != nil {
			fmt.Println("请求失败:", err)
			return
		}
		if httpResp == "" {
			return "", fmt.Errorf("获取飞读Token失败，返回结果为空")
		}
		fmt.Println("请求返回" + httpResp)
	}
	var httpRespJson dto.GetTokenDto2
	err = json.Unmarshal([]byte(httpResp), &httpRespJson)
	if err != nil {
		fmt.Println("JSON 反序列化失败:", err)
	}
	if httpRespJson.Code != 0 {
		return "", fmt.Errorf("获取Token失败，错误代码: %d, 错误信息: %s", httpRespJson.Code, httpRespJson.Msg)
	}
	utils.SetCache(fmt.Sprintf("%d", tenantid), httpRespJson.Data.AccessToken, 50*time.Minute) //缓存设置
	return httpRespJson.Data.AccessToken, nil
}

// 图书推送
func (b *FlyReadAppManager) UploadLibItem(lst []models.Libitem, tenantid int) (resp bool, msg string, err error) {

	var token string
	token, err = b.GetToken(tenantid, false)
	if err != nil {
		return false, "", fmt.Errorf("Token获取失败" + err.Error())
	}
	url := b.getHttpByTenant(tenantid) + "/lcsapi/lcsinv"
	fmt.Println("UploadLibItem:", url)

	payload := dto.UploadBookInfoInput{}
	payload.Container = "lcsinv"
	payload.Component = "shelf"
	payload.Service = "service"
	payload.Token = token
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
		bookitem.Collection_time = item.CreationTime.Format("2006-01-02 15:04:05")
		bookitem.Shelf_time = ""
		payload.Obj.Books = append(payload.Obj.Books, bookitem)
	}

	// 将 map 序列化为 JSON 字节
	data, err := json.Marshal(payload)
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
	var httpResp string
	httpResp, err = utils.GetFileContent("UploadLibItem")
	if err != nil {
		fmt.Println("请求参数" + string(data))
		httpResp, err = utils.Post(url, data, headers)
		if err != nil {
			fmt.Println("请求失败:", err)
			return false, "", fmt.Errorf("失败，返回结果为空" + err.Error())
		}
		if httpResp == "" {
			return false, "", fmt.Errorf("失败，返回结果为空")
		}
		fmt.Println("请求返回" + httpResp)
	}
	var httpRespJson dto.UploadBookInfoResp
	err = json.Unmarshal([]byte(httpResp), &httpRespJson)
	if err != nil {
		fmt.Println("JSON 反序列化失败:", err)
		return false, "", fmt.Errorf("JSON 反序列化失败" + err.Error())
	}
	return httpRespJson.Success, httpRespJson.Message, err
}

// 租户推送
func (b *FlyReadAppManager) UploadTenant(tenant models.Abptenant) (resp bool, err error) {

	tenantid := int(tenant.ID)
	var token string
	token, err = b.GetToken(tenantid, false)
	if err != nil {
		return false, fmt.Errorf("Token获取失败" + err.Error())
	}
	url := b.getHttpByTenant(tenantid) + "/lcsapi/lcsinv"
	fmt.Println("UploadTenant:", url)

	//构建http请求
	var payload dto.UploadBranchInput
	payload.Container = "lcsinv"
	payload.Component = "shelf"
	payload.Service = "branch_sync"
	payload.Token = token
	payload.Obj = dto.Brandobj{
		Branch_no:   fmt.Sprintf("%d", tenant.ID),
		Branch_name: tenant.Name,
	}

	// 将 map 序列化为 JSON 字节
	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("JSON 序列化失败:", err)
		return
	}
	// 自定义请求头（可选）
	headers := map[string]string{
		"tenant-id":     fmt.Sprintf("%d", tenantid),
		"Cookie":        fmt.Sprintf("tenant={%d}", tenantid),
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	// 发起 POST 请求
	var httpResp string
	httpResp, err = utils.GetFileContent("UploadTenant")
	if err != nil {
		fmt.Println("请求参数" + string(data))
		httpResp, err = utils.Post(url, data, headers)
		if err != nil {
			fmt.Println("请求失败:", err)
			return
		}
		if httpResp == "" {
			return resp, fmt.Errorf("返回结果为空")
		}
		fmt.Println("请求返回" + httpResp)
	}
	var httpRespJson dto.UploadBranchResp
	err = json.Unmarshal([]byte(httpResp), &httpRespJson)
	if err != nil {
		fmt.Println("JSON 反序列化失败:", err)
	}
	if !httpRespJson.Success {
		return resp, fmt.Errorf("失败，错误代码: %d, 错误信息: %s", httpRespJson.Code, httpRespJson.Msg)
	}

	return true, nil
}

// 机构推送
func (b *FlyReadAppManager) UploadStruct(_struct models.Libstruct, tenantid int) (resp bool, err error) {
	var token string
	token, err = b.GetToken(tenantid, false)
	if err != nil {
		return false, fmt.Errorf("Token获取失败" + err.Error())
	}
	url := b.getHttpByTenant(tenantid) + "/lcsapi/lcsinv"
	fmt.Println("UploadRegion:", url)

	//构建http请求
	var payload dto.UploadRegionInput
	payload.Container = "lcsinv"
	payload.Component = "shelf"
	payload.Service = "region_sync"
	payload.Token = token

	var Region_name string
	if _struct.RoomName != nil {
		Region_name = *_struct.RoomName
	}
	payload.Obj = dto.RegionObj{
		Branch_no:   fmt.Sprintf("%d", tenantid),
		Region_no:   fmt.Sprintf("%02d%02d%02d", _struct.BuildNo, _struct.FloorNo, _struct.RoomNo),
		Region_name: Region_name,
	}

	// 将 map 序列化为 JSON 字节
	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("JSON 序列化失败:", err)
		return
	}
	// 自定义请求头（可选）
	headers := map[string]string{
		"tenant-id":     fmt.Sprintf("%d", tenantid),
		"Cookie":        fmt.Sprintf("tenant={%d}", tenantid),
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	// 发起 POST 请求
	var httpResp string
	httpResp, err = utils.GetFileContent("UploadRegion")
	if err != nil {
		fmt.Println("请求参数" + string(data))
		httpResp, err = utils.Post(url, data, headers)
		if err != nil {
			fmt.Println("请求失败:", err)
			return
		}
		if httpResp == "" {
			return resp, fmt.Errorf("返回结果为空")
		}
		fmt.Println("请求返回" + httpResp)
	}
	var httpRespJson dto.UploadBranchResp
	err = json.Unmarshal([]byte(httpResp), &httpRespJson)
	if err != nil {
		fmt.Println("JSON 反序列化失败:", err)
	}
	if !httpRespJson.Success {
		return resp, fmt.Errorf("失败，错误代码: %d, 错误信息: %s", httpRespJson.Code, httpRespJson.Msg)
	}
	return true, nil
}

// 推送图书定位
func (b *FlyReadAppManager) UploadLibItemLoc(lst []models.Libitemlocinfo, tenantid int) (resp bool, err error) {
	var token string
	token, err = b.GetToken(tenantid, false)
	if err != nil {
		return false, fmt.Errorf("Token获取失败" + err.Error())
	}
	url := b.getHttpByTenant(tenantid) + "/api/module/base/collection/create-case-item"
	fmt.Println("UploadLibItemLoc:", url)

	//构建http请求
	var payload []dto.UploadBookLocInput

	locmap := make(map[string][]models.Libitemlocinfo)
	//按照层码分类处理成map
	for _, val := range lst {
		key := *val.LayerCode
		locmap[key] = append(locmap[key], val)
	}
	//循环map 把数据构建好
	for key, vals := range locmap {
		var it dto.UploadBookLocInput
		it.CaseNo = key
		for _, item := range vals {
			if item.ItemBarcode != "" {
				it.AssetIdSet = append(it.AssetIdSet, item.ItemBarcode)
			}
		}
		payload = append(payload, it)
	}

	// 将 map 序列化为 JSON 字节
	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("JSON 序列化失败:", err)
		return
	}
	// 自定义请求头（可选）
	headers := map[string]string{
		"tenant-id":     fmt.Sprintf("%d", tenantid),
		"Cookie":        fmt.Sprintf("tenant={%d}", tenantid),
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	// 发起 POST 请求
	var httpResp string
	httpResp, err = utils.GetFileContent("UploadLibItemLoc")
	if err != nil {
		fmt.Println("请求参数" + string(data))
		httpResp, err = utils.Post(url, data, headers)
		if err != nil {
			fmt.Println("请求失败:", err)
			return
		}
		if httpResp == "" {
			return resp, fmt.Errorf("返回结果为空")
		}
		fmt.Println("请求返回" + httpResp)
	}
	var httpRespJson dto.UploadBookInfoLocDto
	err = json.Unmarshal([]byte(httpResp), &httpRespJson)
	if err != nil {
		fmt.Println("JSON 反序列化失败:", err)
	}
	if httpRespJson.Code != 0 {
		return resp, fmt.Errorf("失败，错误代码: %d, 错误信息: %s", httpRespJson.Code, httpRespJson.Msg)
	}
	return true, nil
}

// 书架推送
func (b *FlyReadAppManager) UploadRow(row models.Librow,
	shelfs []models.Libshelf,
	layers []models.Liblayer,
	structs []models.Libstruct,
	tenantid int) (resp bool, err error) {
	var token string
	token, err = b.GetToken(tenantid, false)
	if err != nil {
		return false, fmt.Errorf("Token获取失败" + err.Error())
	}
	url := b.getHttpByTenant(tenantid) + "/lcsapi/lcsinv"
	fmt.Println("UploadRow:", url)

	//构建http请求
	var payload dto.UploadLayersInput
	payload.Container = "lcsinv"
	payload.Component = "shelf"
	payload.Service = "shelf_sync"
	payload.Token = token

	needShelfs := make(map[string][]models.Libshelf) //按照面分组
	layermap := make(map[string][]models.Liblayer)   //根据shelfid 分层
	for _, val := range shelfs {
		if val.Code != nil {
			code := utils.SafeSubstring(*val.Code, 0, len(*val.Code)-3)
			needShelfs[code] = append(needShelfs[code], val)
		}
	}
	// 提取并排序 key
	keys := make([]string, 0, len(needShelfs))
	for key := range needShelfs {
		keys = append(keys, key)
	}
	sort.Strings(keys) // 升序排序，如需降序可用 sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	for _, code := range keys {
		columns := needShelfs[code]

		shelf_no := code               //面编码
		side := *columns[0].Side       //AB面
		shelf_name := *row.Name + side //面名称
		// needStruct := parallel.SmartFilter(structs, 1, 1, func(ly models.Libstruct) bool {
		// 	return *columns[0].StructID == ly.ID
		// })
		var needStruct models.Libstruct
		for _, val := range structs {
			if *columns[0].StructID == val.ID {
				needStruct = val
				break
			}
		}

		region_no := fmt.Sprintf("%02d%02d%02d", needStruct.BuildNo, needStruct.FloorNo, needStruct.RoomNo)
		for _, col := range columns {
			// col_layers := parallel.SmartFilter(layers, 1, 1, func(ly models.Liblayer) bool {
			// 	return ly.ShelfID == col.ID
			// })

			var col_layers []models.Liblayer
			for _, val := range layers {
				if val.ShelfID == col.ID {
					col_layers = append(col_layers, val)
				}
			}

			layermap[col.ID] = append(layermap[col.ID], col_layers...)
		}
		sidelshelfs := dto.UploadLayersShelf{} //面信息
		sidelshelfs.Shelf_info = dto.UploadLayersShelf_info{
			Shelf_no:      shelf_no,
			Shelf_name:    shelf_name,
			Branch_no:     fmt.Sprintf("%d", tenantid),
			Case_num:      6,
			Region_no:     region_no,
			Column_num:    len(columns),
			Side:          side,
			First_call_no: "",
			Last_call_no:  "",
		}
		sidelshelfs.Columns = []dto.UploadLayersColumns{} //面下面的列信息
		for _, col := range columns {
			var it dto.UploadLayersColumns
			it.Column_info = dto.UploadLayersColumnInfo{
				Column_no:     *col.Code,
				Column_name:   *col.Name,
				Shelf_no:      shelf_no,
				Case_num:      6,
				Column_seq:    int(col.ShelfNo),
				First_call_no: "",
				Last_call_no:  "",
			}
			it.Cases = []dto.UploadLayersCase{}
			layer := layermap[col.ID]
			if len(layer) > 0 {
				for _, ly := range layer {
					lydto := dto.UploadLayersCase{
						Case_no:       *ly.Code,
						Case_name:     ly.Name,
						Shelf_no:      shelf_no,
						Column_no:     *col.Code,
						Case_seq:      int(ly.LayerNo),
						First_call_no: "",
						Last_call_no:  "",
					}
					it.Cases = append(it.Cases, lydto)
				}
			}
			sidelshelfs.Columns = append(sidelshelfs.Columns, it)
		}
		payload.Obj.Shelfs = append(payload.Obj.Shelfs, sidelshelfs)
	}

	// 将 map 序列化为 JSON 字节
	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("JSON 序列化失败:", err)
		return
	}
	// 自定义请求头（可选）
	headers := map[string]string{
		"tenant-id":     fmt.Sprintf("%d", tenantid),
		"Cookie":        fmt.Sprintf("tenant={%d}", tenantid),
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	// 发起 POST 请求
	var httpResp string
	httpResp, err = utils.GetFileContent("UploadRow")
	if err != nil {
		fmt.Println("请求参数" + string(data))
		httpResp, err = utils.Post(url, data, headers)
		if err != nil {
			fmt.Println("请求失败:", err)
			return
		}
		if httpResp == "" {
			return resp, fmt.Errorf("返回结果为空")
		}
		fmt.Println("请求返回" + httpResp)
	}
	var httpRespJson dto.UploadLayersDto
	err = json.Unmarshal([]byte(httpResp), &httpRespJson)
	if err != nil {
		fmt.Println("JSON 反序列化失败:", err)
	}
	if httpRespJson.Code != 0 {
		return resp, fmt.Errorf("失败，错误代码: %d, 错误信息: %s", httpRespJson.Code, httpRespJson.Msg)
	}
	return true, nil
}
