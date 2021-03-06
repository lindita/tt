// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package server

import (
	"tt.com/tt/api"
	"tt.com/tt/api/app"
	"tt.com/tt/api/app/v1"
	"tt.com/tt/api/h5"
	"tt.com/tt/internal/conf"
	"tt.com/tt/internal/data"
	"tt.com/tt/internal/data/dao"
	"tt.com/tt/internal/data/db"
	"tt.com/tt/internal/service"
)

// Injectors from wire.go:

func NewApi(config *conf.Config) (*api.Api, func(), error) {
	gormDB := db.NewDb(config)
	client, err := db.NewRedis(config)
	if err != nil {
		return nil, nil, err
	}
	dataSource := data.NewDataSource(config, gormDB, client)
	ttDao := dao.NewTtDao(dataSource)
	daoDao := &dao.Dao{
		Tt: ttDao,
	}
	ttService := service.NewTtService(daoDao)
	serviceService := &service.Service{
		Tt: ttService,
	}
	ttController := v1.NewTtController(serviceService)
	v1V1 := &v1.V1{
		Tt: ttController,
	}
	appApp := &app.App{
		V1: v1V1,
	}
	h5TtController := h5.NewTtController(serviceService)
	h5H5 := &h5.H5{
		Tt: h5TtController,
	}
	apiApi := &api.Api{
		App: appApp,
		H5:  h5H5,
	}
	return apiApi, func() {
	}, nil
}
