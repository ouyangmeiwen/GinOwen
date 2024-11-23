package main

import (
	"GINOWEN/global"
	"GINOWEN/middlewares"
	"GINOWEN/routers"
	"GINOWEN/serviceinit"
	"GINOWEN/utils"

	"github.com/gin-gonic/gin"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample API using Gin and Swagger.
// @termsOfService https://example.com/terms
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

	// 创建 Gin 引擎
	r := gin.New()
	// r.Use(cors.Default())
	r.Use(gin.Recovery())
	r.Use(middlewares.Cors())

	//middlewares.LoadBlacklist()
	r.Use(middlewares.IPBlacklistMiddleware())
	//r.Use(cors.Default())
	routers.InitSwag(r)      //生成swagger文档 Swag init
	routers.InitAllRouter(r) //注册所有路由
	routers.RunAsServer(r)   //启动服务
}
