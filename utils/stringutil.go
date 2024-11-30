package utils

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

func GenerateGUID() string {
	newUUID := uuid.New()
	return strings.ReplaceAll(newUUID.String(), "-", "")
}

func RemoveBrackets(str string) string {
	// 去除开头的 '['
	str = strings.TrimPrefix(str, "[")
	// 去除结尾的 ']'
	str = strings.TrimSuffix(str, "]")
	return str
}

// TrimSpaces 去除字符串前后空格
func TrimSpaces(s string) string {
	return strings.TrimSpace(s)
}

// substring 返回字符串 str 从 start 到 end 的子串
func SafeSubstring(str string, start int, length int) string {
	runes := []rune(str)
	if length > len(runes) {
		length = len(runes)
	}
	truncated := string(runes[start : start+length])
	return truncated
}

// 限制只能是数字和字母
func NumberString(input string) string {
	// 定义正则表达式，只匹配字母和数字
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	// 使用正则表达式替换非字母和数字的字符
	return re.ReplaceAllString(input, "")
}

// 判断字符串是否为空
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// 将字符串转换为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 将字符串转换为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// 反转字符串
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 移除字符串中的所有空格
func RemoveWhitespace(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

// 替换字符串中的子字符串
func ReplaceSubstring(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// 检查字符串是否以指定前缀开头
func StartsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// 检查字符串是否以指定后缀结尾
func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// 按照指定的分隔符分割字符串
func SplitString(s, delimiter string) []string {
	return strings.Split(s, delimiter)
}

// 使用指定的分隔符将字符串列表连接成一个字符串
func JoinStrings(strs []string, delimiter string) string {
	return strings.Join(strs, delimiter)
}

// 将字符串中的每个单词首字母大写
func CapitalizeWords(s string) string {
	return strings.Title(s)
}

// 检查字符串中是否包含指定子字符串
func ContainsSubstring(s, substring string) bool {
	return strings.Contains(s, substring)
}

// 将字符串中的所有字母转换为首字母大写
func Capitalize(s string) string {
	runes := []rune(s)
	if len(runes) == 0 {
		return s
	}
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

func StringListContain(input string, keyToCheck string) bool {
	// 将字符串以逗号分隔
	parts := strings.Split(input, ",")
	// 初始化哈希表
	hashMap := make(map[string]string)
	for _, it := range parts {
		hashMap[it] = it
	}
	if value, exists := hashMap[keyToCheck]; exists {
		fmt.Printf("Key '%s' 存在，值为: %s\n", keyToCheck, value)
		return true
	} else {
		fmt.Printf("Key '%s' 不存在\n", keyToCheck)
		return false
	}
}
func HashContain(hashMap map[string]string, keyToCheck string) bool {
	if value, exists := hashMap[keyToCheck]; exists {
		fmt.Printf("Key '%s' 存在，值为: %s\n", keyToCheck, value)
		return true
	} else {
		fmt.Printf("Key '%s' 不存在\n", keyToCheck)
		return false
	}
}
