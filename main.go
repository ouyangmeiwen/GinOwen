package main

import (
	"GINOWEN/config"
	"GINOWEN/initcontrollers"

	"GINOWEN/middlewares"
	"log"

	"GINOWEN/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	db := initcontrollers.InitDB()

	// 初始化用户和订单控制器
	userController := initcontrollers.InitUserController(db)
	orderController := initcontrollers.InitOrderController(db)
	// 创建 Gin 引擎
	// 加载配置文件
	config.LoadConfig()
	// 创建 Gin 引擎
	r := gin.New()
	// 添加中间件
	api := r.Group("/api")
	api.Use(middlewares.RateLimiter())    // 异常恢复
	api.Use(middlewares.RateLimiter())    // 添加限流中间件
	api.Use(middlewares.CircuitBreaker()) // 添加熔断中间件

	// 注册路由
	routes.RegisterUserRoutes(r, userController)
	routes.RegisterOrderRoutes(r, orderController)

	// 启动服务
	if err := r.Run(":7899"); err != nil {
		log.Fatal(err)
	}

	// 在应用结束时关闭数据库连接
	defer func() {
		// 确保关闭数据库连接
		if db != nil {
			sqlDB, err := db.DB()
			if err == nil {
				sqlDB.Close()
			}
		}
	}()
}
