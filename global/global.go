package global

import (
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type YarmConfig struct {
	System struct {
		Port int `yaml:"port"`
	} `yaml:"system"`

	DB struct {
		Type            string `yaml:"type"`
		Mysql           string `yaml:"mysql"`
		Mssql           string `yaml:"mssql"`
		Oracle          string `yaml:"oracle"`
		Postgres        string `yaml:"postgres"`
		Sqlite          string `yaml:"sqlite"`
		MaxOpenConns    int    `yaml:"maxopenconns"`
		MaxIdleConns    int    `yaml:"maxidleconns"`
		ConnMaxLifetime int    `yaml:"connmaxlifetime"`
	} `yaml:"db"`

	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`

	MongoDB struct {
		URI      string `yaml:"uri"`
		Database string `yaml:"database"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mongodb"`
}

var (
	OWEN_DB     *gorm.DB
	OWEN_DBList map[string]*gorm.DB
	OWEN_REDIS  redis.UniversalClient
	OWEN_MONGO  *mongo.Client
	OWEN_CONFIG YarmConfig
	OWEN_LOCK   sync.RWMutex
	OWEN_LOG    *zap.Logger
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
