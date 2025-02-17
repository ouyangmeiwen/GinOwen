package routers

import (
	"GINOWEN/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterGPIRoutes(r *gin.Engine) {
	api := r.Group("GPI").Use(middlewares.AuthMiddleware("gpi"))
	{
		api.POST("GPIReceive", ApiGroup.gpiApi.GPIReceive)
	}
}
