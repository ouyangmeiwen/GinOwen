package gpiutils

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
func ReadFile(filepath string) (string, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteFile(filepath string, data string) error {
	return ioutil.WriteFile(filepath, []byte(data), os.ModePerm)
}

// 检查文件类型
func IsValidFileType(file *multipart.FileHeader) bool {
	// allowedTypes := []string{"image/jpeg", "image/png", "application/pdf"} // 允许的文件类型
	// ext := strings.ToLower(filepath.Ext(file.Filename))
	// allowedExtensions := []string{".jpg", ".jpeg", ".png", ".pdf"} // 允许的文件扩展名

	allowedTypes := []string{
		"application/msword", // .doc
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document", // .docx
		"application/vnd.ms-excel", // .xls
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", // .xlsx
		"application/pdf", // .pdf
		"text/plain",      // .txt
		"application/zip", // .zip
	} // 允许的文件类型

	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExtensions := []string{
		".doc",
		".docx",
		".xls",
		".xlsx",
		".pdf",
		".txt",
		".zip",
	} // 允许的文件扩展名

	// 检查扩展名
	if !Contains(allowedExtensions, ext) {
		return false
	}

	// 检查 MIME 类型
	fileType := file.Header.Get("Content-Type")
	return Contains(allowedTypes, fileType)
}

// contains 检查字符串是否在切片中
func Contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
