package libs

import (
	"github.com/go-redis/redis"
)

var redisClent *redis.Client

func InitRedisClient(addr string) {
	if redisClent == nil {
		redisClent = redis.NewClient(&redis.Options{
			Addr: addr,
		})
	}
}

func Set(key string, val string) error {
	return redisClent.Set(key, val, 0).Err()
}

func Get(key string) (string, error) {
	return redisClent.Get(key).Result()
}
