package utils

import (
	"GINOWEN/global"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetFileContent(method string) (string, error) {
	fmt.Println("==============================测试接口数据" + method)
	if !strings.Contains(global.OWEN_CONFIG.Redis.Addr, "192.168.229.130") {
		return "", fmt.Errorf("不允许读取本地测试数据")
	}
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
