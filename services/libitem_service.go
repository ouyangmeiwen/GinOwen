package services

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/models/request"
	"GINOWEN/models/response"

	"GINOWEN/utils"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/xuri/excelize/v2"
)

type LibItemService struct {
}

// 导入excel 到数据库
func (LibItemService) ImportExcel(req request.ImportExcelInput) (resp response.ImportExcelDto, err error) {
	begin := time.Now()
	fmt.Println("进入方法的时间" + begin.Format("2006-01-02 15:04:05"))
	if !utils.FileExists(req.Path) {
		return resp, errors.New("文件不存在！")
	}
	// 打开 Excel 文件
	f, err := excelize.OpenFile(req.Path)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	// 获取所有工作表的名称
	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		log.Fatalf("failed to get sheet name")
	}
	// 获取工作表的所有行
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatalf("failed to get rows: %v", err)
	}
	dic := make(map[string]int)
	datas := []models.LibItem{}
	for i, row := range rows {
		// 跳过表头
		if i == 0 {
			for jindex, sv := range row {
				if len(sv) > 0 {
					dic[sv] = jindex //这里将抓换["题名":"1"，"作者":"2"]
				}
			}
			continue
		}
		if len(row) <= 0 {
			continue
		}
		obj := models.LibItem{}
		obj.Id = utils.GenerateGUID()
		obj.CreationTime = time.Now()
		obj.CreatorUserId = sql.NullInt64{Valid: false, Int64: 0}
		obj.LastModificationTime = sql.NullTime{Valid: false, Time: time.Now()}
		obj.LastModifierUserId = sql.NullInt64{Valid: false, Int64: 0}
		obj.IsDeleted = []uint8{0}
		obj.DeletionTime = sql.NullTime{Valid: false, Time: time.Now()}
		obj.DeleterUserId = sql.NullInt64{Valid: false, Int64: 0}
		obj.InfoId = sql.NullString{Valid: false, String: ""}
		println(row[0])
		if len(req.Title) > 0 {
			if value, ok := dic[req.Title]; ok && (len(row) > value) {
				rowval := utils.ConvertToUTF8(row[value])
				utf8_str := utils.SafeSubstring(rowval, 0, 100)
				obj.Title = utils.TrimSpaces(utf8_str)
			}
		}
		if len(obj.Title) <= 0 {
			obj.Title = ""
		}
		if len(req.Author) > 0 {
			if value, ok := dic[req.Author]; ok && (len(row) > value) {
				rowval := utils.ConvertToUTF8(row[value])
				utf8_str := utils.SafeSubstring(rowval, 0, 100)
				utf8_str = utils.TrimSpaces(utf8_str)
				obj.Author = sql.NullString{Valid: true, String: utf8_str}
			}
		}
		if len(req.Barcode) > 0 {
			if value, ok := dic[req.Barcode]; ok && (len(row) > value) {
				rval := row[value]
				obj.Barcode = utils.RemoveBrackets(rval)
			}
		}
		if len(obj.Barcode) <= 0 {
			obj.Barcode = ""
		}
		obj.IsEnable = []uint8{1}
		if len(req.CallNo) > 0 {
			if value, ok := dic[req.CallNo]; ok && (len(row) > value) {
				obj.CallNo = sql.NullString{Valid: true, String: row[value]}
			}
		}
		obj.PreCallNo = sql.NullString{Valid: false, String: ""}
		if len(req.CatalogCode) > 0 {
			if value, ok := dic[req.CatalogCode]; ok && (len(row) > value) {
				obj.CatalogCode = sql.NullString{Valid: true, String: row[value]}
			}
		}
		obj.ItemState = 3
		obj.PressmarkId = sql.NullString{Valid: false, String: ""}
		obj.PressmarkName = sql.NullString{Valid: false, String: ""}
		obj.LocationId = sql.NullString{Valid: false, String: ""}
		if len(req.Locationname) > 0 {
			if value, ok := dic[req.Locationname]; ok && (len(row) > value) {
				obj.LocationName = sql.NullString{Valid: true, String: row[value]}
			}
		}
		obj.BookBarcode = sql.NullString{Valid: false, String: ""}
		if len(req.ISBN) > 0 {
			if value, ok := dic[req.ISBN]; ok && (len(row) > value) {
				obj.ISBN = sql.NullString{Valid: true, String: row[value]}
			}
		}
		obj.PubNo = sql.NullInt16{Valid: false, Int16: 0}
		if len(req.Publisher) > 0 { //如果最后一位没有是不会初始化数组的
			if value, ok := dic[req.Publisher]; ok && (len(row) > value) {
				rowval := utils.ConvertToUTF8(row[value])
				utf8_str := utils.SafeSubstring(rowval, 0, 200)
				utf8_str = utils.TrimSpaces(utf8_str)
				obj.Publisher = sql.NullString{Valid: true, String: utf8_str}
			}
		}
		if len(req.PubDate) > 0 {
			if value, ok := dic[req.PubDate]; ok && (len(row) > value) {
				utf8_str := utils.SafeSubstring(row[value], 0, 16)
				obj.PubDate = sql.NullString{Valid: true, String: utf8_str}
			}
		}
		if len(req.Price) > 0 {
			if value, ok := dic[req.Price]; ok && (len(row) > value) {
				obj.Price = sql.NullString{Valid: true, String: row[value]}
			}
		}
		if len(req.Pages) > 0 {
			if value, ok := dic[req.Pages]; ok && (len(row) > value) {
				obj.Pages = sql.NullString{Valid: true, String: row[value]}
			}
		}
		obj.Summary = sql.NullString{Valid: false, String: ""}
		obj.ItemType = 1
		obj.Remark = sql.NullString{Valid: true, String: "导入数据"}
		obj.OriginType = 1
		obj.CreateType = 1
		obj.TenantId = req.Tenantid
		//global.GVA_DB.Model(&gpi.LibItem{}).Create(obj)
		datas = append(datas, obj)
	}
	fmt.Println("总共插入:" + strconv.Itoa(len(datas)))
	// 分批插入数据
	batchSize := 1000
	var db = global.OWEN_DB //结构绑定数据库对象
	var wg sync.WaitGroup
	utils.BatchInsert(db, &wg, datas, batchSize)
	wg.Wait()
	end := time.Now()
	fmt.Println("结束方法的时间" + end.Format("2006-01-02 15:04:05"))
	fmt.Println("插入" + strconv.Itoa(len(datas)) + ",总共消耗:" + strconv.FormatFloat(end.Sub(begin).Seconds(), 'f', 2, 64))
	return resp, err
}
