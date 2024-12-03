package serviceinit

import (
	"log"
)

type AutoService struct{}

// 定义接口
type Migratable_interface interface {
	CusAutoMigrate()
}
type SysData_interface interface {
	CusSyncDatabase() error
}

// 判断并调用
func callCusAutoMigrateIfExists(service interface{}) {
	if m, ok := service.(Migratable_interface); ok {
		m.CusAutoMigrate() // 调用方法
	} else {
		log.Fatalf("CusAutoMigrate method does not exist")
	}
}

// 判断并调用
func callCusSyncDatabaseIfExists(service interface{}) {
	if m, ok := service.(SysData_interface); ok {
		m.CusSyncDatabase() // 调用方法
	} else {
		log.Fatalf("CusSyncDatabase method does not exist")
	}
}
