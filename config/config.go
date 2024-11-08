package config

import (
	"GINOWEN/global"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func LoadConfig() global.YarmConfig {
	configFile, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer configFile.Close()

	var config global.YarmConfig
	decoder := yaml.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}
	return config
}

func InitDB() *gorm.DB {
	config := global.OWEN_CONFIG
	var dbErr error

	// 配置 Gorm 日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io.Writer 输出日志
		logger.Config{
			SlowThreshold: time.Second, // 慢查询 SQL 阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 彩色打印
		},
	)

	switch config.DB.Type {
	case "mysql":
		DB, dbErr = gorm.Open(mysql.Open(config.DB.Mysql), &gorm.Config{Logger: newLogger})
	case "postgres":
		DB, dbErr = gorm.Open(postgres.Open(config.DB.Postgres), &gorm.Config{Logger: newLogger})
	case "sqlite":
		DB, dbErr = gorm.Open(sqlite.Open(config.DB.Sqlite), &gorm.Config{Logger: newLogger})
	case "sqlserver":
		DB, dbErr = gorm.Open(sqlserver.Open(config.DB.Mssql), &gorm.Config{Logger: newLogger})
	case "oracle":
		// 使用 godror 驱动连接到 Oracle 数据库
		DB, dbErr = gorm.Open(mysql.Open(config.DB.Mysql), &gorm.Config{Logger: newLogger})
	default:
		log.Fatalf("Unsupported database type: %s", config.DB.Type)
	}

	if dbErr != nil {
		log.Fatalf("Failed to connect to the database: %v", dbErr)
	}

	// 配置数据库连接池
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(config.DB.MaxOpenConns)                                      // 最大连接数
	sqlDB.SetMaxIdleConns(config.DB.MaxIdleConns)                                      // 最大空闲连接数
	sqlDB.SetConnMaxLifetime((time.Duration(config.DB.ConnMaxLifetime)) * time.Minute) // 连接的最大生命周期

	return DB
}
func InitRedis() {
	redisCfg := global.OWEN_CONFIG.Redis
	if len(redisCfg.Addr) <= 0 {
		return
	}
	var client redis.UniversalClient
	client = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	} else {
		global.OWEN_REDIS = client
	}
}

func AutoMigrateDB() {
	// 自动迁移数据库结构
	var err error
	// err = DB.AutoMigrate(&models.TestUser{})
	// err = DB.AutoMigrate(&models.TestOrder{})
	if err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}
}

// InitMongoDB 连接 MongoDB
func InitMongoDB() (*mongo.Client, error) {
	config := global.OWEN_CONFIG
	if len(config.MongoDB.URI) <= 0 {
		return nil, errors.New("入参不能空")
	}
	clientOptions := options.Client().ApplyURI(config.MongoDB.URI)
	if config.MongoDB.Username != "" && config.MongoDB.Password != "" {
		clientOptions.SetAuth(options.Credential{
			Username: config.MongoDB.Username,
			Password: config.MongoDB.Password,
		})
	}

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// 检查连接
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	global.OWEN_MONGO = client
	return client, nil
}
