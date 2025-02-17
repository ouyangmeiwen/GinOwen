package models

import (
	"time"

	"gorm.io/gorm"
)

// 审计日志模型
type OwenAuditLog struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    uint           `gorm:"index"`             // 用户ID
	Action    string         `gorm:"type:varchar(255)"` // 操作描述（限制长度255）
	Request   string         `gorm:"type:text"`         // 请求参数
	Response  string         `gorm:"type:text"`         // 响应内容
	Error     string         `gorm:"type:varchar(500)"` // 异常信息，限制500字符
	Status    int            // 响应状态码
	Duration  float64        // 请求耗时
	CreatedAt time.Time      `gorm:"index"` // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"` // 软删除时间
}

// TableName Libitem's table name
func (*OwenAuditLog) TableName() string {
	return "owen_auditlog"
}
