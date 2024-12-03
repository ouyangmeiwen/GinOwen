package serviceinit

import "fmt"

type AutoService struct{}

// 定义接口
type Migratable interface {
	CusAutoMigrate()
}
type SysData interface {
	CusSyncDatabase() error
}

// 判断并调用
func callCusAutoMigrateIfExists(service interface{}) {
	if m, ok := service.(Migratable); ok {
		m.CusAutoMigrate() // 调用方法
	} else {
		fmt.Println("CusAutoMigrate method does not exist")
	}
}

// 判断并调用
func callCusSyncDatabaseIfExists(service interface{}) {
	if m, ok := service.(SysData); ok {
		m.CusSyncDatabase() // 调用方法
	} else {
		fmt.Println("CusAutoMigrate method does not exist")
	}
}
