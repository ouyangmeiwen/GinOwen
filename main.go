package main

import (
	"GINOWEN/config"
	"GINOWEN/docs"
	"GINOWEN/initcontrollers"
	"GINOWEN/middlewares"
	"GINOWEN/utils"
	"fmt"
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
	r.Use(gin.Recovery())
	r.Use(middlewares.Cors())
	// 如果 swagger.json 存放在 docs 目录下，确保提供静态文件服务
	r.Static("/swagger", "./docs")

	// 注册Swagger路由  要放到注册路由的后面
	r.GET(docs.SwaggerInfo.BasePath+"/swagger-ui/*any", func(c *gin.Context) {
		ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/swagger.json"))(c)
	})

	// 添加中间件
	api := r.Group("/api")

	api.Use(middlewares.Recovery())       // 异常恢复
	api.Use(middlewares.RateLimiter())    // 添加限流中间件
	api.Use(middlewares.CircuitBreaker()) // 添加熔断中间件

	// 注册路由
	routes.RegisterUserRoutes(r, userController)
	routes.RegisterOrderRoutes(r, orderController)
	fmt.Println("========================================================")
	port := 7899
	url := "http://localhost:" + fmt.Sprint(port) + "/swagger-ui/index.html"
	fmt.Println(url)
	fmt.Println("========================================================")
	utils.OpenBrowser(url)

	// 启动服务
	if err := r.Run(":" + fmt.Sprint(port)); err != nil {
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
