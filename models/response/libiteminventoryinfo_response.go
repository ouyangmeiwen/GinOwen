package response

import (
	"GINOWEN/models"
)

type GetInventoryDto struct {
	Items      []models.Libiteminventoryinfo
	TotalCount int64
}

type CreateInventoryDto struct {
}

type UpdateInventoryDto struct {
}

type DeleteInventoryDto struct {
}
