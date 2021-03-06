package service

import (
	"context"
	"tt.com/tt/internal/data/dao"
	"tt.com/tt/internal/data/model"
	"tt.com/tt/internal/exception"
	"tt.com/tt/internal/logger"
)

type TtService struct {
	dao *dao.Dao
}

func NewTtService(dao *dao.Dao) *TtService {
	return  &TtService{
		dao: dao,
	}
}

func (s *TtService) GetTt(ctx context.Context) string  {
	logger.GetLogger().Named("tt").Info("tt info custom log!")
	return "success"
}

func (s *TtService) TestRedis(ctx context.Context) string  {
	r := s.dao.Tt.GetTtRedis(ctx)
	return r
}

func (s *TtService) TestMysql(ctx context.Context) model.TtModel  {
	r := s.dao.Tt.GetTtData(ctx)
	return r
}

func (s *TtService) TestPanic(ctx context.Context) {
	panic(exception.NewException(1, "error msg"))
	//panic("like system panic")
}