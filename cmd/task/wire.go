// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package task

import (
	"github.com/google/wire"
	"tt.com/tt/internal/conf"
	"tt.com/tt/internal/data"
	"tt.com/tt/internal/data/daos"
	"tt.com/tt/internal/service"
)

func newService(config *conf.Config) (*service.Service, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		daos.ProviderSet,
		wire.Struct(new(daos.Dao), "*"),
		service.ProviderSet,
		wire.Struct(new(service.Service), "*"),
	))
}

//go:generate wire gen
