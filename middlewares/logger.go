package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 继续处理请求
		c.Next()

		// 记录日志
		log.Infof("Request: %s %s | Duration: %v", c.Request.Method, c.Request.URL, time.Since(start))
	}
}
