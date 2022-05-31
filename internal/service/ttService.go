package service

import (
	"context"
	"tt.com/tt/internal/data/daos"
	"tt.com/tt/internal/data/model"
	"tt.com/tt/internal/exception"
	"tt.com/tt/internal/logger"
)

type TtService struct {
	dao *daos.Dao
}

func NewTtService(dao *daos.Dao) *TtService {
	return  &TtService{
		dao: dao,
	}
}

func (s *TtService) GetTt(ctx context.Context) string  {
	logger.GetLogger().Named("tt").Info("tt info custom log!")
	return "tt"
}

func (s *TtService) TestRedis(ctx context.Context) string  {
	r := s.dao.Tt.GetTtRedis(ctx)
	return"tt redis:" + r
}

func (s *TtService) TestMysql(ctx context.Context) model.TtModel  {
	r := s.dao.Tt.GetTtData(ctx)
	return r
}

func (s *TtService) TestPanic(ctx context.Context) string  {
	panic(exception.NewApiException(1, "error msg"))
	panic("like system panic")
	return "tt test panic"
}