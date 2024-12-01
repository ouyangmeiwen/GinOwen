package routers

import (
	"GINOWEN/middlewares"

	"github.com/gin-gonic/gin"
)

type ScheduledTaskRouter struct {
}

func RegisterTaskRoutes(r *gin.Engine) {
	api := r.Group("Task").Use(middlewares.AuthMiddleware()).Use(middlewares.AuditLogMiddleware())
	{
		api.GET("AddTask", ApiGroup.taskApi.AddTask)
	}
}
