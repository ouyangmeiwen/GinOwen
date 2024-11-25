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
		Port            int    `yaml:"port"`
		RouterPre       string `yaml:"routerpre"`
		TokenExpire     int    `yaml:"tokenexpire"`     //
		Token           string `yaml:"token"`           //
		EnableDebug     bool   `yaml:"enabledebug"`     //
		Blacklistpre    string `yaml:"blacklistpre"`    //
		EnableBlacklist bool   `yaml:"enableblacklist"` //
		IPWhitelist     string `yaml:"ipwhitelist"`     //
		Swaggerui       bool   `yaml:"swaggerui"`       // 是否开启 Swagger UI
		EnableWebsocket bool   `yaml:"enablewebsocket"` // 是否开启 websocket
		CircuitBreaker  struct {
			MaxRequests         int `yaml:"maxrequests"`         // 最大请求数
			Second              int `yaml:"second"`              // 监控时间窗口（秒）
			AddBlackListMinutes int `yaml:"addblacklistminutes"` // 阈值后等待时间
		} `yaml:"circuitbreaker"`
		RateLimiter struct {
			RateLimit           float64 `yaml:"ratelimit"`           // 每秒请求数量
			Burst               int     `yaml:"burst"`               // 最大并发数
			AddBlackListMinutes int     `yaml:"addblacklistminutes"` // 阈值后等待时间
		} `yaml:"ratelimiter"`
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

	RabbitMQ struct {
		URL          string `yaml:"url"`
		ExchangeName string `yaml:"exchangename"`
		ExchangeType string `yaml:"exchangetype"`
	} `mapstructure:"rabbitmq"`

	RabbitMQConsumer struct {
		URL          string `yaml:"url"`
		ExchangeName string `yaml:"exchangename"`
		ExchangeType string `yaml:"exchangetype"`
		QueueName    string `yaml:"queuename"`
		RoutingKey   string `yaml:"routingkey"`
	} `mapstructure:"rabbitmqconsumer"`
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
