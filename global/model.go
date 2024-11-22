package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_Entity struct {
	ID        uint      `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt time.Time `json:"CreatedAt"`            // 创建时间
	CreatorID uint      `json:"CreatorID"`            // 创建者ID需要手动赋值
}

type GVA_AuditedEntity struct {
	ID        uint      `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt time.Time `json:"CreatedAt"`            // 创建时间
	CreatorID uint      `json:"CreatorID"`            // 创建者ID需要手动赋值
	UpdatedAt time.Time `json:"UpdatedAt"`            // 更新时间
	UpdateID  uint      `json:"UpdateID"`             // 修改者ID需要手动赋值
}

type GVA_FullAuditedEntity struct {
	ID        uint           `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt time.Time      `json:"CreatedAt"`            // 创建时间
	CreatorID uint           `json:"CreatorID"`            // 创建者ID需要手动赋值
	UpdatedAt time.Time      `json:"UpdatedAt"`            // 更新时间
	UpdateID  uint           `json:"UpdateID"`             // 修改者ID需要手动赋值
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       // 删除时间
	DeleterID uint           `json:"DeleterID"`            // 删除者ID需要手动赋值
}
