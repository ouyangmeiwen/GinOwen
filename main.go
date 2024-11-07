package main

import (
	"GINOWEN/config"
	"GINOWEN/global"
	"GINOWEN/middlewares"
	"GINOWEN/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// 加载配置文件
	cfg := config.LoadConfig()
	global.GVA_DB = config.InitDB(cfg)
	config.AutoMigrateDB() //数据库自动迁移

	// 创建 Gin 引擎
	r := gin.New()
	// r.Use(cors.Default())
	r.Use(gin.Recovery())
	r.Use(middlewares.Cors())

	routes.InitSwag(r)
	routes.InitApi(r)       //API控制
	routes.InitAllRouter(r) //注册所有路由
	routes.RunAsService(r, cfg)
}
