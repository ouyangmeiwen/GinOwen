// initrouter/db.go
package initcontrollers

import (
	"GINOWEN/config"

	"gorm.io/gorm"
)

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	// 根据配置来选择不同的数据库驱动，例如这里使用的是 SQLite
	appConfig := config.LoadConfig()
	db := config.InitDB(appConfig)
	return db
}
