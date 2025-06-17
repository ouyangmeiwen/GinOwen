package utils

import (
	"GINOWEN/global"
	"time"

	"github.com/redis/go-redis/v9"
)

// GetCache 从 Redis 获取 JSON 并反序列化到 out 指向的结构体中。
// 如果 key 不存在，返回 found=false 且 err=nil。
// 如果 Redis 错误或 JSON 错误，则返回 err。
func GetCache[T any](key string, out *T) (found bool, err error) {
	if len(global.OWEN_CONFIG.Redis.Addr) == 0 {
		return false, nil // Redis 未启用
	}

	val, err := global.OWEN_REDIS.Get(global.Ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil // key 不存在
		}
		return false, err // 其他 Redis 错误
	}

	err = FromJSON(val, out)
	if err != nil {
		return false, err // JSON 解析错误
	}

	return true, nil
}

// SetCache 将任意结构体转成 JSON 后写入 Redis
func SetCache[T any](key string, value T, ttl time.Duration) error {
	if len(global.OWEN_CONFIG.Redis.Addr) == 0 {
		return nil
	}

	jsonStr, err := ToJSON(value)
	if err != nil {
		return err
	}

	return global.OWEN_REDIS.Set(global.Ctx, key, jsonStr, ttl).Err()
}
