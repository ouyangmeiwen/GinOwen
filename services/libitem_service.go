package services

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/models/request"
	"GINOWEN/models/response"

	"GINOWEN/utils"
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
	datas := []models.Libitem{}
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
		obj := models.Libitem{}
		obj.ID = utils.GenerateGUID()
		CreationTime := time.Now()
		obj.CreationTime = &CreationTime
		obj.CreatorUserID = nil
		obj.LastModificationTime = nil
		obj.LastModifierUserID = nil
		obj.IsDeleted = []uint8{0}
		obj.DeletionTime = nil
		obj.DeleterUserID = nil
		obj.InfoID = nil
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
				obj.Author = &utf8_str
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
				obj.CallNo = &row[value]
			}
		}
		obj.PreCallNo = nil
		if len(req.CatalogCode) > 0 {
			if value, ok := dic[req.CatalogCode]; ok && (len(row) > value) {
				obj.CatalogCode = &row[value]
			}
		}
		obj.ItemState = 3
		obj.PressmarkID = nil
		obj.PressmarkName = nil
		obj.LocationID = nil
		if len(req.Locationname) > 0 {
			if value, ok := dic[req.Locationname]; ok && (len(row) > value) {
				rowval := utils.ConvertToUTF8(row[value])
				utf8_str := utils.TrimSpaces(rowval)
				obj.LocationName = &utf8_str
			}
		}
		obj.BookBarcode = nil
		if len(req.ISBN) > 0 {
			if value, ok := dic[req.ISBN]; ok && (len(row) > value) {
				obj.ISBN = &row[value]
			}
		}
		obj.PubNo = nil
		if len(req.Publisher) > 0 { //如果最后一位没有是不会初始化数组的
			if value, ok := dic[req.Publisher]; ok && (len(row) > value) {
				rowval := utils.ConvertToUTF8(row[value])
				utf8_str := utils.SafeSubstring(rowval, 0, 200)
				utf8_str = utils.TrimSpaces(utf8_str)
				obj.Publisher = &utf8_str
			}
		}
		if len(req.PubDate) > 0 {
			if value, ok := dic[req.PubDate]; ok && (len(row) > value) {
				utf8_str := utils.SafeSubstring(row[value], 0, 16)
				obj.PubDate = &utf8_str
			}
		}
		if len(req.Price) > 0 {
			if value, ok := dic[req.Price]; ok && (len(row) > value) {
				obj.Price = &row[value]
			}
		}
		if len(req.Pages) > 0 {
			if value, ok := dic[req.Pages]; ok && (len(row) > value) {
				obj.Pages = &row[value]
			}
		}
		obj.Summary = nil
		obj.ItemType = 1
		remark := "GO导入"
		obj.Remark = &remark
		obj.OriginType = 1
		obj.CreateType = 1
		obj.TenantID = int64(req.Tenantid)
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
func (LibItemService) ImportExcel2(req request.ImportExcel2Input) (resp response.ImportExcelDto, err error) {
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
	datas := []models.Libitem{}
	for i, row := range rows {
		// 跳过表头
		if i == 0 {
			continue
		}
		if len(row) <= 0 {
			continue
		}
		obj := models.Libitem{}
		obj.ID = utils.GenerateGUID()
		CreationTime := time.Now()
		obj.CreationTime = &CreationTime
		obj.CreatorUserID = nil
		obj.LastModificationTime = nil
		obj.LastModifierUserID = nil
		obj.IsDeleted = []uint8{0}
		obj.DeletionTime = nil
		obj.DeleterUserID = nil
		obj.InfoID = nil
		println(row[0])
		if len(row) > req.Title && req.Title >= 0 {
			rowval := utils.ConvertToUTF8(row[req.Title])
			utf8_str := utils.SafeSubstring(rowval, 0, 100)
			obj.Title = utils.TrimSpaces(utf8_str)
		}
		if len(obj.Title) <= 0 {
			obj.Title = ""
		}
		if len(row) > req.Author && req.Author >= 0 {
			rowval := utils.ConvertToUTF8(row[req.Author])
			utf8_str := utils.SafeSubstring(rowval, 0, 100)
			utf8_str = utils.TrimSpaces(utf8_str)
			obj.Author = &utf8_str

		}
		if len(row) > req.Barcode && req.Barcode >= 0 {
			rowval := utils.ConvertToUTF8(row[req.Barcode])
			utf8_str := utils.SafeSubstring(rowval, 0, 30)
			obj.Barcode = utils.RemoveBrackets(utils.TrimSpaces(utf8_str))
		}

		if len(obj.Barcode) <= 0 {
			obj.Barcode = ""
		}
		obj.IsEnable = []uint8{1}
		if len(row) > req.CallNo && req.CallNo >= 0 {
			rowval := utils.ConvertToUTF8(row[req.CallNo])
			utf8_str := utils.TrimSpaces(rowval)
			obj.CallNo = &utf8_str
		}
		obj.PreCallNo = nil
		if len(row) > req.CatalogCode && req.CatalogCode >= 0 {
			rowval := utils.ConvertToUTF8(row[req.CatalogCode])
			utf8_str := utils.TrimSpaces(rowval)
			obj.CatalogCode = &utf8_str
		}
		obj.ItemState = 3
		obj.PressmarkID = nil
		obj.PressmarkName = nil
		obj.LocationID = nil
		if len(row) > req.Locationname && req.Locationname >= 0 {
			rowval := utils.ConvertToUTF8(row[req.Locationname])
			utf8_str := utils.TrimSpaces(rowval)
			obj.LocationName = &utf8_str
		}

		obj.BookBarcode = nil
		if len(row) > req.ISBN && req.ISBN >= 0 {
			rowval := utils.ConvertToUTF8(row[req.ISBN])
			utf8_str := utils.TrimSpaces(rowval)
			obj.ISBN = &utf8_str
		}

		obj.PubNo = nil
		if len(row) > req.Publisher && req.Publisher >= 0 {
			rowval := utils.ConvertToUTF8(row[req.Publisher])
			utf8_str := utils.SafeSubstring(rowval, 0, 200)
			utf8_str = utils.TrimSpaces(rowval)
			obj.Publisher = &utf8_str
		}
		if len(row) > req.PubDate && req.PubDate >= 0 {
			rowval := utils.ConvertToUTF8(row[req.PubDate])
			utf8_str := utils.TrimSpaces(rowval)
			obj.PubDate = &utf8_str
		}

		if len(row) > req.Price && req.Price >= 0 {
			rowval := utils.ConvertToUTF8(row[req.Price])
			utf8_str := utils.TrimSpaces(rowval)
			obj.Price = &utf8_str
		}
		if len(row) > req.Pages && req.Pages >= 0 {
			rowval := utils.ConvertToUTF8(row[req.Pages])
			utf8_str := utils.TrimSpaces(rowval)
			obj.Pages = &utf8_str
		}
		obj.Summary = nil
		obj.ItemType = 1
		remark := "GO导入"
		obj.Remark = &remark
		obj.OriginType = 1
		obj.CreateType = 1
		obj.TenantID = int64(req.Tenantid)
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
