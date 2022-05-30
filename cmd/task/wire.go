// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package task

import (
	"github.com/google/wire"
	"tt.com/tt/internal/conf"
	"tt.com/tt/internal/data"
	"tt.com/tt/internal/data/daos"
	"tt.com/tt/internal/services"
)

func newService(config *conf.Config) (*services.Service, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		daos.ProviderSet,
		wire.Struct(new(daos.Dao), "*"),
		services.ProviderSet,
		wire.Struct(new(services.Service), "*"),
	))
}

//go:generate wire gen
