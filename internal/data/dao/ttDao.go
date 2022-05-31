package dao

import (
	"context"
	"time"
	"tt.com/tt/internal/data"
	"tt.com/tt/internal/data/model"
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
	var result model.TtModel
	d.dataSource.DefautltDB.First(&result, "id = ?", 5)
	return result
}

func (d *TtDao) GetTtRedis(ctx context.Context) string {
	d.dataSource.DefaultRedis.Set(ctx, "tt", "tt redis result", 1000 * time.Second)
	value, _ := d.dataSource.DefaultRedis.Get(ctx, "tt").Result()
	return value
}