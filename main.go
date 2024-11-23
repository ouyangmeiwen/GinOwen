package main

import (
	"GINOWEN/global"
	"GINOWEN/middlewares"
	"GINOWEN/rabbitmq"
	"GINOWEN/routers"
	"GINOWEN/serviceinit"
	"GINOWEN/utils"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// @title GinOwen API
// @version 1.0
// @description 这是一个全局启用了安全验证的 API。
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT Authorization header (Bearer token)
func main() {

	//test.CreateAutoMigration()

	// 加载配置文件
	global.OWEN_CONFIG = serviceinit.LoadConfig()
	// 初始化日志
	utils.InitLogger()
	defer utils.Sync()

	//global.OWEN_LOG.Debug("开始程序！")

	global.OWEN_DB = serviceinit.InitDB()
	middlewares.StartAuditLogCleanup(global.OWEN_DB) // 启动日志清理任务

	serviceinit.InitRedis() //初始化redis

	serviceinit.AutoMigrateDB() //数据库自动迁移

	serviceinit.InitMongoDB() //mongodb

	serviceinit.InitRabbiMQ() //rabbitmq
	defer func() {
		if rabbitmq.Instance != nil {
			// 使用 Close 方法来关闭连接和通道
			if err := rabbitmq.Instance.Close(); err != nil {
				log.Fatalf("Error closing RabbitMQ: %v", err)
			}
		}
	}()

	// 创建 Gin 引擎
	r := gin.New()
	// r.Use(cors.Default())
	r.Use(gin.Recovery())
	r.Use(middlewares.Cors())

	//middlewares.LoadBlacklist()
	if global.OWEN_CONFIG.System.EnableBlacklist {
		r.Use(middlewares.IPBlacklistMiddleware())
	}
	if global.OWEN_CONFIG.System.CircuitBreaker.MaxRequests > 0 && global.OWEN_CONFIG.System.CircuitBreaker.Second > 0 {
		r.Use(middlewares.CircuitBreaker(uint32(global.OWEN_CONFIG.System.CircuitBreaker.MaxRequests), time.Duration(global.OWEN_CONFIG.System.CircuitBreaker.Second))) //10 秒最多20个请求 主要用于检测和隔离异常服务，防止因为后端服务的故障导致请求堆积或服务雪崩。
	}
	if global.OWEN_CONFIG.System.RateLimiter.Burst > 0 && global.OWEN_CONFIG.System.RateLimiter.RateLimit > 0 {
		r.Use(middlewares.RateLimiter(global.OWEN_CONFIG.System.RateLimiter.RateLimit, global.OWEN_CONFIG.System.RateLimiter.Burst)) //10个请求 20个并发 主要用于限制请求速率，防止服务因为流量过大而超载。
	}
	//r.Use(cors.Default())
	if global.OWEN_CONFIG.System.Swaggerui {
		routers.InitSwag(r) //生成swagger文档 Swag init
	}
	routers.InitAllRouter(r) //注册所有路由
	routers.RunAsServer(r)   //启动服务
}
