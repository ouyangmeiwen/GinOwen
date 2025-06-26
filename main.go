package main

import (
	"GINOWEN/api"
	"GINOWEN/extend"
	"GINOWEN/extend/rabbitmqextend"
	"GINOWEN/extend/websocketextend"
	"GINOWEN/global"
	"GINOWEN/middlewares"
	"GINOWEN/routers"
	"GINOWEN/serviceinit"
	"GINOWEN/services"
	"GINOWEN/utils"
	"fmt"
	"log"
	"sync"
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

	now := time.Now()
	utcstr := now.Format("2006-01-02 15:04:05")
	fmt.Printf("UTC GinOwen API started at %s\n", utcstr)
	parsedTime, err := time.Parse("2006-01-02 15:04:05", utcstr)
	if err != nil {
		fmt.Println("UTC GinOwen API started at (parse error):", err)
	} else {
		fmt.Println("UTC GinOwen API started at", parsedTime)
	}

	locstr := utils.FormatInLocation("2006-01-02 15:04:05", now)
	fmt.Printf("Asia/Shanghai GinOwen API started at %s\n", locstr)
	ts, err := utils.ParseInLocation("2006-01-02 15:04:05", locstr)
	if err != nil {
		fmt.Println("Asia/Shanghai GinOwen API started at (parse error):", err)
	} else {
		fmt.Println("Asia/Shanghai GinOwen API started at:", ts)
	}

	// 加载配置文件
	global.OWEN_CONFIG = serviceinit.LoadConfig()
	// 初始化日志
	utils.InitLogger()
	defer utils.Sync()
	serviceinit.InitDB()

	extend.Test()
	if _, ok := global.OWEN_DBList["from"]; ok {
		if global.OWEN_CONFIG.DB["from"].CanCreateModel {
			extend.CreateDBModles(global.OWEN_DBList["from"]) //生成数据库结构根据数据库配置
		}
		if global.OWEN_CONFIG.DB["from"].CanAutoMigration {
			extend.CreateAutoMigrationFile()

		}
		if global.OWEN_CONFIG.DB["from"].CanAutoSynData {
			extend.CreateAutoSyncFile()
		}
	} //占位置
	if global.OWEN_DB != nil {
		middlewares.StartAuditLogCleanup(global.OWEN_DB) // 启动日志清理任务
	}

	serviceinit.InitRedis() //初始化redis
	if global.OWEN_DB != nil {
		serviceinit.AutoMigrateDB(global.OWEN_DB) //数据库自动迁移
	}
	serviceinit.InitMongoDB() //mongodb

	serviceinit.InitRabbiMQ() //rabbitmq
	defer func() {
		if rabbitmqextend.InstancePublisher != nil {
			rabbitmqextend.InstancePublisher.Close()
		}
		if rabbitmqextend.InstanceConsumer != nil {
			rabbitmqextend.InstanceConsumer.Close()
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
	if global.OWEN_CONFIG.System.EnableWebsocket {
		websocketextend.RegisterWebsocket(r) //注册websocket
	}
	routers.InitAllRouter(r)
	if global.OWEN_CONFIG.System.TaskInterval > 0 {
		go startScheduledTaskService(&api.ServicesGroup.TaskService, global.OWEN_CONFIG.System.TaskInterval) //任务调度
	} //注册所有路由
	routers.RunAsServer(r) //启动服务
}

var mu_task sync.Mutex // 用于保护共享资源的锁
// 定时任务服务
func startScheduledTaskService(taskService *services.ScheduledTaskService, interval int) {
	// 定期检查计划任务并执行
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	// 使用 for range 代替 select{}
	for range ticker.C {
		mu_task.Lock() // 加锁，确保对共享资源的访问是安全的
		// 处理任务
		err := taskService.ProcessTasks()
		if err != nil {
			log.Printf("Error processing tasks: %v", err)
		} else {
			log.Println("Tasks processed successfully")
		}
		mu_task.Unlock() // 解锁
	}
}
