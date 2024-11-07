package main

import (
	"GINOWEN/config"
	"GINOWEN/initcontrollers"
	"fmt"

	"GINOWEN/middlewares"
	"log"

	"GINOWEN/routes"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	// 启用 CORS
	r.Use(cors.Default())

	// 添加中间件
	api := r.Group("/api")
	api.Use(middlewares.RateLimiter())    // 异常恢复
	api.Use(middlewares.RateLimiter())    // 添加限流中间件
	api.Use(middlewares.CircuitBreaker()) // 添加熔断中间件

	// 注册路由
	routes.RegisterUserRoutes(r, userController)
	routes.RegisterOrderRoutes(r, orderController)

	// 注册Swagger路由
	r.GET("/swagger/*any", func(c *gin.Context) {
		fmt.Println("Swagger route accessed") // 这是调试日志
		ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
	})

	// 启动服务
	if err := r.Run(":7899"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Swagger UI available at http://localhost:7899/swagger/index.html")

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
