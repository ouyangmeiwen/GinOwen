package main

import (
	"GINOWEN/config"
	"GINOWEN/global"
	"GINOWEN/middlewares"
	"GINOWEN/routers"
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
	// 加载配置文件
	global.OWEN_CONFIG = config.LoadConfig()
	// 初始化日志
	utils.InitLogger()
	defer utils.Sync()

	//global.OWEN_LOG.Debug("开始程序！")

	global.OWEN_DB = config.InitDB()
	middlewares.StartAuditLogCleanup(global.OWEN_DB) // 启动日志清理任务

	config.InitRedis() //初始化redis

	config.AutoMigrateDB() //数据库自动迁移

	config.InitMongoDB() //mongodb

	// 创建 Gin 引擎
	r := gin.New()
	// r.Use(cors.Default())
	r.Use(gin.Recovery())
	r.Use(middlewares.Cors())
	//r.Use(cors.Default())
	routers.InitSwag(r)      //生成swagger文档 Swag init
	routers.InitAllRouter(r) //注册所有路由
	routers.RunAsServer(r)   //启动服务
}
