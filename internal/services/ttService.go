package services

import (
	"context"
	"github.com/spf13/cast"
	"tt.com/tt/internal/data/daos"
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
	r := s.dao.Tt.GetTtData(ctx)
	return "tt"+cast.ToString(r.Id)
}