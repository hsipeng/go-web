package db

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8" // 注意导入的是新版本
	"lirawx.cn/go-web/config"
)

var (
	// Rdb redis db
	Rdb *redis.Client
)

// InitRedisClient 初始化连接
func InitRedisClient(cfg config.RedisConfig) (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Password: cfg.Password, // no password set
		DB:       0,            // use default DB
		PoolSize: 100,          // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = Rdb.Ping(ctx).Result()
	return err
}
