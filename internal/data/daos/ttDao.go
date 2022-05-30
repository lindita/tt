package daos

import (
	"context"
	"fmt"
	"test.com/tt/internal/data"
	"test.com/tt/internal/data/model"
	"test.com/tt/internal/logger"
	"time"
)

type TtDao struct {
	dataSource *data.DataSource
}

func NewTtDao(dataSource *data.DataSource) *TtDao {
	return &TtDao{
		dataSource: dataSource,
	}
}

func (d *TtDao) GetTtData(ctx context.Context) model.TtModel {
	d.dataSource.DefaultRedis.Set(context.TODO(), "aaa", "o12222222", 10000 * time.Second).Val()
	value, _ := d.dataSource.DefaultRedis.Get(context.TODO(), "aaa").Result()
	fmt.Println(value)
	logger.GetLogger().Named("xxx").Error("xxxxxffffffffffff")
	var xx model.TtModel
	d.dataSource.DefautltDB.First(&xx, "id = ?", 5)
	fmt.Printf("%+v", xx)
	return xx
}