package services

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"GINOWEN/utils"
)

type LibrowService struct {
}

func (LibrowService) QueryRows(input request.QueryRowInput) (resp response.QueryRowDto, err error) {
	// Implementation of the method to query rows from the library
	// This will typically involve database operations to fetch the data
	// based on the input criteria provided in QueryRowInput.

	// var layerstest []models.Libitem
	// global.OWEN_DB.Model(&models.Libitem{}).Find(&layerstest)

	// // 并发过滤
	// start := time.Now()
	// filteredParallel := parallel.ParallelFilter(layerstest, 1, func(ly models.Libitem) bool {
	// 	return strings.Contains(ly.Title, "4")
	// })
	// durationParallel := time.Since(start)
	// fmt.Println("ParallelFilter duration:", durationParallel)
	// fmt.Println("ParallelFilter count:", len(filteredParallel))

	// // 普通 for 循环过滤
	// start = time.Now()
	// var filteredSerial []models.Libitem
	// for _, ly := range layerstest {
	// 	if strings.Contains(ly.Title, "4") {
	// 		filteredSerial = append(filteredSerial, ly)
	// 	}
	// }
	// durationSerial := time.Since(start)
	// fmt.Println("Serial for-loop duration:", durationSerial)
	// fmt.Println("Serial count:", len(filteredSerial))
	islimitquery := false
	var rows []models.Librow
	row_query := global.OWEN_DB.Model(&models.Librow{})
	if input.RowNo != nil && *input.RowNo > 0 {
		row_query = row_query.Where("RowNo = ?", *input.RowNo)
		islimitquery = true
	}
	if input.StructCode != nil && len(*input.StructCode) > 0 {
		row_query = row_query.Where("Code LIKE ?", *input.StructCode+"%")
		islimitquery = true
	}
	row_query = row_query.Where("IsDeleted = ?", 0).Order("RowNo ASC")
	err = row_query.Find(&rows).Error
	if err == nil {
		var rowids []string
		for _, row := range rows {
			rowids = append(rowids, row.ID)
		}
		var shelfs []models.Libshelf
		shelf_query := global.OWEN_DB.Model(&models.Libshelf{}).Where("IsDeleted = ? and IsEnable= ?", 0, 1)
		if islimitquery {
			shelf_query = shelf_query.Where("RowIdentity IN ?", rowids)
		}
		err = shelf_query.Order("Code ASC").Find(&shelfs).Error
		var shelfids []string
		if err == nil {

			for _, shelf := range shelfs {
				shelfids = append(shelfids, shelf.ID)
			}
			var layers []models.Liblayer
			layer_query := global.OWEN_DB.Model(&models.Liblayer{}).Where("IsDeleted = ? and IsEnable= ?", 0, 1)
			if islimitquery {
				layer_query = layer_query.Where("ShelfId IN ?", shelfids)
			}
			err = layer_query.Order("LayerNo ASC").Find(&layers).Error
			layer_map := make(map[string][]response.LiblayerDto)
			if err == nil {
				for _, layer := range layers {
					layer_map[layer.ShelfID] = append(layer_map[layer.ShelfID], response.LiblayerDto{
						Liblayer:  layer,
						IsDeleted: utils.Uint8slicetoboolslice(layer.IsDeleted),
						IsEnable:  utils.Uint8slicetoboolslice(layer.IsEnable),
					})
				}
			}
			shelf_map := make(map[string][]response.LibshelfDto)
			for _, shelf := range shelfs {
				shelf_map[shelf.RowIdentity] = append(shelf_map[shelf.RowIdentity], response.LibshelfDto{
					Libshelf:   shelf,
					IsDeleted:  utils.Uint8slicetoboolslice(shelf.IsDeleted),
					IsEnable:   utils.Uint8slicetoboolslice(shelf.IsEnable),
					IsBosseyed: utils.Uint8slicetoboolslice(shelf.IsBosseyed),
					Layers:     layer_map[shelf.ID],
				})
			}
			for _, row := range rows {
				resp.Rows = append(resp.Rows, response.LibrowDto{
					Librow:    row,
					IsDeleted: utils.Uint8slicetoboolslice(row.IsDeleted),
					Shelfs:    shelf_map[row.ID],
				})
			}
		}
	}
	return resp, err
}
