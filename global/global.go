package global

import (
	"sync"

	"github.com/qiniu/qmgo"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type YarmConfig struct {
	System struct {
		Port int `yaml:"port"`
	} `yaml:"system"`

	DB struct {
		Type string `yaml:"type"`
		URL  string `yaml:"url"`
	} `yaml:"db"`

	Redis struct {
		URL string `yaml:"url"`
		DB  int    `yaml:"db"`
	} `yaml:"redis"`

	MongoDB struct {
		URL string `yaml:"url"`
	} `yaml:"mongodb"`
}

var (
	OWEN_DB     *gorm.DB
	OWEN_DBList map[string]*gorm.DB
	OWEN_REDIS  redis.UniversalClient
	OWEN_MONGO  *qmgo.QmgoClient
	OWEN_CONFIG YarmConfig
	OWEN_LOCK   sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	OWEN_LOCK.RLock()
	defer OWEN_LOCK.RUnlock()
	return OWEN_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	OWEN_LOCK.RLock()
	defer OWEN_LOCK.RUnlock()
	db, ok := OWEN_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
