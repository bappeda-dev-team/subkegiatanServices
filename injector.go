//go:build wireinject
// +build wireinject

package main

import (
	"subkegiatanServices/app"

	"subkegiatanServices/controller"
	"subkegiatanServices/repository"
	"subkegiatanServices/service"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var subkegiatanSet = wire.NewSet(
	repository.NewSubkegiatanRepositoryImpl,
	wire.Bind(new(repository.SubkegiatanRepository), new(*repository.SubkegiatanRepositoryImpl)),
	service.NewSubkegiatanServiceImpl,
	wire.Bind(new(service.SubkegiatanService), new(*service.SubkegiatanServiceImpl)),
	controller.NewSubkegiatanControllerImpl,
	wire.Bind(new(controller.SubkegiatanController), new(*controller.SubkegiatanControllerImpl)),
)

func InitializedServer() *echo.Echo {
	wire.Build(
		app.GetConnection,
		wire.Value([]validator.Option{}),
		validator.New,
		subkegiatanSet,
		app.NewRouter,
	)
	return nil
}
