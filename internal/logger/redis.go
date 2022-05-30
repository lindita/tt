package logger

import (
	"context"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

type RedisLogger struct {
	log *log.Logger
}

func (l *RedisLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	fmt.Printf("%+v", v)
	_ = l.log.Output(2, fmt.Sprintf(format, v...))
}

func NewRedisLogger() *RedisLogger {
	writer := lumberjack.Logger{
		Filename:   "./logs/redis.log",     // 日志文档路径
		MaxSize:    128,                      // 每个日志文档保存的最大尺寸 单位：M
		MaxBackups: 30,                       // 日志文档最多保存多少个备份
		MaxAge:     7,                        // 文档最多保存多少天
		Compress:   true,                     // 是否压缩
	}
	newLogger := log.New(&writer, "redis: ", log.LstdFlags|log.Lshortfile)
	return &RedisLogger{
		log: newLogger,
	}
}