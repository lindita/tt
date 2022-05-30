// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package server

import (
	"github.com/google/wire"
	appV1 "tt.com/tt/api/app/v1"
	"tt.com/tt/internal/conf"
	"tt.com/tt/internal/data"
	"tt.com/tt/internal/data/daos"
	"tt.com/tt/internal/services"
)

func newController(config *conf.Config) (*appV1.Controller, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		daos.ProviderSet,
		wire.Struct(new(daos.Dao), "*"),
		services.ProviderSet,
		wire.Struct(new(services.Service), "*"),
		appV1.ProviderSet,
		wire.Struct(new(appV1.Controller), "*"),
	))
}

//go:generate wire gen
