package services

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/models/dto"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"GINOWEN/utils"
	"fmt"
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
	WorkTime1 := utils.FormatInLocation("2016-01-02 15:04:05", work.WorkTime)

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
			WorkStarTime = utils.FormatInLocation("2016-01-02 15:04:05", *work.WorkStarTime)
		}

		var WorkEndTime string
		if work.WorkEndTime != nil {
			WorkEndTime = utils.FormatInLocation("2016-01-02 15:04:05", *work.WorkEndTime)
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
			WorkTime:     utils.FormatInLocation("2016-01-02 15:04:05", work.WorkTime),
			WorkStarTime: WorkStarTime,
			WorkEndTime:  WorkEndTime,
			TaskStatus:   int(work.TaskStatus),
			CreateTime:   utils.FormatInLocation("2016-01-02 15:04:05", work.CreationTime),
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
