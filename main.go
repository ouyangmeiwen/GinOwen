package main

import (
	"GINOWEN/config"
	"GINOWEN/global"
	"GINOWEN/middlewares"
	"GINOWEN/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	// 加载配置文件
	global.GVA_CONFIG = config.LoadConfig()
	global.GVA_DB = config.InitDB()
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
