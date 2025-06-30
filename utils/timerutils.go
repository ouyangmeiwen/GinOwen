package utils

import (
	"time"
)

// 东八区时间
func LoadLocation() (*time.Location, error) {
	return time.LoadLocation("Asia/Shanghai")
}

// 将字符串解析出东八区时间
// "yyyy-mm-dd hh:mm:ss" 使用该格式
func ParseInLocation(layout string, timestr string) (ts time.Time, err error) {
	// 解析时间
	// 获取当前时区（本地时区）
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return time.Time{}, err
	}
	localtime, err := time.ParseInLocation(layout, timestr, location)
	if err != nil {
		return time.Time{}, err
	}
	return localtime, nil
}

// 将时间类型转成东八区时间字符串
// "yyyy-mm-dd hh:mm:ss" 使用该格式
func FormatInLocation(layout string, t time.Time) string {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		// 如果加载失败就退回原时间
		return t.Format(layout)
	}
	return t.In(loc).Format(layout)
}

// 东八区时间系统当前时间
func NowInLocation() time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(loc)
	return now
}

// 东八区时间系统当前时间
func NowDateLocation() time.Time {
	var now = NowInLocation()
	// 取当前日期部分（清除时间）
	dateOnly := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return dateOnly
}

func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当前时间的字符串表示 "2006-01-02 15:04:05"
// fmt.Println(GetCurrentTimeString("2006-01-02 15:04:05"))
// fmt.Println(GetCurrentTimeString("2006/01/02"))
// fmt.Println(GetCurrentTimeString("15:04:05"))
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

// 获取当前本地时间
func Now() time.Time {
	return time.Now()
}

// 获取当前UTC时间
func NowUTC() time.Time {
	return time.Now().UTC()
}

// 将时间格式化为指定格式
func Format(t time.Time, layout string) string {
	return t.Format(layout)
}

// 解析字符串为时间
func Parse(layout, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

// 加指定时间
func Add(t time.Time, duration time.Duration) time.Time {
	return t.Add(duration)
}

// 时间加天、月、年
func AddDate(t time.Time, years, months, days int) time.Time {
	return t.AddDate(years, months, days)
}

// 获取当前时间戳（秒）
func Unix(t time.Time) int64 {
	return t.Unix()
}

// 获取当前时间戳（毫秒）
func UnixMilli(t time.Time) int64 {
	return t.UnixMilli()
}

// 获取当前时间戳（纳秒）
func UnixNano(t time.Time) int64 {
	return t.UnixNano()
}

// 比较两个时间
func IsBefore(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

func IsAfter(t1, t2 time.Time) bool {
	return t1.After(t2)
}

func IsEqual(t1, t2 time.Time) bool {
	return t1.Equal(t2)
}
