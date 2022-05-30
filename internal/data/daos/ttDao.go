package daos

import (
	"context"
	"fmt"
	"time"
	"tt.com/tt/internal/data"
	"tt.com/tt/internal/data/model"
	"tt.com/tt/internal/logger"
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
	d.dataSource.DefaultRedis.Set(ctx, "tt", "testtt", 10000 * time.Second).Val()
	value, _ := d.dataSource.DefaultRedis.Get(ctx, "testtt").Result()
	fmt.Println(value)
	logger.GetLogger().Named("xxx").Error("error")
	var xx model.TtModel
	d.dataSource.DefautltDB.First(&xx, "id = ?", 5)
	fmt.Printf("%+v", xx)
	return xx
}