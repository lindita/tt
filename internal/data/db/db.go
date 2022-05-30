package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"test.com/tt/internal/conf"
	"test.com/tt/internal/logger"
)

func NewDb(config *conf.Config) *gorm.DB {
	dataSourceTemplate := "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&timeout=30s&writeTimeout=10s&readTimeout=10s"
	masterDSNConfig := config.Database.Default.Master
	masterDSN := fmt.Sprintf(
		dataSourceTemplate,
		masterDSNConfig.Username,
		masterDSNConfig.Password,
		masterDSNConfig.Host,
		masterDSNConfig.Port,
		masterDSNConfig.Dbname,
		masterDSNConfig.Charset,
	)

	db, err := gorm.Open(mysql.Open(masterDSN), &gorm.Config{
		Logger: logger.NewGormLog(),
	})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(masterDSNConfig.ConnMin)
	sqlDB.SetMaxOpenConns(masterDSNConfig.ConnMax)
	if config.Runtime == "dev" {
		return db.Debug()
	}
	return db
}
