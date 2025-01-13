package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

// 判断文件是否存在
func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}

// 判断目录是否存在
func DirectoryExists(directory string) bool {
	info, err := os.Stat(directory)
	return !os.IsNotExist(err) && info.IsDir()
}

// 创建新文件
func CreateFile(filepath string) (*os.File, error) {
	return os.Create(filepath)
}

// 创建目录（包括父目录）
func CreateDirectory(directory string) error {
	return os.MkdirAll(directory, os.ModePerm)
}

// 删除文件
func DeleteFile(filepath string) error {
	return os.Remove(filepath)
}

// 删除目录及其内容
func DeleteDirectory(directory string) error {
	return os.RemoveAll(directory)
}

// 清空目录内容
func ClearDirectory(directory string) error {
	files, err := os.ReadDir(directory)
	if err != nil {
		return err
	}
	for _, file := range files {
		err := os.RemoveAll(filepath.Join(directory, file.Name()))
		if err != nil {
			return err
		}
	}
	return nil
}

// 移动文件或目录
func Move(src, dst string) error {
	return os.Rename(src, dst)
}

// 复制文件
func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

// 复制目录
func CopyDirectory(srcDir, dstDir string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(dstDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		}
		return CopyFile(path, targetPath)
	})
}

// 获取文件大小
func GetFileSize(filepath string) (int64, error) {
	info, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// 读取文件内容
func ReadFile(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// 写入文件（覆盖）
func WriteFile(filepath string, content string) error {
	return os.WriteFile(filepath, []byte(content), os.ModePerm)
}

// 写入文件（追加）
func AppendToFile(filepath string, content string) error {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return err
}

// 递归列出目录下的所有文件
func ListFilesRecursive(directory string) ([]string, error) {
	var files []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// 获取目录大小
func GetDirectorySize(directory string) (int64, error) {
	var size int64
	err := filepath.Walk(directory, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}

// 获取文件扩展名
func GetFileExtension(filename string) string {
	return strings.ToLower(filepath.Ext(filename))
}

// 计算文件 MD5 哈希
func GetFileMD5(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// 判断是否是符号链接
func IsSymlink(filepath string) (bool, error) {
	info, err := os.Lstat(filepath)
	if err != nil {
		return false, err
	}
	return info.Mode()&os.ModeSymlink != 0, nil
}

// 验证文件类型
// 检查文件类型
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

// 判断切片中是否包含某字符串
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
