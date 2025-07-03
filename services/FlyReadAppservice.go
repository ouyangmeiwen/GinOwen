package services

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/models/dto"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"GINOWEN/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
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

// 盘点指令下发
func (f *FlyReadAppService) Inventory(req request.InventoryInput, tenantid int) (resp response.InventoryDto, err error) {

	if req.Workid == "" {
		req.Workid = utils.UUID32()
	}
	var layers []models.Liblayer
	if !req.IsAll {
		layers, err = ManagerGroup.frymanager.GetEnableLayers(tenantid, req.IsAll, req.Devicetype, req.RobotId, req.RobotRouterId)
		if err != nil {
			return resp, err
		}
		var layerneed []models.Liblayer
		for _, ly := range layers {
			for _, id := range req.Layerids {
				if id == ly.ID || (ly.Code != nil && id == *ly.Code) {
					layerneed = append(layerneed, ly)
				}
			}
		}

		if len(layerneed) != len(req.Layerids) {
			return resp, fmt.Errorf("启用层不匹配！")
		}
		layers = layerneed
	}

	var dto_resp dto.InventoryDto
	dto_resp, err = ManagerGroup.frymanager.Inventory(req.IsAll, req.Workid, req.Triggers, req.Workname, layers, tenantid, req.Devicetype, req.RobotId, req.RobotRouterId)
	if err != nil {
		return resp, err
	}
	resp.Success = dto_resp.Success
	return resp, nil
}

func (f *FlyReadAppService) InventoryHis(req request.InventoryHisInput, tenantid int) (resp response.InventoryHisDto, err error) {
	var dto_resp dto.InventoryHisDto
	dto_resp, err = ManagerGroup.frymanager.InventoryHis(req.IsHistory, tenantid)
	if err != nil {
		return resp, err
	}
	resp.Success = dto_resp.Success
	resp.InventoryHisRespObj = dto_resp.Obj
	return resp, nil
}

func (f *FlyReadAppService) InventoryList(req request.InventoryListInput, tenantid int) (resp response.InventoryListDto, err error) {
	var dto_resp dto.InventoryListDto
	var dtstart time.Time
	var dtend time.Time
	dtstart, err = utils.ParseInLocation("2006-01-02", req.DtStart)
	if err != nil {
		return resp, fmt.Errorf("开始时间格式错误: %v", err)
	}
	dtend, err = utils.ParseInLocation("2006-01-02", req.DtEnd)
	if err != nil {
		return resp, fmt.Errorf("结束时间格式错误: %v", err)
	}

	dto_resp, err = ManagerGroup.frymanager.InventoryList(req.PageIndex, req.PageSize, dtstart, dtend, tenantid)
	if err != nil {
		return resp, err
	}
	resp.List = dto_resp.Obj.List
	return resp, nil
}

func (f *FlyReadAppService) SetBussiness(input request.SetBussinessInput, tenantid int) (resp response.SetBussinessDto, err error) {
	var sets dto.FlyReadSetting
	sets, err = ManagerGroup.frymanager.GetFlyReadSetting(tenantid)
	if err != nil {
		return resp, err
	}
	if input.Addloc == "" || input.Newlocmode == "" || input.Lostday == "" {
		return resp, fmt.Errorf("addloc, newlocmode, lostday 不能为空")
	}
	sets.FlyAddLoc = input.Addloc
	sets.FlyNewLocMode = input.Newlocmode
	sets.FlyLostDay = input.Lostday
	switch input.Newlocmode {
	case "1":
		sets.FlyLocType = "2"
		sets.FlyTempLoc = ""
	case "2":
		sets.FlyLocType = "2"
		sets.FlyTempLoc = "1"
	case "3":
		sets.FlyLocType = "2"
		sets.FlyTempLoc = "2"
	case "4":
		sets.FlyLocType = "1"
	case "5":
		sets.FlyLocType = "0"
	}
	if input.Shape == "0" || input.Shape == "1" {
		sets.RowShape = input.Shape
	} else {
		sets.RowShape = "0"
	}
	err = ManagerGroup.frymanager.SetFlyReadSetting(sets, tenantid)
	if err != nil {
		return resp, err
	}
	resp.Success = true
	return resp, nil
}

// 获取盘点设置
func (f *FlyReadAppService) GetInventorySet(input request.GetInventorySetInput, tenantid int) (resp response.GetInventorySetDto, err error) {
	fset, e := ManagerGroup.frymanager.GetFlyReadSetting(tenantid)
	if e != nil {
		return resp, e
	}
	var firsttask models.Libinventorytask
	err = global.OWEN_DB.Model(&models.Libinventorytask{}).Where("TenantId=? and IsDeleted=0 and OriginType=10 and TaskType=0 and TriggerSatus=0", tenantid).First(&firsttask).Error
	if err != nil {
		return resp, fmt.Errorf("获取盘点设置失败: %v", err)
	}
	if firsttask.ID == "" {
		return resp, fmt.Errorf("没有设置盘点任务")
	}
	resp.SysStartTime = fset.FlyStartTime
	resp.SysEndTime = fset.FlyEndTime
	resp.IsEnable = utils.Uint8slicetoboolslice(firsttask.IsEnable)
	if firsttask.InventoryStartDate != nil {
		resp.InventoryStartDate = firsttask.InventoryStartDate.Format("2006-01-02")
	} else {
		resp.InventoryStartDate = ""
	}
	if firsttask.InventoryEndDate != nil {
		resp.InventoryEndDate = firsttask.InventoryEndDate.Format("2006-01-02")
	} else {
		resp.InventoryEndDate = ""
	}
	resp.Interval = int(firsttask.Interval)
	if firsttask.DeviceType != nil {
		resp.DeviceType = *firsttask.DeviceType
	}
	return resp, nil
}

func (f *FlyReadAppService) GetEnableRow(input request.GetEnableRowInput, tenantid int) (resp response.GetEnableRowDto, err error) {
	var dto_resp dto.GetEnableRowDto
	dto_resp, err = ManagerGroup.frymanager.GetEnableRow(input.IsQueryAll, input.Devicetype, tenantid)
	if err != nil {
		return resp, err
	}
	resp.GetEnableRowDto = dto_resp
	return resp, nil
}

// 获取机器人列表
func (f *FlyReadAppService) GetRobotlist(input request.GetRobotlistInput, tenantid int) (resp response.GetRobotlistDto, err error) {

	var dto_input dto.RobotlistInput
	var dto_resp dto.RobotlistDto
	dto_resp, err = ManagerGroup.frymanager.GetRobotlist(dto_input, tenantid)
	if err != nil {
		return resp, err
	}
	resp.RobotlistDto = dto_resp
	return resp, nil
}

func (f *FlyReadAppService) InventorySet(input request.InventorySetInput, tenantid int) (resp response.InventorySetDto, err error) {
	if (input.SysStartTime == "" && input.SysEndTime != "") || input.SysStartTime != "" && input.SysEndTime == "" {
		return resp, fmt.Errorf("盘点工作时间设置不合理")
	}
	if input.Interval <= 0 || input.Interval >= 24 {
		return resp, fmt.Errorf("间隔小时设置设置不合理")
	}
	// 获取当前日期
	loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(loc)
	today := now.Format("2006-01-02")
	var SysStartTime time.Time
	var SysEndTime time.Time
	if input.SysStartTime != "" {
		SysStartTime, err = utils.ParseInLocation("2006-01-02 15:04", today+" "+input.SysStartTime)
		if err != nil {
			return resp, fmt.Errorf("SysStartTime 格式错误: %v", err)
		}
	}
	if input.SysEndTime != "" {
		SysEndTime, err = utils.ParseInLocation("2006-01-02 15:04", today+" "+input.SysEndTime)
		if err != nil {
			return resp, fmt.Errorf("SysEndTime 格式错误: %v", err)
		}
	}

	var InventoryStartDate *time.Time
	InventoryStartDate1, er1 := utils.ParseInLocation("2006-01-02", input.InventoryStartDate)
	if er1 != nil {
		InventoryStartDate = nil
	} else {
		InventoryStartDate = &InventoryStartDate1
	}
	var InventoryEndDate *time.Time
	InventoryEndDate1, err2 := utils.ParseInLocation("2006-01-02", input.InventoryEndDate)
	if err2 != nil {
		InventoryEndDate = nil
	} else {
		InventoryEndDate = &InventoryEndDate1
	}

	var fsets dto.FlyReadSetting
	fsets, err = ManagerGroup.frymanager.GetFlyReadSetting(tenantid)
	if err != nil {
		return resp, err
	}
	if input.SysStartTime != "" && input.SysEndTime != "" {
		if SysStartTime.After(SysEndTime) {
			return resp, fmt.Errorf("SysStartTime 不能大于 SysEndTime")
		} else {
			fsets.FlyStartTime = input.SysStartTime
			fsets.FlyEndTime = input.SysEndTime
		}
	} else if input.SysStartTime == "" && input.SysEndTime == "" {
		fsets.FlyStartTime = ""
		fsets.FlyEndTime = ""
	}
	if input.DeviceType != "0" && input.DeviceType != "1" {
		return resp, fmt.Errorf("间隔全盘仅限deviceType只能是1,2")
	}
	fsets.FlyDeviceType = input.DeviceType
	err = ManagerGroup.frymanager.SetFlyReadSetting(fsets, tenantid)
	if err != nil {
		return resp, err
	}
	var firsttask models.Libinventorytask
	err = global.OWEN_DB.Model(&models.Libinventorytask{}).Where("TenantId=? and IsDeleted=0 and OriginType=10 and TaskType=0 and TriggerSatus=0", tenantid).First(&firsttask).Error
	if err != nil {
		return resp, fmt.Errorf("获取盘点设置失败: %v", err)
	}
	if firsttask.ID != "" {
		err = global.OWEN_DB.
			Model(&models.Libinventorytask{}).
			Where("TenantId = ? AND IsDeleted = 0 AND OriginType = 10 AND TaskType = 0 AND TriggerSatus = 0 and ID!=?", tenantid, firsttask.ID).
			Delete(&models.Libinventorytask{}).Error
		if err != nil {
			return resp, fmt.Errorf("删除其他盘点任务失败: %v", err)
		}
		firsttask.InventoryStartDate = InventoryStartDate
		firsttask.InventoryEndDate = InventoryEndDate
		firsttask.Interval = int64(input.Interval)
		firsttask.IsEnable = utils.BoolToUint8Slice(input.IsEnable)
		// 二期改造
		firsttask.DeviceType = &input.DeviceType
		firsttask.RobotID = nil
		firsttask.RobotName = nil
		firsttask.RobotRouterID = nil
		firsttask.RobotRouterName = nil
		err = global.OWEN_DB.Model(&models.Libinventorytask{}).Where("ID=?", firsttask.ID).Save(firsttask).Error
		if err != nil {
			return resp, fmt.Errorf("更新盘点任务失败: %v", err)
		}
		//remove  old  works
		err = global.OWEN_DB.
			Model(&models.Libinventorywork{}).
			Where("TenantId = ? AND IsDeleted = 0 AND OriginType = 10 and TaskID=?  and SendStatus=0", tenantid, firsttask.ID).
			Delete(&models.Libinventorywork{}).Error
		if err != nil {
			return resp, fmt.Errorf("删除旧任务失败: %v", err)
		}

	} else {
		firsttask.ID = utils.UUID32()
		firsttask.CreationTime = time.Now()
		firsttask.CreatorUserID = nil
		firsttask.LastModificationTime = nil
		firsttask.LastModifierUserID = nil
		firsttask.IsDeleted = utils.BoolToUint8Slice(false)
		firsttask.DeleterUserID = nil
		firsttask.DeletionTime = nil
		taskname := "间隔盘点-全盘"
		firsttask.TaskName = &taskname
		firsttask.TaskType = 0     // 0:全盘巡检
		firsttask.TriggerSatus = 0 // 0:未触发
		firsttask.InventoryStartDate = InventoryStartDate
		firsttask.InventoryEndDate = InventoryEndDate
		firsttask.Interval = int64(input.Interval)
		firsttask.IsEnable = utils.BoolToUint8Slice(input.IsEnable)
		firsttask.TenantID = int64(tenantid)
		firsttask.OriginType = 10 // 10:飞阅
		firsttask.DeviceType = &input.DeviceType
		firsttask.RobotID = nil
		firsttask.RobotName = nil
		firsttask.RobotRouterID = nil
		firsttask.RobotRouterName = nil
		err = global.OWEN_DB.Model(&models.Libinventorytask{}).Create(&firsttask).Error
		if err != nil {
			return resp, fmt.Errorf("创建盘点任务失败: %v", err)
		}
	}

	if firsttask.InventoryStartDate != nil &&
		firsttask.InventoryStartDate.Before(time.Now()) &&
		firsttask.InventoryEndDate.After(time.Now()) &&
		firsttask.InventoryEndDate != nil &&
		input.IsEnable {

		loc, _ := time.LoadLocation("Asia/Shanghai")
		now := time.Now().In(loc)

		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)

		hour := 0                  // 假设初始值
		interval := input.Interval // 假设 input.Interval 是 int 类型
		for {
			if hour >= 24 {
				break
			}
			worktime := today.Add(time.Duration(hour) * time.Hour)
			if worktime.Before(now) {
				hour += interval
				continue
			}

			if input.SysStartTime != "" {
				if worktime.Before(SysStartTime) {
					hour += interval
					continue
				}
			}
			if input.SysEndTime != "" {
				if worktime.After(SysEndTime) {
					hour += input.Interval
					continue
				}
			}

			var work models.Libinventorywork
			work.ID = utils.UUID32()
			work.CreationTime = time.Now()
			work.CreatorUserID = nil
			work.LastModificationTime = nil
			work.LastModifierUserID = nil
			work.IsDeleted = utils.BoolToUint8Slice(false)
			work.DeleterUserID = nil
			work.DeletionTime = nil
			TaskName := *firsttask.TaskName + worktime.Format("2006010215")
			work.TaskName = &TaskName
			work.TaskID = &firsttask.ID
			work.WorkTime = worktime
			hangfire := fmt.Sprintf("Collection_全盘_%d_%s", tenantid, worktime.Format("2006010215"))
			work.HangFireKey = &hangfire
			Comment := "图书全盘"
			work.Comment = &Comment
			work.SendStatus = 0 // 0:未发送
			work.TenantID = int64(tenantid)
			work.OriginType = 10 // 10:飞阅
			work.DeviceType = &input.DeviceType
			err = global.OWEN_DB.Model(&models.Libinventorywork{}).Create(&work).Error
			if err != nil {
				return resp, fmt.Errorf("创建盘点工作失败: %v", err)
			}
			hour += input.Interval
		}
	}
	return resp, nil
}

func (f *FlyReadAppService) CheckWorkTime(tenantid int, worktime time.Time) (err error) {
	if worktime.Before(time.Now()) {
		return fmt.Errorf("工作时间不能小于当前时间")
	}
	var fsets dto.FlyReadSetting
	fsets, err = ManagerGroup.frymanager.GetFlyReadSetting(tenantid)
	if err != nil {
		return err
	}
	// 获取当前日期
	today := utils.FormatInLocation("2006-01-02", worktime)
	var SysStartTime time.Time
	var SysEndTime time.Time
	if fsets.FlyStartTime != "" {
		SysStartTime, err = utils.ParseInLocation("2006-01-02 15:04", today+" "+fsets.FlyStartTime)
		if err != nil {
			return fmt.Errorf("SysStartTime 格式错误: %v", err)
		}
		if SysStartTime.After(worktime) {
			return fmt.Errorf("工作时间不能小于系统开始时间")
		}
	}
	if fsets.FlyEndTime != "" {
		SysEndTime, err = utils.ParseInLocation("2006-01-02 15:04", today+" "+fsets.FlyEndTime)
		if err != nil {
			return fmt.Errorf("SysEndTime 格式错误: %v", err)
		}
		if SysEndTime.Before(worktime) {
			return fmt.Errorf("工作时间不能大于系统结束时间")
		}
	}
	return nil
}

func (f *FlyReadAppService) CreatWork(input request.CreatWorkInput, tenantid int) (resp response.CreatWorkDto, err error) {
	if len(input.LayerIds) <= 0 {
		return resp, fmt.Errorf("图书层不能为空")
	}
	if len(strings.TrimSpace(input.DeviceType)) <= 0 {
		return resp, fmt.Errorf("设备类型不能空")
	}
	switch input.DeviceType {
	case "2":
		if len(strings.TrimSpace(input.RobotId)) <= 0 {
			return resp, fmt.Errorf("机器人ID不能空")
		}
		if len(strings.TrimSpace(input.RobotRouterId)) <= 0 {
			return resp, fmt.Errorf("机器人路由ID不能空")
		}
	case "1", "0":
		{

		}
	default:
		return resp, fmt.Errorf("设备类型只能是0,1,2")
	}

	var works []time.Time
	loc, _ := time.LoadLocation("Asia/Shanghai")
	WorkTime, err1 := utils.ParseInLocation("2006-01-02 15:04:05", input.WorkTime)
	if err1 != nil {
		return resp, fmt.Errorf("工作时间格式错误: %v", err1)
	}
	err = f.CheckWorkTime(tenantid, WorkTime)
	if err != nil {
		return resp, fmt.Errorf("工作时间不符合系统设置: %v", err)
	}
	works = append(works, WorkTime)

	hour := WorkTime.Hour()
	minutes := WorkTime.Minute()

	if len(input.WorkTimes) > 0 {
		for _, wt := range input.WorkTimes {
			worktime_temp, err2 := utils.ParseInLocation("2006-01-02 15:04:05", wt)
			if err2 != nil {
				return resp, fmt.Errorf("工作时间格式错误: %v", err2)
			}
			worktime_each := time.Date(worktime_temp.Year(), worktime_temp.Month(), worktime_temp.Day(), hour, minutes, 0, 0, loc)
			err = f.CheckWorkTime(tenantid, worktime_each)
			if err != nil {
				return resp, fmt.Errorf("工作时间不符合系统设置: %v", err)
			}
			works = append(works, worktime_each)
		}
	}
	layermap := make(map[string]models.Liblayer)
	layers, e2 := ManagerGroup.frymanager.GetEnableLayers(tenantid, true, input.DeviceType, input.RobotId, input.RobotRouterId)
	if e2 != nil {
		return resp, fmt.Errorf("获取启用层失败: %v", e2)
	}
	if len(layers) <= 0 {
		return resp, fmt.Errorf("没有启用层")
	}
	for _, ly := range input.LayerIds {
		found := false
		for _, layer := range layers {
			if layer.ID == ly || (layer.Code != nil && *layer.Code == ly) {
				found = true
				layermap[ly] = layer
				break
			}
		}
		if !found {
			return resp, fmt.Errorf("图书层 %s 不在启用层列表中", ly)
		}
	}

	var firsttask models.Libinventorytask
	firsttask.ID = utils.UUID32()
	firsttask.CreationTime = time.Now()
	firsttask.CreatorUserID = nil
	firsttask.LastModificationTime = nil
	firsttask.LastModifierUserID = nil
	firsttask.IsDeleted = utils.BoolToUint8Slice(false)
	firsttask.DeleterUserID = nil
	firsttask.DeletionTime = nil
	firsttask.TaskName = &input.TaskName
	firsttask.TaskType = 1     // 1:单次巡检
	firsttask.TriggerSatus = 1 // 0:未触发
	firsttask.InventoryStartDate = nil
	firsttask.InventoryEndDate = nil
	firsttask.Interval = 0 // 单次巡检不需要间隔
	firsttask.IsEnable = utils.BoolToUint8Slice(true)
	firsttask.TenantID = int64(tenantid)
	firsttask.OriginType = 10 // 10:飞阅
	firsttask.DeviceType = &input.DeviceType
	firsttask.RobotID = &input.RobotId
	firsttask.RobotName = &input.RobotName
	firsttask.RobotRouterID = &input.RobotRouterId
	firsttask.RobotRouterName = &input.RobotRouterName
	err = global.OWEN_DB.Model(&models.Libinventorytask{}).Create(&firsttask).Error
	if err != nil {
		return resp, fmt.Errorf("创建盘点任务失败: %v", err)
	}

	for _, worktime := range works {

		var work models.Libinventorywork
		work.ID = utils.UUID32()
		work.CreationTime = time.Now()
		work.CreatorUserID = nil
		work.LastModificationTime = nil
		work.LastModifierUserID = nil
		work.IsDeleted = utils.BoolToUint8Slice(false)
		work.DeleterUserID = nil
		work.DeletionTime = nil
		TaskName := *firsttask.TaskName + worktime.Format("2006010215")
		work.TaskName = &TaskName
		work.TaskID = &firsttask.ID
		work.WorkTime = worktime
		hangfire := fmt.Sprintf("Collection_定时盘%d_%s", tenantid, work.ID)
		work.HangFireKey = &hangfire
		Comment := ""
		work.Comment = &Comment
		work.SendStatus = 0 // 0:未发送
		work.TenantID = int64(tenantid)
		work.OriginType = 10 // 10:飞阅
		work.DeviceType = &input.DeviceType
		err = global.OWEN_DB.Model(&models.Libinventorywork{}).Create(&work).Error
		if err != nil {
			return resp, fmt.Errorf("创建盘点工作失败: %v", err)
		}
		for _, layerid := range input.LayerIds {
			var detail models.Libinventoryworkdetail
			detail.ID = utils.UUID32()
			detail.CreationTime = time.Now()
			detail.CreatorUserID = nil
			detail.TaskStatus = 0 // 0:未开始
			detail.WorkID = &work.ID
			lyone := layermap[layerid]
			detail.LayerID = &lyone.ID
			detail.LayerCode = lyone.Code
			detail.LayerName = &lyone.Name
			detail.TenantID = int64(tenantid)
			err = global.OWEN_DB.Model(&models.Libinventoryworkdetail{}).Create(&detail).Error
			if err != nil {
				return resp, fmt.Errorf("创建盘点工作详情失败: %v", err)
			}
		}
	}
	return resp, nil
}
func (f *FlyReadAppService) UpdateWork(input request.UpdateWorkInput, tenantid int) (resp response.UpdateWorkDto, err error) {
	if len(input.LayerIds) <= 0 {
		return resp, fmt.Errorf("层集合为空不允许创建任务")
	}
	if len(strings.TrimSpace(input.DeviceType)) <= 0 {
		return resp, fmt.Errorf("设备类型不能空")
	}
	switch input.DeviceType {
	case "2":
		if len(strings.TrimSpace(input.RobotId)) <= 0 {
			return resp, fmt.Errorf("机器人ID不能空")
		}
		if len(strings.TrimSpace(input.RobotRouterId)) <= 0 {
			return resp, fmt.Errorf("机器人路由ID不能空")
		}
	case "1", "0":
		{

		}
	default:
		return resp, fmt.Errorf("设备类型只能是0,1,2")
	}
	WorkTime, err1 := utils.ParseInLocation("2006-01-02 15:04:05", input.WorkTime)
	if err1 != nil {
		return resp, fmt.Errorf("工作时间格式错误: %v", err1)
	}
	err = f.CheckWorkTime(tenantid, WorkTime)
	if err != nil {
		return resp, fmt.Errorf("工作时间不符合系统设置: %v", err)
	}
	layermap := make(map[string]models.Liblayer)
	layers, e2 := ManagerGroup.frymanager.GetEnableLayers(tenantid, true, input.DeviceType, input.RobotId, input.RobotRouterId)
	if e2 != nil {
		return resp, fmt.Errorf("获取启用层失败: %v", e2)
	}
	if len(layers) <= 0 {
		return resp, fmt.Errorf("没有启用层")
	}
	for _, ly := range input.LayerIds {
		found := false
		for _, layer := range layers {
			if layer.ID == ly || (layer.Code != nil && *layer.Code == ly) {
				found = true
				layermap[ly] = layer
				break
			}
		}
		if !found {
			return resp, fmt.Errorf("图书层 %s 不在启用层列表中", ly)
		}
	}
	var work models.Libinventorywork
	err = global.OWEN_DB.Model(&models.Libinventorywork{}).Where("TenantId=? and IsDeleted=0 and Id=? ", tenantid, input.WorkId).First(&work).Error
	if err != nil || work.ID == "" {
		return resp, fmt.Errorf("当前Work不存在")
	}
	if work.TaskStatus != 0 {
		return resp, fmt.Errorf("work状态是%d,不可以编辑", work.TaskStatus)
	}
	if work.SendStatus != 0 {
		return resp, fmt.Errorf("work状态已经发送视觉盘点,不可以编辑")
	}

	var task models.Libinventorytask
	err = global.OWEN_DB.Model(&models.Libinventorytask{}).Where("TenantId=? and IsDeleted=0 and Id=? ", tenantid, work.TaskID).First(&task).Error
	if err != nil || task.ID == "" {
		return resp, fmt.Errorf("当前task不存在")
	}
	if task.TriggerSatus != 1 {
		return resp, fmt.Errorf("当前Task不是定时任务不能编辑")
	}
	task.DeviceType = &input.DeviceType
	task.TaskName = &input.WorkName
	task.RobotID = &input.RobotId
	task.RobotName = &input.RobotName
	task.RobotRouterID = &input.RobotRouterId
	task.RobotRouterName = &input.RobotRouterName
	err = global.OWEN_DB.Model(&models.Libinventorytask{}).Where("Id=?", task.ID).Updates(map[string]interface{}{
		"DeviceType":      task.DeviceType,
		"TaskName":        task.TaskName,
		"RobotID":         task.RobotID,
		"RobotName":       task.RobotName,
		"RobotRouterID":   task.RobotRouterID,
		"RobotRouterName": task.RobotRouterName,
	}).Error
	if err != nil || task.ID == "" {
		return resp, fmt.Errorf("当前task不存在")
	}
	WorkTime1 := utils.FormatInLocation("2006-01-02 15:04:05", work.WorkTime)

	if work.TaskName != &input.WorkName || WorkTime1 != input.WorkTime || work.DeviceType != &input.DeviceType {
		work.TaskName = &input.WorkName
		work.WorkTime = WorkTime
		work.DeviceType = &input.DeviceType
		err = global.OWEN_DB.Model(&models.Libinventorywork{}).Where("Id=?", work.ID).Updates(map[string]interface{}{
			"TaskName":   work.TaskName,
			"WorkTime":   work.WorkTime,
			"DeviceType": work.DeviceType,
		}).Error
	}

	var details []models.Libinventoryworkdetail
	err = global.OWEN_DB.Model(&models.Libinventoryworkdetail{}).Where("TenantId=? and WorkId=? ", tenantid, work.ID).Find(&details).Error
	if err != nil {
		return resp, fmt.Errorf("详情查询报错")
	}
	isneedupdate := false
	if len(input.LayerIds) != len(details) {
		isneedupdate = true
	}
	for _, de := range details {
		found := false
		for _, code := range input.LayerIds {
			if code == *de.LayerCode {
				found = true
				break
			}
		}
		if !found {
			isneedupdate = true
			break
		}
	}
	if isneedupdate {
		err = global.OWEN_DB.Model(&models.Libinventoryworkdetail{}).Where("TenantId=? and WorkId=? ", tenantid, work.ID).Delete(&models.Libinventoryworkdetail{}).Error
		if err != nil {
			return resp, fmt.Errorf("删除明细失败")
		}
		for _, layerid := range input.LayerIds {
			var detail models.Libinventoryworkdetail
			detail.ID = utils.UUID32()
			detail.CreationTime = time.Now()
			detail.CreatorUserID = nil
			detail.TaskStatus = 0 // 0:未开始
			detail.WorkID = &work.ID
			lyone := layermap[layerid]
			detail.LayerID = &lyone.ID
			detail.LayerCode = lyone.Code
			detail.LayerName = &lyone.Name
			detail.TenantID = int64(tenantid)
			err = global.OWEN_DB.Model(&models.Libinventoryworkdetail{}).Create(&detail).Error
			if err != nil {
				return resp, fmt.Errorf("创建盘点工作详情失败: %v", err)
			}
		}
	}
	return resp, nil
}

type detailgroup struct {
	WorkId            string
	Count             int
	ExceptionMsgCount int
}

// 任务列表
func (f *FlyReadAppService) WorkList(input request.WorkListInput, tenantid int) (resp response.PageWorkListDto, err error) {

	worksAll := global.OWEN_DB.Model(&models.Libinventorywork{}).Where("TenantID=? and IsDeleted=0 and OriginType=10", tenantid)
	if input.WorkId != "" {
		worksAll = worksAll.Where("Id=?", input.WorkId)
	}
	if input.DeviceType != "" {
		worksAll = worksAll.Where("DeviceType=?", input.DeviceType)
	}
	if input.WorkName != "" {
		worksAll = worksAll.Where("WorkName like ?", "%"+input.WorkName+"%")
	}
	if input.TriggerSatus > 0 {
		worksAll = worksAll.Where("TriggerSatus=?", input.TriggerSatus)
	}
	workquery := worksAll.Session(&gorm.Session{})
	if input.TaskStatus > 0 {
		workquery = workquery.Where("TaskStatus=?", input.TaskStatus)
	}
	workquery = workquery.Order(" CreationTime desc")
	var total int64
	workquery.Count(&total)

	var itemworks []models.Libinventorywork
	err = workquery.Offset(input.SkipCount).Limit(input.MaxResultCount).Find(&itemworks).Error
	if err != nil {
		return resp, err
	}
	if len(itemworks) <= 0 {
		return resp, nil
	}

	var items []response.WorkListDto

	var workids []string
	for _, work := range itemworks {
		workids = append(workids, work.ID)

		var TaskName string
		if work.TaskName != nil {
			TaskName = *work.TaskName
		}

		var WorkStarTime string
		if work.WorkStarTime != nil {
			WorkStarTime = utils.FormatInLocation("2006-01-02 15:04:05", *work.WorkStarTime)
		}

		var WorkEndTime string
		if work.WorkEndTime != nil {
			WorkEndTime = utils.FormatInLocation("2006-01-02 15:04:05", *work.WorkEndTime)
		}

		var Comment string
		if work.Comment != nil {
			Comment = *work.Comment
		}

		var ExceptionMsg string
		if work.ExceptionMsg != nil {
			ExceptionMsg = *work.ExceptionMsg
		}

		item := response.WorkListDto{
			WorkName:     TaskName,
			WorkId:       work.ID,
			WorkTime:     utils.FormatInLocation("2006-01-02 15:04:05", work.WorkTime),
			WorkStarTime: WorkStarTime,
			WorkEndTime:  WorkEndTime,
			TaskStatus:   int(work.TaskStatus),
			CreateTime:   utils.FormatInLocation("2006-01-02 15:04:05", work.CreationTime),
			TriggerSatus: int(work.TaskStatus),
			Comment:      Comment,
			ExceptionMsg: ExceptionMsg,
		}
		items = append(items, item)
	}

	var details []models.Libinventoryworkdetail
	err = global.OWEN_DB.Model(&models.Libinventoryworkdetail{}).Where("TenantID=? and WorkId in ?", tenantid, workids).Find(&details).Error
	if err != nil {
		return resp, err
	}

	details_map := make(map[string]*detailgroup)
	for _, d := range details {
		if _, exists := details_map[*d.WorkID]; !exists {
			ExceptionMsgCount := 0
			if d.ExceptionMsg != nil && *d.ExceptionMsg != "" {
				ExceptionMsgCount = 1
			}
			details_map[*d.WorkID] = &detailgroup{
				WorkId:            *d.WorkID,
				Count:             1,
				ExceptionMsgCount: ExceptionMsgCount,
			}
		} else {
			v := details_map[*d.WorkID]
			if d.ExceptionMsg != nil && *d.ExceptionMsg != "" {
				v.ExceptionMsgCount = v.ExceptionMsgCount + 1
			}
			v.Count = v.Count + 1
		}
	}

	for index := range items {
		if v, exist := details_map[items[index].WorkId]; exist {
			items[index].DetailsCount = v.Count
			items[index].ExceptionMsgCount = v.ExceptionMsgCount
		} else {
			items[index].DetailsCount = 0
			items[index].ExceptionMsgCount = 0
		}
	}

	resp.TotalCount = int(total)
	resp.Items = items

	var InitWorkCount int64
	worksAll.Session(&gorm.Session{}).Where("TaskStatus=0").Count(&InitWorkCount)
	resp.InitWorkCount = int(InitWorkCount)

	var InventoryCount int64
	worksAll.Session(&gorm.Session{}).Where("TaskStatus=1").Count(&InventoryCount)
	resp.InventoryCount = int(InventoryCount)

	var SucessCount int64
	worksAll.Session(&gorm.Session{}).Where("TaskStatus=2").Count(&SucessCount)
	resp.SucessCount = int(SucessCount)

	var FailtureCount int64
	worksAll.Session(&gorm.Session{}).Where("TaskStatus=3").Count(&FailtureCount)
	resp.FailtureCount = int(FailtureCount)

	return resp, nil
}

func (f *FlyReadAppService) deleteWorkDetail(workid string, tenantid int) (err error) {
	err = global.OWEN_DB.Model(&models.Libinventoryworkdetail{}).Where("TenantID=? and WorkID=?", tenantid, workid).Delete(&models.Libinventoryworkdetail{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *FlyReadAppService) DeleteWork(input request.DeleteWorkInput, tenantid int) (resp response.DeleteWorkDto, err error) {

	var work models.Libinventorywork
	err = global.OWEN_DB.Model(&models.Libinventorywork{}).Where("IsDeleted=0 and TenantID=? and ID=?", tenantid, input.Workid).First(&work).Error
	if err != nil {
		return resp, err
	}
	if work.ID == "" {
		return resp, fmt.Errorf("workid不存在")
	}
	if work.SendStatus != 0 {
		return resp, fmt.Errorf("已创建无法删除，任务1天后未执行将自动清除！")
	}
	if work.TaskStatus != 0 {
		return resp, fmt.Errorf("任务状态%d,不允许删除！", work.TaskStatus)
	}
	err = f.deleteWorkDetail(input.Workid, tenantid)
	if err != nil {
		return resp, err
	}
	err = global.OWEN_DB.Delete(&work).Error
	if err != nil {
		return resp, err
	}
	resp.Success = true
	return resp, nil
}

func (f *FlyReadAppService) GetTaskDetailAsync(workid string, tenantid int) (resp response.TaskDetail, err error) {
	var work models.Libinventorywork
	err = global.OWEN_DB.Model(&models.Libinventorywork{}).Where("TenantID=? and IsDeleted=0 and Id=?", tenantid, workid).First(&work).Error
	if err != nil {
		return resp, err
	}
	if work.TaskID == nil || *work.TaskID == "" {
		return resp, fmt.Errorf("work不存在")
	}
	var task models.Libinventorytask

	err = global.OWEN_DB.Model(&models.Libinventorytask{}).Where("TenantID=? and IsDeleted=0 and Id=?", tenantid, work.TaskID).First(&task).Error
	if err != nil {
		return resp, err
	}

	var detailscount int64
	global.OWEN_DB.Model(&models.Libinventoryworkdetail{}).Where("TenantID=? and WorkID=?", tenantid, work.ID).Count(&detailscount)

	return response.TaskDetail{
		DeviceType:      utils.SafeString(task.DeviceType),
		RobotId:         utils.SafeString(task.RobotID),
		RobotName:       utils.SafeString(task.RobotName),
		RobotRouterId:   utils.SafeString(task.RobotRouterID),
		RobotRouterName: utils.SafeString(task.RobotRouterName),
		WorkName:        utils.SafeString(work.TaskName),
		WorkId:          work.ID,
		TriggerSatus:    int(work.TaskStatus),
		WorkTime:        utils.FormatInLocation("2006-01-02 15:04:05", work.WorkTime),
		WorkStarTime:    utils.SafeTimeString(work.WorkStarTime, "2006-01-02 15:04:05"),
		WorkEndTime:     utils.SafeTimeString(work.WorkEndTime, "2006-01-02 15:04:05"),
		WorkLayerCount:  int(detailscount),
	}, nil
}

func (f *FlyReadAppService) DetailList(input request.DetailListInput, tenantid int) (resp response.PageDetailListDto, err error) {
	var details []models.Libinventoryworkdetail
	err = global.OWEN_DB.Model(&models.Libinventoryworkdetail{}).Where("TenantID=? and WorkId=?", tenantid, input.WorkId).Find(&details).Error
	if err != nil {
		return resp, fmt.Errorf("%v", err)
	}
	if len(details) <= 0 {
		return resp, nil
	}
	var detailids []string
	for _, v := range details {
		detailids = append(detailids, v.ID)
	}
	logquery := global.OWEN_DB.Model(&models.Libiteminventorylog{}).Where("TenantID=? and InventoryWorkId in ? and InventoryWorkType=1", tenantid, detailids)
	if input.LayerId != "" {
		logquery = logquery.Where("LayerId=?", input.LayerId)
	}
	if input.Barcode != "" {
		logquery = logquery.Where("Barcode=?", input.Barcode)
	}
	if input.Title != "" {
		logquery = logquery.Where("Title like ?", "%"+input.Title+"%")
	}
	if input.CallNo != "" {
		logquery = logquery.Where("CallNo=?", input.CallNo)
	}
	if input.ISBN != "" {
		logquery = logquery.Where("ISBN=?", input.ISBN)
	}
	if input.InventoryState > 0 {
		logquery = logquery.Where("InventoryState=?", input.InventoryState)
	}
	if input.ShowOff != "1" {
		logquery = logquery.Where("(InventoryState!=2 and InventoryState!=8)")
	}
	if input.Sorting != "" {
		logquery = logquery.Order(input.Sorting)
	} else {
		logquery = logquery.Order("LayerCode asc")
	}

	var total int64
	logquery.Count(&total)

	var logs []models.Libiteminventorylog
	err = logquery.Offset(input.SkipCount).Limit(input.MaxResultCount).Find(&logs).Error
	if err != nil {
		return resp, err
	}

	var layersids []string

	for _, v := range logs {
		if v.LocLayerID != nil && *v.LocLayerID != "" {
			layersids = append(layersids, *v.LocLayerID)
		}
	}

	var layers []models.Liblayer
	err = global.OWEN_DB.Model(&models.Liblayer{}).Where("TenantID=? and Id in ?", tenantid, layersids).Find(&layers).Error
	if err != nil {
		return resp, err
	}
	layers_map := make(map[string]models.Liblayer)
	for _, v := range layers {
		layers_map[v.ID] = v
	}

	for _, v := range logs {
		var dto response.DetailListDto
		dto.LayerId = utils.SafeString(v.LayerID)
		dto.LayerCode = utils.SafeString(v.LayerCode)
		dto.LayerName = utils.SafeString(v.LayerName)
		dto.Title = utils.SafeString(v.ItemTitle)
		dto.ISBN = utils.SafeString(v.ItemISBN)
		dto.Author = utils.SafeString(v.ItemAuthor)
		dto.Publisher = utils.SafeString(v.ItemPublisher)
		dto.CallNo = utils.SafeString(v.ItemCallNo)
		dto.Barcode = v.ItemBarcode
		dto.LocationName = utils.SafeString(v.LocationName)
		dto.InventoryState = int(v.InventoryState)
		dto.LocLayerName = utils.SafeString(v.LocLayerName)
		dto.LocLayerId = utils.SafeString(v.LocLayerID)
		dto.LocLayerCode = utils.SafeString(v.LocLayerCode)
		dto.Remark = utils.SafeString(v.Remark)

		if (v.LocLayerName == nil || *v.LocLayerName == "") && (v.LocLayerID != nil && *v.LocLayerID != "") {
			dto.LocLayerName = layers_map[*v.LocLayerID].Name
		}
		resp.Items = append(resp.Items, dto)
	}
	resp.TotalCount = int(total)
	var taskDetail response.TaskDetail
	taskDetail, err = f.GetTaskDetailAsync(input.WorkId, tenantid)
	if err != nil {
		return resp, err
	}
	resp.Detail = taskDetail
	return resp, nil
}

func (f *FlyReadAppService) DetailStatusList(input request.DetailStatusListInput, tenantid int) (resp response.PageDetailStatusListDto, err error) {
	var details []models.Libinventoryworkdetail
	details_query := global.OWEN_DB.Model(&models.Libinventoryworkdetail{}).Where("TenantID=? and WorkId=?", tenantid, input.WorkId)
	if input.TaskStatus >= 0 {
		details_query = details_query.Where("TaskStatus=?", input.TaskStatus)
	}
	var total int64
	details_query.Count(&total)

	details_query = details_query.Order("LayerCode asc")

	err = details_query.Offset(input.SkipCount).Limit(input.MaxResultCount).Find(&details).Error
	if err != nil {
		return resp, err
	}
	if len(details) <= 0 {
		return resp, nil
	}

	if input.TaskStatus < 0 {
		isneedUpdate := true
		for _, v := range details {
			if !(v.TaskStatus == 2 || v.TaskStatus == 3) {
				isneedUpdate = false
				break
			}
		}
		if isneedUpdate {
			var work models.Libinventorywork
			err = global.OWEN_DB.Model(&models.Libinventorywork{}).Where("TenantID=? and IsDeleted=0 and Id=?", tenantid, input.WorkId).First(&work).Error
			if err != nil {
				return resp, nil
			}
			if work.ID != "" && work.TaskStatus != 2 && work.TriggerSatus != 2 {
				if work.TaskStatus != 0 {

					var detail models.Libinventoryworkdetail
					err = global.OWEN_DB.Model(&models.Libinventoryworkdetail{}).Where("TenantID=? and WorkId=? and  (ExceptionMsg!=null and ExceptionMsg!='')", tenantid, input.WorkId).Order("CreationTime desc").First(&detail).Error
					if err != nil {
						fmt.Println("detail not  found")
					}
					err = global.OWEN_DB.Model(&models.Libinventorywork{}).
						Where("TenantID=? AND IsDeleted=0 AND Id=?", tenantid, input.WorkId).
						Updates(map[string]interface{}{
							"TaskStatus":   2, // 替换为你要更新的值
							"ExceptionMsg": utils.SafeString(detail.ExceptionMsg),
							// 其他需要更新的字段
						}).Error
					if err != nil {
						return resp, err
					}
				}
			}
		}
	}

	resp.TotalCount = int(total)

	for _, v := range details {
		var dto response.DetailStatusListDto
		dto.LayerId = utils.SafeString(v.LayerID)
		dto.LayerCode = utils.SafeString(v.LayerCode)
		dto.LayerName = utils.SafeString(v.LayerName)
		dto.TaskStatus = int(v.TaskStatus)
		dto.ExceptionMsg = utils.SafeString(v.ExceptionMsg)
		dto.Remark = utils.SafeString(v.Remark)
		resp.Items = append(resp.Items, dto)
	}

	var taskDetail response.TaskDetail
	taskDetail, err = f.GetTaskDetailAsync(input.WorkId, tenantid)
	if err != nil {
		return resp, err
	}
	resp.Detail = taskDetail
	return resp, nil
}

func (f *FlyReadAppService) InventoryMonthList(input request.InventoryMonthListInput, tenantid int) (resp []response.InventoryMonthListDto, err error) {

	// 取当前日期部分（清除时间）
	dateOnly := utils.NowDateLocation()
	// 往前推一个月
	prevMonth := dateOnly.AddDate(0, -1, 0)

	var statlist []models.Libinventorystat
	err = global.OWEN_DB.Model(&models.Libinventorystat{}).Where("TenantID=? and OriginType=10 and StatType=0 and StatDate>?", tenantid, prevMonth).Order("StatDate desc,InventoryState asc").Find(&statlist).Error
	if err != nil {
		return resp, err
	}
	if len(statlist) <= 0 {
		return resp, nil
	}

	for _, v := range statlist {
		var dto response.InventoryMonthListDto
		dto.DateTime = utils.FormatInLocation("2006-01-02 15:04:05", v.StatDate)
		dto.Count = int(v.InventoryCount)
		dto.InventoryState = int(v.InventoryState)
		resp = append(resp, dto)
	}

	return resp, nil
}
func (f *FlyReadAppService) BooksNewIndex(input request.BooksNewIndexInput, tenantid int) (resp response.BooksNewIndexDto, err error) {
	var totalCount int64
	global.OWEN_DB.Model(&models.Libitem{}).Where("TenantID=? and IsDeleted=0", tenantid).Count(&totalCount)
	var dto_resp dto.QueryShelfNumDto
	dto_resp, err = ManagerGroup.frymanager.QueryShelfNum(tenantid)
	if err != nil {
		return resp, err
	}
	resp.TotalCount = int(totalCount)
	resp.Seg_num = dto_resp.Obj.Seg_num
	resp.Ocr_num = dto_resp.Obj.Ocr_num
	resp.Match_num = dto_resp.Obj.Match_num
	resp.Confidence_num = dto_resp.Obj.Confidence_num
	resp.Unconfidence_num = dto_resp.Obj.Unconfidence_num

	return resp, nil
}

func (f *FlyReadAppService) GetNotHitRank(req request.GetNotHitRankInput, tenantid int) (resp []response.GetNotHitRankDto, err error) {
	var dto_reso dto.QueryNotHitRankDto
	dto_reso, err = ManagerGroup.frymanager.QueryNotHitRank(req.Query_limit, tenantid)
	if err != nil {
		return resp, err
	}
	for _, k := range dto_reso.Obj {
		var one response.GetNotHitRankDto
		one.GetNotHitListDto = k
		resp = append(resp, one)
	}
	return resp, nil
}

type FaultLayerStat struct {
	LayerName string
	LayerId   string
	Count     int
	Time      time.Time
}

func (f *FlyReadAppService) GetFaultRank(input request.GetFaultRankInput, tenantid int) (resp []response.GetFaultRankDto, err error) {
	var faultList []FaultLayerStat
	sql := `
SELECT 
    MAX(LayerName) AS LayerName,
    LayerId,
    COUNT(*) AS Count,
    MAX(COALESCE(LastModificationTime, CreationTime)) AS Time
FROM libiteminventoryinfo
WHERE 
    InventoryState = 3 AND
    LayerId IS NOT NULL AND TRIM(LayerId) != ''
GROUP BY LayerId
ORDER BY Count DESC
LIMIT ?
`
	err = global.OWEN_DB.Raw(sql, input.Count).Scan(&faultList).Error
	if err != nil {
		return resp, err
	}
	for _, v := range faultList {
		var dto response.GetFaultRankDto
		dto.LayerId = v.LayerId
		dto.LayerName = v.LayerName
		dto.Count = v.Count
		dto.Time = utils.FormatInLocation("2006-01-02 15:04:05", v.Time)
		resp = append(resp, dto)
	}
	return resp, nil
}
func (f *FlyReadAppService) BooksIndex(input request.BooksIndexInput, tenantid int) (resp response.BooksIndexDto, err error) {

	var ItemsTotal int64
	global.OWEN_DB.Model(&models.Libitem{}).Where("TenantID=? and IsDeleted=0", tenantid).Count(&ItemsTotal)

	var OffBooksTotal int64
	global.OWEN_DB.Model(&models.Libiteminventorylog{}).Where("TenantID=? and InventoryWorkType=3", tenantid).Count(&OffBooksTotal)

	var OnbooksTotal int64
	global.OWEN_DB.Model(&models.Libiteminventoryinfo{}).Where("TenantID=? and (InventoryState=1 or InventoryState=3)", tenantid).Count(&OnbooksTotal)

	var BookExceptionTotal int64
	global.OWEN_DB.Model(&models.Libiteminventoryinfo{}).Where("TenantID=? and ExceptionMsg!=null and ExceptionMsg!=''", tenantid).Count(&BookExceptionTotal)

	resp.ItemsTotal = int(ItemsTotal)
	resp.OffBooksTotal = int(OffBooksTotal)
	resp.OnbooksTotal = int(OnbooksTotal)
	resp.BookExceptionTotal = int(BookExceptionTotal)
	return resp, nil
}

type logCount struct {
	ItemBarcode string
	Count       int
	ItemTitle   string
}

func (f *FlyReadAppService) BookRankIndex(input request.BookRankIndexInput, tenantid int) (resp []response.BookRankIndexDto, err error) {
	today := utils.NowDateLocation()
	dtEnd := time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, today.Location())
	dtStart := dtEnd.AddDate(0, -1, 0)
	if input.CountType == "2" {
		dtStart, dtEnd = utils.GetStatRange()
	}
	if input.IsOcrBarcode {
		var OffBooksCurrent []logCount
		sql := `
		select ItemBarcode,Count(1) as Count from  libiteminventorylog where TenantId=? AND InventoryWorkType=3  AND CreationTime>=?
GROUP BY ItemBarcode 
ORDER BY Count DESC
LIMIT ?
		`
		err = global.OWEN_DB.Raw(sql, tenantid, dtEnd, input.Count).Scan(&OffBooksCurrent).Error
		if err != nil {
			return resp, err
		}
		var barcodes []string
		if len(OffBooksCurrent) > 0 {
			for _, v := range OffBooksCurrent {
				//barcode := strings.TrimPrefix(v.ItemBarcode, "*")
				barcodes = append(barcodes, v.ItemBarcode)
			}
			var OffBooksLast []logCount
			OffBooksLast_map := make(map[string]int)
			sql = `
		select ItemBarcode,Count(1) as Count from  libiteminventorylog where TenantId=? AND InventoryWorkType=3  
		AND CreationTime>=? AND  CreationTime<? AND ItemBarcode IN ? 
GROUP BY ItemBarcode 
ORDER BY Count DESC
LIMIT ?
		`
			err = global.OWEN_DB.Raw(sql, tenantid, dtStart, dtEnd, barcodes, input.Count).Scan(&OffBooksLast).Error
			if err != nil {

			}
			if len(OffBooksLast) > 0 {
				for _, v := range OffBooksLast {
					OffBooksLast_map[v.ItemBarcode] = v.Count
				}
			}

			var barcodes_book []string
			for _, v := range barcodes {
				barcode := strings.TrimPrefix(v, "*")
				barcodes_book = append(barcodes_book, barcode)
			}

			var books []models.Libitem
			books_map := make(map[string]models.Libitem)
			err = global.OWEN_DB.Model(&models.Libitem{}).Where("TenantID=? and IsDeleted=0 and Barcode in ?", tenantid, barcodes_book).Find(&books).Error
			if err != nil {

			}
			if len(books) > 0 {
				for _, v := range books {
					books_map[v.Barcode] = v
				}
			}

			// 排序：按 Count 倒序
			sort.Slice(OffBooksCurrent, func(i, j int) bool {
				return OffBooksCurrent[i].Count > OffBooksCurrent[j].Count // 逆序
			})

			for i := 0; i < len(OffBooksCurrent); i++ {
				item := OffBooksCurrent[i]
				var dto response.BookRankIndexDto
				book_barcode := strings.TrimPrefix(item.ItemBarcode, "*")
				dto.Title = books_map[book_barcode].Title
				dto.Count = item.Count
				dto.LastCount = OffBooksLast_map[item.ItemBarcode]
				dto.Index = i + 1
				dto.CallNo = utils.SafeString(books_map[book_barcode].CallNo)
				resp = append(resp, dto)
			}
		}

	} else {
		var OffBooksCurrent []logCount
		sql := `
		select ItemTitle,Count(1) as Count from  libiteminventorylog where TenantId=? AND InventoryWorkType=3  AND CreationTime>=?
GROUP BY ItemTitle 
ORDER BY Count DESC
LIMIT ?
		`
		err = global.OWEN_DB.Raw(sql, tenantid, dtEnd, input.Count).Scan(&OffBooksCurrent).Error
		if err != nil {
			return resp, err
		}
		if len(OffBooksCurrent) > 0 {
			var titles []string
			for _, v := range OffBooksCurrent {
				titles = append(titles, v.ItemTitle)
			}
			var OffBooksLast []logCount
			OffBooksLast_map := make(map[string]int)
			sql = `
		select ItemTitle,Count(1) as Count from  libiteminventorylog where TenantId=? AND InventoryWorkType=3  
		AND CreationTime>=? AND  CreationTime<? AND ItemTitle IN ? 
GROUP BY ItemTitle 
ORDER BY Count DESC
LIMIT ?
		`
			err = global.OWEN_DB.Raw(sql, tenantid, dtStart, dtEnd, titles, input.Count).Scan(&OffBooksLast).Error
			if err != nil {

			}
			if len(OffBooksLast) > 0 {
				for _, v := range OffBooksLast {
					OffBooksLast_map[v.ItemTitle] = v.Count
				}
			}

			// 排序：按 Count 倒序
			sort.Slice(OffBooksCurrent, func(i, j int) bool {
				return OffBooksCurrent[i].Count > OffBooksCurrent[j].Count // 逆序
			})

			for i := 0; i < len(OffBooksCurrent); i++ {
				item := OffBooksCurrent[i]
				var dto response.BookRankIndexDto
				dto.Title = item.ItemTitle
				dto.Count = item.Count
				dto.LastCount = OffBooksLast_map[item.ItemTitle]
				dto.Index = i + 1
				resp = append(resp, dto)
			}
		}

	}
	return resp, nil
}

func (f *FlyReadAppService) InventoryFlyReadIndex(input request.InventoryFlyReadIndexInput, tenantid int) (resp response.InventoryFlyReadIndexDto, err error) {
	current, e1 := ManagerGroup.frymanager.InventoryHis(false, tenantid)
	if e1 != nil {
		return resp, e1
	}
	last, e2 := ManagerGroup.frymanager.InventoryHis(true, tenantid)
	if e2 != nil {
		return resp, e2
	}

	resp.Current = response.InventoryFlyReadHisDto{
		CreateTime:  current.Obj.CreateTime,
		UpdateTime:  current.Obj.UpdateTime,
		WorkId:      current.Obj.WorkId,
		CheckType:   current.Obj.CheckType,
		CheckAll:    current.Obj.CheckAll,
		DeviceType:  current.Obj.DeviceType,
		Status:      current.Obj.Status,
		TaskAllNum:  current.Obj.TaskAllNum,
		FinishNum:   current.Obj.FinishNum,
		BookCount:   current.Obj.BookCount,
		AssertCount: current.Obj.AssertCount,
	}
	resp.Last = response.InventoryFlyReadHisDto{
		CreateTime:  last.Obj.CreateTime,
		UpdateTime:  last.Obj.UpdateTime,
		WorkId:      last.Obj.WorkId,
		CheckType:   last.Obj.CheckType,
		CheckAll:    last.Obj.CheckAll,
		DeviceType:  last.Obj.DeviceType,
		Status:      last.Obj.Status,
		TaskAllNum:  last.Obj.TaskAllNum,
		FinishNum:   last.Obj.FinishNum,
		BookCount:   last.Obj.BookCount,
		AssertCount: last.Obj.AssertCount,
	}
	if resp.Last.WorkId != "" {
		var work models.Libinventorywork
		err = global.OWEN_DB.Model(&models.Libinventorywork{}).Where("TenantID=? and IsDeleted=0 and Id=?", tenantid, resp.Last.WorkId).First(&work).Error
		if err != nil {
			return resp, nil
		}
		var details []models.Libinventoryworkdetail
		err = global.OWEN_DB.Model(&models.Libinventoryworkdetail{}).Where("TenantID=?  and WorkID=?", tenantid, resp.Last.WorkId).Find(&details).Error
		if err != nil {
			return resp, nil
		}
		if len(details) <= 0 {
			return resp, err
		}
		var ids []string
		for _, v := range details {
			ids = append(ids, v.ID)
		}
		var logs []models.Libiteminventorylog
		err = global.OWEN_DB.Model(&models.Libiteminventorylog{}).Where("TenantID=? and OriginType=10 and InventoryWorkType=1  and Id in ?", tenantid, ids).Find(&logs).Error
		if err != nil {
			return resp, nil
		}
		if len(logs) <= 0 {
			return resp, nil
		}
		logs_map := make(map[int]int)
		for _, v := range logs {
			logs_map[int(v.InventoryState)]++
		}

		resp.Last.OnNums = logs_map[1]
		resp.Last.FaultNums = logs_map[3]
		resp.Last.OffNums = logs_map[2]
		resp.Last.OnAndFaultNums = resp.Last.FaultNums + resp.Last.OnNums
		resp.Last.TotalNums = resp.Last.OnAndFaultNums
		if resp.Last.UpdateTime == "" && work.WorkEndTime != nil {
			resp.Last.UpdateTime = utils.FormatInLocation("2006-01-02 15:04:05", *work.WorkEndTime)
		}
		if resp.Last.CreateTime == "" && work.WorkStarTime != nil {
			resp.Last.CreateTime = utils.FormatInLocation("2006-01-02 15:04:05", *work.WorkStarTime)
		}

	}
	return resp, nil
}

func (f *FlyReadAppService) GetStructTreeList(tenantid int) (resp []response.StructDto, err error) {

	var structlist []models.Libstruct
	err = global.OWEN_DB.Model(&models.Libstruct{}).Where("TenantID=? and IsDeleted=0", tenantid).Find(&structlist).Error
	if err != nil {
		return resp, err
	}
	var struct_map_bulid []models.Libstruct
	var struct_map_floor []models.Libstruct
	var struct_map_room []models.Libstruct

	for _, v := range structlist {
		if v.FloorNo == 0 && v.RoomNo == 0 {
			//build
			struct_map_bulid = append(struct_map_bulid, v)
		} else if v.FloorNo != 0 && v.RoomNo == 0 {
			//floor
			struct_map_floor = append(struct_map_floor, v)
		} else if v.FloorNo != 0 && v.RoomNo != 0 {
			//room
			struct_map_room = append(struct_map_room, v)
		}
	}
	for _, build := range struct_map_bulid {
		var dto response.StructDto
		dto.Libstruct = build //1级
		buildcode := fmt.Sprintf("%02d", build.BuildNo)
		for _, floor := range struct_map_floor {
			floorcode := fmt.Sprintf("%02d%02d%02d", floor.BuildNo, floor.FloorNo, floor.RoomNo)
			if strings.HasPrefix(floorcode, buildcode) {
				var dto_floor response.StructDto
				dto_floor.Libstruct = floor
				dto.Children = append(dto.Children, dto_floor) //二级
			}
		}

		for index := range dto.Children {
			floorcode := fmt.Sprintf("%02d%02d", dto.Children[index].BuildNo, dto.Children[index].FloorNo)
			for _, room := range struct_map_room {
				roomcode := fmt.Sprintf("%02d%02d%02d", room.BuildNo, room.FloorNo, room.RoomNo)
				if strings.HasPrefix(roomcode, floorcode) {
					var dto_room response.StructDto
					dto_room.Libstruct = room
					dto.Children[index].Children = append(dto.Children[index].Children, dto_room) //三级
				}
			}
		}
		resp = append(resp, dto)
	}
	return resp, nil
}
