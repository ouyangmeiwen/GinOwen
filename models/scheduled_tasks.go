package models

import "time"

type ScheduledTask struct {
	ID              int       `gorm:"primaryKey;autoIncrement"`                                                          // id 字段是主键，自动自增
	TaskName        string    `gorm:"type:varchar(255);not null"`                                                        // 任务名称，非空，类型为 varchar(255)
	ScheduleTime    time.Time `gorm:"not null"`                                                                          // 计划执行时间，非空
	IntervalSeconds *int      `gorm:"default:null"`                                                                      // 执行间隔（秒），可为 NULL
	LastRunTime     time.Time `gorm:"default:null"`                                                                      // 上次执行时间，可以为空
	NextRunTime     time.Time `gorm:"not null"`                                                                          // 下次执行时间，非空
	Status          string    `gorm:"type:enum('pending', 'running', 'completed', 'failed');default:'pending';not null"` // 任务状态，默认值为 'pending'
}

func (*ScheduledTask) TableName() string {
	return "owen_scheduledtask"
}
