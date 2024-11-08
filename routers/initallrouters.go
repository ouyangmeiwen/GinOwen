package routers

import (
	"GINOWEN/docs"
	"GINOWEN/global"
	"GINOWEN/middlewares"
	"GINOWEN/utils"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitAllRouter(r *gin.Engine) {
	// 添加中间件
	api := r.Group("/api")          //API开头的增加熔断
	api.Use(middlewares.Recovery()) // 异常恢复

	// 注册路由
	RegisterUserRoutes(r)
	RegisterOrderRoutes(r)
	RegisterLibitemRoutes(r)
	RegisterUploadfileRoutes(r)

}
func InitSwag(r *gin.Engine) {
	// 如果 swagger.json 存放在 docs 目录下，确保提供静态文件服务
	r.Static("/swagger", "./docs")
	// 注册Swagger路由  要放到注册路由的后面
	r.GET(docs.SwaggerInfo.BasePath+"/swagger-ui/*any", func(c *gin.Context) {
		ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/swagger.json"))(c)
	})
}
func RunAsService(r *gin.Engine) {
	fmt.Println("========================================================")
	port := global.OWEN_CONFIG.System.Port
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
		if global.OWEN_DB != nil {
			sqlDB, err := global.OWEN_DB.DB()
			if err == nil {
				sqlDB.Close()
			}
		}
	}()
}
