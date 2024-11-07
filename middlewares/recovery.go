package middlewares

import (
	"GINOWEN/utils"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				utils.ResponseError(c, 500, "Internal Server Error")
			}
		}()
		c.Next()
	}
}
