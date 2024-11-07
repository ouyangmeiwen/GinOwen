package config

import (
	"GINOWEN/models"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
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

func LoadConfig() Config {
	configFile, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer configFile.Close()

	var config Config
	decoder := yaml.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}
	return config
}

func InitDB(config Config) *gorm.DB {
	var dbErr error
	switch config.DB.Type {
	case "mysql":
		DB, dbErr = gorm.Open(mysql.Open(config.DB.URL), &gorm.Config{})
	case "postgres":
		DB, dbErr = gorm.Open(postgres.Open(config.DB.URL), &gorm.Config{})
	case "sqlite":
		DB, dbErr = gorm.Open(sqlite.Open(config.DB.URL), &gorm.Config{})
	case "sqlserver":
		DB, dbErr = gorm.Open(sqlserver.Open(config.DB.URL), &gorm.Config{})
	default:
		log.Fatalf("Unsupported database type: %s", config.DB.Type)
	}

	if dbErr != nil {
		log.Fatalf("Failed to connect to the database: %v", dbErr)
	}

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}
	// 配置数据库连接池
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(100)                 // 最大连接数
	sqlDB.SetMaxIdleConns(10)                  // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // 连接的最大生命周期

	// 自动迁移数据库结构
	err = DB.AutoMigrate(&models.User{})
	err = DB.AutoMigrate(&models.Order{})

	if err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}
	return DB
}
