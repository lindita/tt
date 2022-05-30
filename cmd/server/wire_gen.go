// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package server

import (
	"tt.com/tt/api"
	"tt.com/tt/api/app"
	"tt.com/tt/api/app/v1"
	"tt.com/tt/internal/conf"
	"tt.com/tt/internal/data"
	"tt.com/tt/internal/data/daos"
	"tt.com/tt/internal/data/db"
	"tt.com/tt/internal/services"
)

// Injectors from wire.go:

func NewApi(config *conf.Config) (*api.Api, func(), error) {
	appH5 := &app.AppH5{}
	gormDB := db.NewDb(config)
	client, err := db.NewRedis(config)
	if err != nil {
		return nil, nil, err
	}
	dataSource := data.NewDataSource(config, gormDB, client)
	ttDao := daos.NewTtDao(dataSource)
	dao := &daos.Dao{
		Tt: ttDao,
	}
	ttService := services.NewTtService(dao)
	service := &services.Service{
		Tt: ttService,
	}
	ttController := v1.NewTtController(service)
	appV1 := &app.AppV1{
		Tt: ttController,
	}
	appApp := &app.App{
		H5: appH5,
		V1: appV1,
	}
	apiApi := &api.Api{
		App: appApp,
	}
	return apiApi, func() {
	}, nil
}
