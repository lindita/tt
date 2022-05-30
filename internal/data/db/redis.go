package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"test.com/tt/internal/conf"
	"test.com/tt/internal/logger"
	"time"
)

func NewRedis(config *conf.Config) (*redis.Client, error) {
	redis.SetLogger(logger.NewRedisLogger())
	policy := config.Redis.Policy
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Default.GetAddr(),
		Password: config.Redis.Default.Password,
		DB:           0,                   // use default DB
		PoolSize:     policy.PoolSize,
		DialTimeout:  time.Second * policy.DialTimeout,
		ReadTimeout:  time.Second * policy.ReadTimeout,
		WriteTimeout: time.Second * policy.WriteTimeout,
		MaxRetries:   3,
	})
	_, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}