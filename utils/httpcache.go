package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFileContent(method string) (string, error) {
	// 获取程序当前工作目录（运行时所在目录）
	rootDir, err := os.Getwd()
	if err != nil {
		return "Error getting current working directory", err
	}

	// 构建完整文件路径
	filePath := filepath.Join(rootDir, "httpcache", fmt.Sprintf("%s.json", method))

	// 读取文件
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "Error reading file: " + err.Error(), err
	}

	// 转换为字符串并打印
	content := string(data)
	return content, nil
}
