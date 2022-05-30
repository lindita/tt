// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package server

import (
	"github.com/google/wire"
	"tt.com/tt/api"
	"tt.com/tt/api/app"
	v1 "tt.com/tt/api/app/v1"
	"tt.com/tt/internal/conf"
	"tt.com/tt/internal/data"
	"tt.com/tt/internal/data/daos"
	"tt.com/tt/internal/services"
)

func NewApi(config *conf.Config) (*api.Api, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		daos.ProviderSet,
		wire.Struct(new(daos.Dao), "*"),
		services.ProviderSet,
		wire.Struct(new(services.Service), "*"),
		v1.ProviderSet,
		//增加目录需要在这加上
		wire.Struct(new(app.AppV1), "*"),
		wire.Struct(new(app.AppH5), "*"),
		wire.Struct(new(app.App), "*"),
		wire.Struct(new(api.Api), "*"),
	))
}

//go:generate wire gen
