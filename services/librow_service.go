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
	var rows []models.Librow
	row_query := global.OWEN_DB.Model(&models.Librow{})
	if input.RowNo != nil && *input.RowNo > 0 {
		row_query = row_query.Where("RowNo = ?", *input.RowNo)
	}
	if input.StructCode != nil && len(*input.StructCode) > 0 {
		row_query = row_query.Where("Code LIKE ?", *input.StructCode+"%")
	}
	row_query = row_query.Where("IsDeleted = ?", 0).Order("RowNo ASC")
	err = row_query.Find(&rows).Error
	if err == nil {
		var rowids []string
		for _, row := range rows {
			rowids = append(rowids, row.ID)
		}
		var shelfs []models.Libshelf
		err = global.OWEN_DB.Model(&models.Libshelf{}).Where("IsDeleted = ? and IsEnable= ?", 0, 1).Where("RowIdentity IN ?", rowids).Order("Code ASC").Find(&shelfs).Error
		var shelfids []string
		if err == nil {
			for _, shelf := range shelfs {
				shelfids = append(shelfids, shelf.ID)
			}
			var layers []models.Liblayer
			err = global.OWEN_DB.Model(&models.Liblayer{}).Where("IsDeleted = ? and IsEnable= ?", 0, 1).Where("ShelfId IN ?", shelfids).Order("LayerNo ASC").Find(&layers).Error
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
