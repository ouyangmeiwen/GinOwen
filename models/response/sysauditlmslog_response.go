package response

import (
	"GINOWEN/models"
	"GINOWEN/utils"
)

type QueryLmsItem struct {
	models.Sysauditlmslog
}

type QueryLmsDto struct {
	TotalCount int64          //总数量
	Items      []QueryLmsItem //数据集
}

type CreateLmsLogDto struct {
	utils.RespBaseResult
}
type UpdateLmsLogDto struct {
	utils.RespBaseResult
}

type DeleteLmsLogDto struct {
	utils.RespBaseResult
}
