package middlewares

import (
	"GINOWEN/auth"
	"GINOWEN/global"
	"GINOWEN/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 权限控制中间件
func AuthMiddleware(requiredPermissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		// 解析JWT
		claims, err := auth.ParseJWT(strings.TrimPrefix(tokenStr, "Bearer "))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// 从数据库加载用户角色
		var role models.OwenRole
		if err := global.OWEN_DB.First(&role, claims.RoleID).Error; err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "role not found"})
			c.Abort()
			return
		}

		// 检查角色权限
		userPermissions := strings.Split(role.Permissions, ",")
		hasPermission := false
		for _, permission := range requiredPermissions {
			if contains(userPermissions, permission) {
				hasPermission = true
				break
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
		if v == item {
			return true
		}
	}
	return false
}
