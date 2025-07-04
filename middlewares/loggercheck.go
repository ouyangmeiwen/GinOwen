package middlewares

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/utils"
	"bytes"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 字符串截取函数，避免溢出
func truncateString(s string, maxLength int) string {
	return utils.SafeSubstring(s, 0, maxLength)
}

// 自定义 ResponseWriter，用于捕获响应内容
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // 将响应写入 body 缓冲区
	return w.ResponseWriter.Write(b)
}

// 审计日志中间件
func AuditLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 捕获请求参数
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody)) // 重置 Body 以供后续读取
		}

		// 捕获响应内容
		writer := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = writer
		// 处理请求
		c.Next()
		userIDint := utils.GetCurrentUserTd(c)
		if userIDint < 0 {
			// c.JSON(http.StatusInternalServerError, "invalid userID type")
			c.Abort()
			return
		}
		// 记录响应状态、内容和执行时间
		statusCode := c.Writer.Status()
		responseBody := writer.body.String()
		duration := time.Since(start).Seconds()

		// 捕获异常信息
		var errMsg string
		if len(c.Errors) > 0 {
			errMsg = c.Errors.String()
		}

		// 创建审计日志记录
		auditLog := models.OwenAuditLog{
			UserID:    uint(userIDint), // 假设已经通过 JWT 或 Session 获取用户 ID
			Action:    truncateString(c.Request.RequestURI, 255),
			Request:   truncateString(string(requestBody), 1000),
			Response:  truncateString(responseBody, 1000),
			Error:     truncateString(errMsg, 500),
			Status:    statusCode,
			Duration:  duration,
			CreatedAt: time.Now(),
		}

		// 保存日志到数据库
		if err := global.OWEN_DB.Create(&auditLog).Error; err != nil {
			log.Printf("failed to save audit log: %v", err)
		}
	}
}

// 定期清理审计日志
func StartAuditLogCleanup(db *gorm.DB) {
	ticker := time.NewTicker(7 * 24 * time.Hour) // 每周清理一次
	go func() {
		for {
			<-ticker.C
			cleanupAuditLogs(db)
		}
	}()
}

func cleanupAuditLogs(db *gorm.DB) {
	expirationDate := time.Now().AddDate(0, 0, -30) // 30天前的日期
	db.Where("created_at < ?", expirationDate).Unscoped().Delete(&models.OwenAuditLog{})
}
