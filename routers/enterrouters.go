package routers

import (
	"GINOWEN/docs"
	"GINOWEN/global"
	"GINOWEN/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitAllRouter(r *gin.Engine) {
	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
	// 注册路由
	RegisterLibitemRoutes(r)
	RegisterUploadfileRoutes(r)
	RegisterSysauditlmslogRoutes(r)
	RegisterJWTRoutes(r)
}
func InitSwag(r *gin.Engine) {
	// 如果 swagger.json 存放在 docs 目录下，确保提供静态文件服务
	r.Static("/swagger", "./docs")
	// 注册Swagger路由  要放到注册路由的后面
	r.GET(docs.SwaggerInfo.BasePath+"/swagger-ui/*any", func(c *gin.Context) {
		ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/swagger.json"))(c)
	})
}

// initServer 函数创建并返回一个具有精细化配置的 http.Server 实例
func initServer(address string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second, // 读取超时
		WriteTimeout:   20 * time.Second, // 写入超时
		IdleTimeout:    10 * time.Second, // 空闲连接超时
		MaxHeaderBytes: 1 << 20,          // 最大请求头大小 (1 MB)
	}
}

func RunAsServer(r *gin.Engine) {
	port := global.OWEN_CONFIG.System.Port
	url := fmt.Sprintf("http://localhost:%d/swagger-ui/index.html", port)

	log.Printf("Swagger URL: %s\n", url)
	utils.OpenBrowser(url)

	// 初始化服务器
	srv := initServer(fmt.Sprintf(":%d", port), r)

	// 优雅关闭信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 在协程中启动服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server listen error: %v\n", err)
		}
	}()
	log.Printf("Server started on port %d\n", port)

	// 监听信号并优雅关闭
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// 确保数据库连接关闭
	if global.OWEN_DB != nil {
		sqlDB, err := global.OWEN_DB.DB()
		if err == nil {
			sqlDB.Close()
			log.Println("Database connection closed.")
		}
	}

	log.Println("Server exiting.")
}
