package utils

import (
	"time"
)

func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// "yyyy-mm-dd hh:mm:ss" 使用该格式
func FormatLocalTime(timestr string) (ts time.Time, err error) {
	// 解析时间
	// 获取当前时区（本地时区）
	location, err := time.LoadLocation("Local")
	if err != nil {
		return time.Time{}, err
	}
	localtime, err := time.ParseInLocation("2006-01-02 15:04:05", timestr, location)
	if err != nil {
		return time.Time{}, err
	}
	return localtime, nil
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
