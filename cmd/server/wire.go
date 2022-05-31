// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package server

import (
	"github.com/google/wire"
	"tt.com/tt/api"
	"tt.com/tt/api/app"
	"tt.com/tt/api/app/v1"
	"tt.com/tt/api/h5"
	"tt.com/tt/internal/conf"
	"tt.com/tt/internal/data"
	"tt.com/tt/internal/data/dao"
	"tt.com/tt/internal/service"
)

func NewApi(config *conf.Config) (*api.Api, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		dao.ProviderSet,
		wire.Struct(new(dao.Dao), "*"),
		service.ProviderSet,
		wire.Struct(new(service.Service), "*"),
		//增加目录需要在这加上
		v1.ProviderSet,
		h5.ProviderSet,
		wire.Struct(new(v1.V1), "*"),
		wire.Struct(new(h5.H5), "*"),
		wire.Struct(new(app.App), "*"),
		wire.Struct(new(api.Api), "*"),
	))
}

//go:generate wire gen
