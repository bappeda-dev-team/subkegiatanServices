package app

import (
	"subkegiatanServices/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(subkegiatanController controller.SubkegiatanController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/subkegiatan/create", subkegiatanController.Create)
	e.PUT("/subkegiatan/update/:id", subkegiatanController.Update)
	e.DELETE("/subkegiatan/delete/:id", subkegiatanController.Delete)
	e.GET("/subkegiatan/detail/:id", subkegiatanController.FindById)
	e.GET("/subkegiatan/findall", subkegiatanController.FindAll)
	e.GET("/subkegiatan/findkode/:kode_subkegiatan", subkegiatanController.FindByKodeSubKegiatan)

	return e
}
