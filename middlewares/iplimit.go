package middlewares

import (
	"GINOWEN/global"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
)

// 加载黑名单（Redis 版本）
func LoadBlacklist() map[string]time.Time {
	blacklist := make(map[string]time.Time)

	// 获取所有黑名单的键
	keys, err := global.OWEN_REDIS.Keys(ctx, global.OWEN_CONFIG.System.Blacklistpre+"blacklist:*").Result()
	if err != nil {
		fmt.Println("Error loading blacklist keys from Redis:", err)
		return blacklist
	}

	// 遍历键并获取每个 IP 的解锁时间
	for _, key := range keys {
		val, err := global.OWEN_REDIS.Get(ctx, key).Result()
		if err != nil {
			fmt.Printf("Error retrieving value for key %s: %v\n", key, err)
			continue
		}
		var unlockTime time.Time
		err = json.Unmarshal([]byte(val), &unlockTime)
		if err != nil {
			fmt.Printf("Error parsing value for key %s: %v\n", key, err)
			continue
		}
		ip := key[len(global.OWEN_CONFIG.System.Blacklistpre+"blacklist:"):] // 从键名中提取 IP 地址
		blacklist[ip] = unlockTime
	}

	return blacklist
}

// 移除黑名单（Redis 版本）
func RemoveFromBlacklist(ip string) {
	key := fmt.Sprintf(global.OWEN_CONFIG.System.Blacklistpre+"blacklist:%s", ip)

	// 从 Redis 删除指定的键
	err := global.OWEN_REDIS.Del(ctx, key).Err()
	if err != nil {
		fmt.Printf("Error removing blacklist IP %s from Redis: %v\n", ip, err)
	} else {
		fmt.Printf("Successfully removed IP %s from blacklist\n", ip)
	}
}

// 保存黑名单（Redis 版本）
func SaveToBlacklist(ip string, unlockTime time.Time) {
	key := fmt.Sprintf(global.OWEN_CONFIG.System.Blacklistpre+"blacklist:%s", ip)
	value, err := json.Marshal(unlockTime)
	if err != nil {
		fmt.Println("Error marshaling unlock time:", err)
		return
	}

	// 设置 Redis 键值对，使用解锁时间作为过期时间
	err = global.OWEN_REDIS.Set(ctx, key, value, time.Until(unlockTime)).Err()
	if err != nil {
		fmt.Printf("Error saving blacklist IP %s to Redis: %v\n", ip, err)
	}
}

// IP 黑名单中间件
func IPBlacklistMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := fmt.Sprintf(global.OWEN_CONFIG.System.Blacklistpre+"blacklist:%s", ip)

		// 检查 IP 是否在 Redis 中
		val, err := global.OWEN_REDIS.Get(ctx, key).Result()
		if err == redis.Nil {
			// IP 不在黑名单中
			c.Next()
			return
		} else if err != nil {
			fmt.Println("Error checking IP in Redis:", err)
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		// 解析解锁时间
		var unlockTime time.Time
		err = json.Unmarshal([]byte(val), &unlockTime)
		if err != nil {
			fmt.Printf("Error parsing unlock time for IP %s: %v\n", ip, err)
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		// 如果当前时间在解锁时间之前，阻止访问
		if time.Now().Before(unlockTime) {
			c.JSON(403, gin.H{"error": fmt.Sprintf("Your IP (%s) is blacklisted until %s", ip, unlockTime)})
			c.Abort()
			return
		}

		c.Next()
	}
}
