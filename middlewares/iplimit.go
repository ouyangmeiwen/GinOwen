package middlewares

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	// 黑名单 IP 列表
	Blacklist = make(map[string]time.Time)
)

// 加载黑名单
func LoadBlacklist() {
	// 读取黑名单文件
	data, err := os.ReadFile("blacklist.json")
	if err != nil {
		// 如果文件不存在，则初始化黑名单为空
		if os.IsNotExist(err) {
			Blacklist = make(map[string]time.Time)
		} else {
			fmt.Println("Error reading blacklist file:", err)
		}
		return
	}
	// 解析黑名单数据
	err = json.Unmarshal(data, &Blacklist)
	if err != nil {
		fmt.Println("Error parsing blacklist file:", err)
	}
}

// 保存黑名单
func SaveBlacklist() {
	// 异步保存黑名单到文件
	go func() {
		data, err := json.Marshal(Blacklist)
		if err != nil {
			fmt.Println("Error marshaling blacklist:", err)
			return
		}
		err = os.WriteFile("blacklist.json", data, 0644)
		if err != nil {
			fmt.Println("Error saving blacklist file:", err)
		}
	}()
}

// IP 黑名单中间件
func IPBlacklistMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		// 检查请求的 IP 是否在黑名单中
		if unlockTime, exists := Blacklist[ip]; exists {
			fmt.Printf(fmt.Sprintf("%s,%s\n", "now:", time.Now().Format("2006-01-02 15:04:05")))
			fmt.Printf(fmt.Sprintf("%s,%s\n", "unlockTime:", unlockTime.Format("2006-01-02 15:04:05")))
			// 比较 Unix 时间戳（秒级）
			if time.Now().Before(unlockTime) {
				c.JSON(403, gin.H{"error": fmt.Sprintf("Your IP (%s) is blacklisted until %s", ip, unlockTime)})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
