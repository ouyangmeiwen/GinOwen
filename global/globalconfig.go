package global

// SystemConfig 系统配置
type SystemConfig struct {
	Port            int    `yaml:"port"`
	RouterPre       string `yaml:"routerpre"`
	TokenExpire     int    `yaml:"tokenexpire"`
	Token           string `yaml:"token"`
	EnableDebug     bool   `yaml:"enabledebug"`
	Blacklistpre    string `yaml:"blacklistpre"`
	EnableBlacklist bool   `yaml:"enableblacklist"`
	IPWhitelist     string `yaml:"ipwhitelist"`
	Swaggerui       bool   `yaml:"swaggerui"`
	EnableWebsocket bool   `yaml:"enablewebsocket"`
	TaskInterval    int    `yaml:"taskinterval"`
	TaskDB          string `yaml:"taskdb"`  //taskpre
	Taskpre         string `yaml:"taskpre"` //taskpre

	CircuitBreaker struct {
		MaxRequests         int `yaml:"maxrequests"`
		Second              int `yaml:"second"`
		AddBlackListMinutes int `yaml:"addblacklistminutes"`
	} `yaml:"circuitbreaker"`

	RateLimiter struct {
		RateLimit           float64 `yaml:"ratelimit"`
		Burst               int     `yaml:"burst"`
		AddBlackListMinutes int     `yaml:"addblacklistminutes"`
	} `yaml:"ratelimiter"`
}

// DBConfig 数据库配置
type DBConfig struct {
	Type             string `yaml:"type"`
	CanAutoMigration bool   `yaml:"canautomigration"`
	CanAutoSynData   bool   `yaml:"canautosyndata"`
	MySQL            string `yaml:"mysql"`
	MSSQL            string `yaml:"mssql"`
	Oracle           string `yaml:"oracle"`
	Postgres         string `yaml:"postgres"`
	SQLite           string `yaml:"sqlite"`
	MaxOpenConns     int    `yaml:"maxopenconns"`
	MaxIdleConns     int    `yaml:"maxidleconns"`
	ConnMaxLifetime  int    `yaml:"connmaxlifetime"`
}

// RedisConfig Redis 配置
type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// MongoDBConfig MongoDB 配置
type MongoDBConfig struct {
	URI      string `yaml:"uri"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// RabbitMQConfig RabbitMQ 配置
type RabbitMQConfig struct {
	URL          string `yaml:"url"`
	ExchangeName string `yaml:"exchangename"`
	ExchangeType string `yaml:"exchangetype"`
}

// RabbitMQConsumerConfig RabbitMQ Consumer 配置
type RabbitMQConsumerConfig struct {
	URL          string `yaml:"url"`
	ExchangeName string `yaml:"exchangename"`
	ExchangeType string `yaml:"exchangetype"`
	QueueName    string `yaml:"queuename"`
	RoutingKey   string `yaml:"routingkey"`
}

// YarmConfig 总配置
type YarmConfig struct {
	System           SystemConfig           `yaml:"system"`
	DB               map[string]DBConfig    `yaml:"db"` // 使用 map 来处理多个数据库配置（如 one, two）
	Redis            RedisConfig            `yaml:"redis"`
	MongoDB          MongoDBConfig          `yaml:"mongodb"`
	RabbitMQ         RabbitMQConfig         `yaml:"rabbitmq"`
	RabbitMQConsumer RabbitMQConsumerConfig `yaml:"rabbitmqconsumer"`
}
