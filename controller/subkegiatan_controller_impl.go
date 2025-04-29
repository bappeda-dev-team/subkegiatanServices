package controller

import (
	"net/http"
	"subkegiatanServices/model/web"
	"subkegiatanServices/service"

	"github.com/labstack/echo/v4"
)

type SubkegiatanControllerImpl struct {
	subkegiatanService service.SubkegiatanService
}

func NewSubkegiatanControllerImpl(subkegiatanService service.SubkegiatanService) *SubkegiatanControllerImpl {
	return &SubkegiatanControllerImpl{subkegiatanService: subkegiatanService}
}

// Create godoc
// @Summary Create Subkegiatan
// @Description Create a new subkegiatan
// @Tags CREATE Subkegiatan
// @Accept json
// @Produce json
// @Param subkegiatan body web.SubkegiatanCreateRequest true "Create Subkegiatan"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /subkegiatan/create [post]
func (controller *SubkegiatanControllerImpl) Create(c echo.Context) error {
	SubkegiatanCreateRequest := web.SubkegiatanCreateRequest{}
	if err := c.Bind(&SubkegiatanCreateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	subkegiatanResponse, err := controller.subkegiatanService.Create(c.Request().Context(), SubkegiatanCreateRequest)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Create Subkegiatan",
				Data:   err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Create Subkegiatan",
		Data:   subkegiatanResponse,
	})
}

// Update godoc
// @Summary Update Subkegiatan
// @Description Update an existing subkegiatan
// @Tags UPDATE Subkegiatan
// @Accept json
// @Produce json
// @Param id path string true "Subkegiatan ID"
// @Param subkegiatan body web.SubkegiatanUpdateRequest true "Update Subkegiatan"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /subkegiatan/update/{id} [put]
func (controller *SubkegiatanControllerImpl) Update(c echo.Context) error {
	SubkegiatanUpdateRequest := web.SubkegiatanUpdateRequest{}
	if err := c.Bind(&SubkegiatanUpdateRequest); err != nil {
		return err
	}

	id := c.Param("id")

	SubkegiatanUpdateRequest.Id = id

	subkegiatanResponse, err := controller.subkegiatanService.Update(c.Request().Context(), SubkegiatanUpdateRequest)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Update Subkegiatan",
				Data:   err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Update Subkegiatan",
		Data:   subkegiatanResponse,
	})
}

// Delete godoc
// @Summary Delete Subkegiatan
// @Description Delete an existing subkegiatan
// @Tags DELETE Subkegiatan
// @Accept json
// @Produce json
// @Param id path string true "Subkegiatan ID"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /subkegiatan/delete/{id} [delete]
func (controller *SubkegiatanControllerImpl) Delete(c echo.Context) error {
	id := c.Param("id")

	err := controller.subkegiatanService.Delete(c.Request().Context(), id)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Delete Subkegiatan",
				Data:   err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Delete Subkegiatan",
		Data:   nil,
	})
}

// FindById godoc
// @Summary FindById subkegiatan
// @Description Find By Id an existing subkegiatan
// @Tags GET Subkegiatan
// @Accept json
// @Produce json
// @Param id path string true "Subkegiatan ID"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /subkegiatan/detail/{id} [get]
func (controller *SubkegiatanControllerImpl) FindById(c echo.Context) error {
	id := c.Param("id")

	subkegiatanResponse, err := controller.subkegiatanService.FindById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Find Subkegiatan",
				Data:   err.Error(),
			},
		)
	}
	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Find Subkegiatan",
		Data:   subkegiatanResponse,
	})
}

// FindAll godoc
// @Summary FindAll subkegiatan
// @Description FindAll an existing subkegiatan
// @Tags GET Subkegiatan
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /subkegiatan/findall [get]
func (controller *SubkegiatanControllerImpl) FindAll(c echo.Context) error {
	subkegiatanResponse, err := controller.subkegiatanService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Find All Subkegiatan",
				Data:   err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Find All Subkegiatan",
		Data:   subkegiatanResponse,
	})
}

// @Summary FindByKodeSubKegiatan subkegiatan
// @Description FindByKodeSubKegiatan an existing subkegiatan
// @Tags GET Subkegiatan
// @Accept json
// @Produce json
// @Param kode_subkegiatan path string true "Kode Subkegiatan"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /subkegiatan/findkode/{kode_subkegiatan} [get]
func (controller *SubkegiatanControllerImpl) FindByKodeSubKegiatan(c echo.Context) error {
	kodeSubKegiatan := c.Param("kode_subkegiatan")

	subkegiatanResponse, err := controller.subkegiatanService.FindByKodeSubKegiatan(c.Request().Context(), kodeSubKegiatan)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Find Subkegiatan by Kode Subkegiatan",
				Data:   err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Find Subkegiatan by Kode Subkegiatan",
		Data:   subkegiatanResponse,
	})
}
