package middlewares

import (
	"GINOWEN/auth"
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/utils"

	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 权限控制中间件
func AuthMiddleware(requiredPermissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 打印所有请求头
		fmt.Println("Request Headers:", c.Request.Header)
		// 获取 Authorization 头
		tokenStr := c.GetHeader("Authorization")
		// if tokenStr == "" {
		// 	if global.OWEN_CONFIG.System.EnableDebug {
		// 		tokenStr = utils.GetToken(c)
		// 		if tokenStr == "" {
		// 			// 如果没有提供 Authorization 头，返回错误信息
		// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		// 			c.Abort()
		// 			return
		// 		}
		// 	}
		// }
		// 移除 Bearer 前缀
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
		if tokenStr == "" {
			// 如果 Bearer 后面没有 token 字符串，返回错误
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token after Bearer"})
			c.Abort()
			return
		}

		// 解析 JWT Token
		claims, err := auth.ParseJWT(tokenStr)
		if err != nil {
			// 如果解析失败，返回错误信息
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// 从数据库加载角色信息
		var role models.OwenRole

		if len(global.OWEN_CONFIG.Redis.Addr) > 0 {
			rolestr, err := global.OWEN_REDIS.Get(global.Ctx, tokenStr).Result()
			if err != nil {
				// 如果解析失败，返回错误信息
				c.JSON(http.StatusUnauthorized, gin.H{"error": "token expire"})
				c.Abort()
				return
			}
			err = utils.FromJSON(rolestr, &role)
			if err != nil {
				// 如果解析失败，返回错误信息
				c.JSON(http.StatusUnauthorized, gin.H{"error": "token valid"})
				c.Abort()
				return
			}
		} else {
			if err := global.OWEN_DB.First(&role, claims.RoleID).Error; err != nil {
				c.JSON(http.StatusForbidden, gin.H{"error": "role not found"})
				c.Abort()
				return
			}
		}
		// 检查用户是否具有所需权限
		userPermissions := strings.Split(role.Permissions, ",")
		hasPermission := false
		if len(requiredPermissions) <= 0 {
			hasPermission = true //默认不需要权限
		} else {
			for _, permission := range requiredPermissions {
				if contains(userPermissions, permission) {
					hasPermission = true
					break
				}
			}

		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
			c.Abort()
			return
		}

		// 将用户信息传递给后续处理
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}

// contains 检查权限是否在权限列表中
func contains(list []string, item string) bool {
	for _, v := range list {
		if v == "all" { //管理员权限
			return true
		}
		if v == item {
			return true
		}
	}
	return false
}
