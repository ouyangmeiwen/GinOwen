package request

import "time"

type AddScheduledTaskInput struct {
	TaskName        string    `json:"TaskName" form:"TaskName"`         // 任务名称，非空，类型为 varchar(255)
	ScheduleTime    time.Time `json:"ScheduleTime" form:"ScheduleTime"` // 计划执行时间，非空 2024-12-01T17:10:45+08:00
	IntervalSeconds *int      `json:"IntervalSeconds" form:"IntervalSeconds" default:"0"`
}
