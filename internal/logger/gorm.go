package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func NewGormLog() logger.Interface {
	writer := lumberjack.Logger{
		Filename:   "./logs/mysql.log",     // 日志文档路径
		MaxSize:    128,                      // 每个日志文档保存的最大尺寸 单位：M
		MaxBackups: 30,                       // 日志文档最多保存多少个备份
		MaxAge:     7,                        // 文档最多保存多少天
		Compress:   true,                     // 是否压缩
	}
	newLogger := logger.New(
		log.New(&writer, "gorm: ", log.LstdFlags|log.Lshortfile), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,   // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:      false,         // 禁用彩色打印
		},
	)
	return newLogger
}