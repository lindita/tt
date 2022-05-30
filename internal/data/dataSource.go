package data

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"test.com/tt/internal/conf"
)

type DataSource struct {
	Config       *conf.Config
	DefautltDB   *gorm.DB
	DefaultRedis *redis.Client
}

func NewDataSource(config *conf.Config, db *gorm.DB, redis *redis.Client) *DataSource {
	data := &DataSource{Config: config, DefautltDB: db, DefaultRedis: redis}
	return data
}

type contextTxKey struct{}

// ExecTx gorm Transaction 开启事务
func (d *DataSource) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.DefautltDB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

// TransactionDB 根据此方法来判断当前的 db 是不是使用 事务的 DB
func (d *DataSource) DBTx(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.DefautltDB
}