package main

import (
	"GINOWEN/config"
	"GINOWEN/global"
	"GINOWEN/middlewares"
	"GINOWEN/routers"
	"GINOWEN/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化日志
	utils.InitLogger()
	defer utils.Sync()

	global.OWEN_LOG.Debug("开始程序！")

	// 加载配置文件
	global.OWEN_CONFIG = config.LoadConfig()
	global.OWEN_DB = config.InitDB()
	config.InitRedis() //初始化redis

	config.AutoMigrateDB() //数据库自动迁移

	// 创建 Gin 引擎
	r := gin.New()
	// r.Use(cors.Default())
	r.Use(gin.Recovery())
	r.Use(middlewares.Cors())

	routers.InitSwag(r)      //生成swagger文档 Swag init
	routers.InitApi(r)       //API控制
	routers.InitAllRouter(r) //注册所有路由
	routers.RunAsService(r)
}
