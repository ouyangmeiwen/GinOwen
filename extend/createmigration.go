package extend

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gorm.io/gen"
	"gorm.io/gorm"
)

func CreateAutoMigrationFile() {
	// 指定 model 文件夹路径
	modelDir := "./extenddb/model"
	outputFile := "./serviceinit/auto_migrate.go"

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
			log.Fatalf("Failed to parse file %s: %v", file, err)
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
	log.Printf("生成from数据库表结构文件")
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
	_, err = file.WriteString(`package serviceinit

import (
	"GINOWEN/extenddb/model"
	"GINOWEN/global"
	"log"
)

func(s AutoService) CusAutoMigrate() {
	DB := global.OWEN_DBList["to"]
	var err error
`)
	if err != nil {
		return err
	}

	// 写入每个结构体的迁移代码
	for _, structName := range structNames {
		line := fmt.Sprintf("\terr = DB.AutoMigrate(&model.%s{})\n", structName)
		line += fmt.Sprintf(`
	if err != nil {
		log.Fatalf("Failed to migrate table: %s")
		return
	}
`, structName)
		_, err = file.WriteString(line)
		if err != nil {
			return err
		}
	}

	// 写入文件尾
	_, err = file.WriteString(`
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return
	}
}
`)
	return err
}

func CreateAutoSyncFile() {
	// 指定 model 文件夹路径
	modelDir := "./extenddb/model"
	outputFile := "./serviceinit/auto_sync.go"

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
			log.Fatalf("Failed to parse file %s: %v", file, err)
			continue
		}
		structNames = append(structNames, names...)
	}

	// 生成 SyncDatabase 文件
	err = generateAutoSyncFile(outputFile, structNames)
	if err != nil {
		log.Fatalf("Failed to generate auto-sync file: %v", err)
	}
	log.Printf("Auto-sync file generated at %s", outputFile)
	log.Printf("生成from数据库表数据文件")
}

// 生成包含 SyncDatabase 代码的文件
func generateAutoSyncFile(outputFile string, structNames []string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入文件头
	_, err = file.WriteString(`package serviceinit

import (
	"log"
	"GINOWEN/extenddb/model"
	"GINOWEN/global"
)

func(s AutoService) CusSyncDatabase() error {
	batchSize := 500 // 每批次同步的数据量
	var offset int // 用于分页

`)

	if err != nil {
		return err
	}

	// 写入每个结构体的同步代码
	for _, structName := range structNames {
		// 生成每个结构体同步的分页代码
		line := fmt.Sprintf(`
		// Syncing %s model data
		offset = 0
		for {
			var %sData []model.%s
			// 分页查询数据，不依赖排序字段
			if err := global.OWEN_DBList["from"].Offset(offset).Limit(batchSize).Find(&%sData).Error; err != nil {
				log.Fatalf("Failed to query table: %s")
				return err
			}

			// 如果没有更多数据，退出循环
			if len(%sData) == 0 {
				break
			}

			// 更新 offset
			offset += batchSize

			// 将数据插入到 db2
			if err := global.OWEN_DBList["to"].CreateInBatches(%sData, batchSize).Error; err != nil {
				log.Fatalf("Failed to insert table: %s")
				return err
			}
		}
		`, structName, structName, structName, structName, structName, structName, structName, structName)
		_, err = file.WriteString(line)
		if err != nil {
			return err
		}
	}

	// 写入文件尾
	_, err = file.WriteString(`
	return nil
}
`)

	return err
}

// 遍历目录并重命名以_开头的文件
func renameFilesInDir(dir string) error {
	// 遍历目录
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理文件，不处理子目录
		if !info.IsDir() && strings.HasPrefix(info.Name(), "_") {
			// 新文件名去掉前缀_
			newName := strings.TrimLeft(info.Name(), "_")
			newPath := filepath.Join(filepath.Dir(path), newName)

			// 重命名文件
			err := os.Rename(path, newPath)
			if err != nil {
				return fmt.Errorf("无法重命名文件 %s: %v", path, err)
			}
			fmt.Printf("文件 %s 已重命名为 %s\n", path, newPath)
		}

		return nil
	})

	return err
}

func CreateDBModles(db *gorm.DB) {
	// 生成实例
	outpath := "extenddb/dbgen"
	g := gen.NewGenerator(gen.Config{
		// 相对执行`go run`时的路径, 会自动创建目录
		OutPath: outpath,

		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,

		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true, // generate pointer when field is nullable

		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false, // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values

		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false, // detect integer field's unsigned type, adjust generated data type
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false, // generate with gorm index tag
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true, // generate with gorm column type tag
	})
	// 设置目标 db
	g.UseDB(db)

	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf
	dataMap := map[string]func(detailType string) (dataType string){
		"tinyint":   func(detailType string) (dataType string) { return "int64" },
		"smallint":  func(detailType string) (dataType string) { return "int64" },
		"mediumint": func(detailType string) (dataType string) { return "int64" },
		"bigint":    func(detailType string) (dataType string) { return "int64" },
		"int":       func(detailType string) (dataType string) { return "int64" },
		//"varchar":   func(detailType string) (dataType string) { return "sql.NullString" },
		//"text":      func(detailType string) (dataType string) { return "sql.NullString" },
	}

	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	// 自定义模型结体字段的标签
	// 将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		toStringField := `balance, `
		if strings.Contains(toStringField, columnName) {
			return columnName + ",string"
		}
		return columnName
	})

	// 将非默认字段名的字段定义为自动时间戳和软删除字段;
	// 自动时间戳默认字段名为:`updated_at`、`created_at, 表字段数据类型为: INT 或 DATETIME
	// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME
	autoUpdateTimeField := gen.FieldGORMTag("update_time", "column:update_time;type:int unsigned;autoUpdateTime")
	autoCreateTimeField := gen.FieldGORMTag("create_time", "column:create_time;type:int unsigned;autoCreateTime")
	softDeleteField := gen.FieldType("delete_time", "soft_delete.DeletedAt")
	// 模型自定义选项组
	fieldOpts := []gen.ModelOpt{jsonField, autoCreateTimeField, autoUpdateTimeField, softDeleteField}

	// 创建模型的结构体,生成文件在 model 目录; 先创建的结果会被后面创建的覆盖
	// 这里创建个别模型仅仅是为了拿到`*generate.QueryStructMeta`类型对象用于后面的模型关联操作中

	//Address := g.GenerateModel("address")

	// 创建全部模型文件, 并覆盖前面创建的同名模型
	allModel := g.GenerateAllTable(fieldOpts...)

	// 创建有关联关系的模型文件
	// User := g.GenerateModel("user",
	// 	append(
	// 		fieldOpts,
	// 		// user 一对多 address 关联, 外键`uid`在 address 表中
	// 		gen.FieldRelate(field.HasMany, "Address", Address, &field.RelateConfig{GORMTag: "foreignKey:UID"}),
	// 	)...,
	// )
	// Address = g.GenerateModel("address",
	// 	append(
	// 		fieldOpts,
	// 		gen.FieldRelate(field.BelongsTo, "User", User, &field.RelateConfig{GORMTag: "foreignKey:UID"}),
	// 	)...,
	// )

	// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖
	//g.ApplyBasic(User, Address)
	g.ApplyBasic(allModel...)
	g.Execute()
	log.Printf("生成from数据库表结构")
	renameFilesInDir("extenddb")
}

func Test() {

}
