package utils

import (
	"time"
)

func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当前时间的字符串表示
func GetCurrentTimeString(format string) string {
	return time.Now().Format(format)
}

// 计算两个时间之间的差异
func TimeDifference(t1, t2 time.Time) time.Duration {
	return t1.Sub(t2)
}

// 将字符串转换为时间
func ParseTime(timeStr, format string) (time.Time, error) {
	return time.Parse(format, timeStr)
}

// 示例格式: "2006-01-02 15:04:05"
func FormatTime(t time.Time, format string) string {
	return t.Format(format)
}
