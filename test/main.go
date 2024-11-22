package test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateAutoMigration() {
	// 指定 model 文件夹路径
	modelDir := "./model"
	outputFile := "./auto_migrate.go"

	// 获取 model 文件夹下所有的 Go 文件
	goFiles, err := getGoFilesInDir(modelDir)
	if err != nil {
		log.Fatalf("Failed to list Go files: %v", err)
	}

	// 提取所有结构体类型名
	var structNames []string
	for _, file := range goFiles {
		names, err := getStructNamesFromFile(file)
		if err != nil {
			log.Printf("Failed to parse file %s: %v", file, err)
			continue
		}
		structNames = append(structNames, names...)
	}

	// 生成 AutoMigrate 文件
	err = generateAutoMigrateFile(outputFile, structNames)
	if err != nil {
		log.Fatalf("Failed to generate auto-migrate file: %v", err)
	}

	log.Printf("Auto-migrate file generated at %s", outputFile)
}

// 获取指定目录中的所有 Go 文件
func getGoFilesInDir(dir string) ([]string, error) {
	var goFiles []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			goFiles = append(goFiles, path)
		}
		return nil
	})
	return goFiles, err
}

// 从 Go 文件中提取所有结构体类型名
func getStructNamesFromFile(filePath string) ([]string, error) {
	var structNames []string

	// 打开文件并解析 AST
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, filePath, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	// 遍历 AST 节点
	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		// 查找类型声明
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			// 判断是否为结构体
			if _, ok := typeSpec.Type.(*ast.StructType); ok {
				structNames = append(structNames, typeSpec.Name.Name)
			}
		}
	}

	return structNames, nil
}

// 生成包含 AutoMigrate 代码的文件
func generateAutoMigrateFile(outputFile string, structNames []string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入文件头
	_, err = file.WriteString(`package main

import (
	"log"
	"gorm.io/gorm"
	"your_project_path/model"
)

func AutoMigrate(DB *gorm.DB) {
	var err error
`)
	if err != nil {
		return err
	}

	// 写入每个结构体的迁移代码
	for _, structName := range structNames {
		line := fmt.Sprintf("\terr = DB.AutoMigrate(&model.%s{})\n", structName)
		_, err = file.WriteString(line)
		if err != nil {
			return err
		}
	}

	// 写入文件尾
	_, err = file.WriteString(`
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
`)
	return err
}
