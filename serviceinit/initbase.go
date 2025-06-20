package serviceinit

import (
	"GINOWEN/extend/rabbitmqextend"
	"GINOWEN/global"
	"GINOWEN/models"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	//"github.com/dzwvip/oracle"
	"github.com/glebarez/sqlite" // 使用 godror 驱动
	// 注册 godror 驱动
	"GINOWEN/global/cfg"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func LoadConfig() cfg.YarmConfig {

	configfile := "config.yaml"
	configFile, err := os.Open(configfile)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer configFile.Close()

	var config cfg.YarmConfig
	decoder := yaml.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}
	return config
}

func InitDB() {
	//config := global.OWEN_CONFIG
	dbMap := global.OWEN_CONFIG.DB
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
	for key, dbCfg := range dbMap {
		if len(dbCfg.Type) <= 0 {
			continue
		}
		var DB *gorm.DB
		var sqlDB *sql.DB
		switch dbCfg.Type {
		case "mysql":
			DB, dbErr = gorm.Open(mysql.Open(dbCfg.MySQL), &gorm.Config{Logger: newLogger})
		case "postgres":
			DB, dbErr = gorm.Open(postgres.Open(dbCfg.Postgres), &gorm.Config{Logger: newLogger})
		case "sqlite":
			DB, dbErr = gorm.Open(sqlite.Open(dbCfg.SQLite), &gorm.Config{Logger: newLogger})
		case "mssql":
			DB, dbErr = gorm.Open(sqlserver.Open(dbCfg.MSSQL), &gorm.Config{Logger: newLogger})
		case "oracle":
			// sqlDB, _ = sql.Open("godror", dbCfg.Oracle)
			// // DB, dbErr = gorm.Open(gorm.Config{
			// // 	ConnPool: oracleDB,  // 将底层的 sql.DB 连接池传给 GORM
			// // 	Logger:   newLogger, // 可选：设置 GORM 日志输出级别
			// // })
			// var currentTime string
			// err := sqlDB.QueryRow("SELECT TO_CHAR(SYSDATE, 'YYYY-MM-DD HH24:MI:SS') AS current_time FROM dual").Scan(&currentTime)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// // 打印当前时间
			// fmt.Println("验证oracle数据库:", currentTime)

		default:
			log.Fatalf("Unsupported database type: %s", dbCfg.Type)
		}

		if dbErr != nil {
			log.Fatalf("Failed to connect to the database: %v", dbErr)
		}

		// 配置数据库连接池
		var err error
		if sqlDB == nil {
			sqlDB, err = DB.DB()
			if err != nil {
				log.Fatalf("Failed to get database instance: %v", err)
			}
		}
		sqlDB.SetMaxOpenConns(dbCfg.MaxOpenConns)                                      // 最大连接数
		sqlDB.SetMaxIdleConns(dbCfg.MaxIdleConns)                                      // 最大空闲连接数
		sqlDB.SetConnMaxLifetime((time.Duration(dbCfg.ConnMaxLifetime)) * time.Minute) // 连接的最大生命周期
		if key == "default" {
			global.OWEN_DB = DB
			global.SQL_DB = sqlDB
		}
		if global.OWEN_DBList == nil {
			global.OWEN_DBList = make(map[string]*gorm.DB)
		}
		global.OWEN_DBList[key] = DB
		log.Printf("初始化数据库%s", dbCfg.Type)
	}
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

func AutoMigrateDB(DB *gorm.DB) {
	// 自动迁移数据库结构
	var err error
	err = DB.AutoMigrate(&models.OwenUser{})
	err = DB.AutoMigrate(&models.OwenRole{})
	err = DB.AutoMigrate(&models.OwenAuditLog{})
	err = DB.AutoMigrate(&models.ScheduledTask{})
	//是否需要初始化model
	if _, ok := global.OWEN_DBList["to"]; ok {
		s1 := AutoService{}
		if global.OWEN_CONFIG.DB["to"].CanAutoMigration {
			callCusAutoMigrateIfExists(s1)
			log.Printf("生成to数据库表结构")
		}
		if global.OWEN_CONFIG.DB["to"].CanAutoSynData {
			callCusSyncDatabaseIfExists(s1)
			log.Printf("生成to数据库表结构数据")
		}
	}
	if err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}
	addDefaultData(DB)
}
func addDefaultData(DB *gorm.DB) {

	// 检查并添加默认角色数据
	if err := DB.First(&models.OwenRole{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		defaultRoles := []models.OwenRole{
			{Name: "Admin", Permissions: "all"},
			{Name: "User", Permissions: "sysauditlmslog"},
		}
		for _, role := range defaultRoles {
			if err := DB.Create(&role).Error; err != nil {
				log.Printf("Failed to insert default role %s: %v", role.Name, err)
			}
		}
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	// 检查并添加默认用户数据
	if err := DB.First(&models.OwenUser{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		defaultUsers := []models.OwenUser{
			{Username: "admin", Password: string(hashedPassword), RoleID: 1}, // 示例用户
			{Username: "user", Password: string(hashedPassword), RoleID: 2},
		}
		for _, user := range defaultUsers {
			if err := DB.Create(&user).Error; err != nil {
				log.Printf("Failed to insert default user %s: %v", user.Username, err)
			}
		}
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

func InitRabbiMQ() {

	if len(global.OWEN_CONFIG.RabbitMQConsumer.URL) > 0 {
		rabbitmqextend.RegisterMQConsumer(global.OWEN_CONFIG.RabbitMQConsumer.URL,
			global.OWEN_CONFIG.RabbitMQConsumer.ExchangeName,
			global.OWEN_CONFIG.RabbitMQConsumer.ExchangeType,
			global.OWEN_CONFIG.RabbitMQConsumer.QueueName,
			global.OWEN_CONFIG.RabbitMQConsumer.RoutingKey)
	}

	if len(global.OWEN_CONFIG.RabbitMQ.URL) > 0 {
		rabbitmqextend.RegisterMQPublisher(global.OWEN_CONFIG.RabbitMQ.URL,
			global.OWEN_CONFIG.RabbitMQ.ExchangeName,
			global.OWEN_CONFIG.RabbitMQ.ExchangeType)

		// 发布消息
		err := rabbitmqextend.InstancePublisher.PublishMessage("default_routingKey.init", []byte("service.is.ok"))
		if err != nil {
			log.Fatalf("Failed to publish message: %v", err)
		}
		//test
		// 发送文本消息
		textMsg := rabbitmqextend.Data{
			DataType: "TextMessage",
			Body:     json.RawMessage(`{"Content": "Hello, RabbitMQ!"}`),
		}
		err = rabbitmqextend.InstancePublisher.PublishData("default_routingKey.init", textMsg)
		if err != nil {
			log.Fatalf("Error publishing message: %v", err)
		}

		// 发送图像消息
		imageMsg := rabbitmqextend.Data{
			DataType: "ImageMessage",
			Body:     json.RawMessage(`{"ImageURL": "http://example.com/image.jpg", "AltText": "A sample image"}`),
		}
		err = rabbitmqextend.InstancePublisher.PublishData("default_routingKey.init", imageMsg)
		if err != nil {
			log.Fatalf("Error publishing message: %v", err)
		}

	}

}
