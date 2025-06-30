package convert

import (
	"fmt"
	"strconv"
)

// --- 字符串转数字 ---

func Atoi(s string) (int, error) {
	return strconv.Atoi(s)
}

func ParseInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func ParseUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func ParseFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// --- 数字转字符串 ---

func Itoa(i int) string {
	return strconv.Itoa(i)
}

func FormatInt64(i int64) string {
	return strconv.FormatInt(i, 10)
}

func FormatUint64(u uint64) string {
	return strconv.FormatUint(u, 10)
}

func FormatFloat64(f float64, prec int) string {
	return strconv.FormatFloat(f, 'f', prec, 64)
}

// --- 示例辅助：必须成功的封装（不推荐业务使用） ---

func MustAtoi(s string) int {
	i, err := Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("invalid int string: %s", s))
	}
	return i
}

func MustFloat64(s string) float64 {
	f, err := ParseFloat64(s)
	if err != nil {
		panic(fmt.Sprintf("invalid float string: %s", s))
	}
	return f
}
