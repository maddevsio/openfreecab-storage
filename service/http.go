package service

import (
	"net/http"
	"strconv"

	"github.com/dhconnelly/rtreego"
	"github.com/gen1us2k/log"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/maddevsio/openfreecab-storage/service/data"
)

type HTTPService struct {
	BaseService

	os          *OpenStorage
	logger      log.Logger
	e           *echo.Echo
	treeService *RtreeService
}

func (h *HTTPService) Name() string {
	return "http_api"
}

func (h *HTTPService) Init(os *OpenStorage) error {
	h.os = os
	h.logger = log.NewLogger(h.Name())
	e := echo.New()
	h.treeService = h.os.RtreeService()
	h.e = e
	h.e.GET("/nearest/:lat/:lon", h.nearestNeighbors)
	h.e.POST("/add/", h.addData)
	h.e.POST("/clean/:companyName/", h.cleanByCompanyName)
	return nil
}

func (h *HTTPService) Run() error {
	h.e.Run(standard.New(h.os.Config().HTTPBindAddr))
	return nil
}

func (h *HTTPService) addData(c echo.Context) error {
	h.logger.Info("Adding Driver")
	driverData := &data.DriverData{}
	if err := c.Bind(driverData); err != nil {
		return c.JSON(http.StatusBadRequest, &data.DefaultResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	h.treeService.AddDriver(driverData)
	return c.JSON(http.StatusOK, &data.DefaultResponse{
		Success: true,
		Message: "",
	})
}

func (h *HTTPService) nearestNeighbors(c echo.Context) error {
	lat, err := strconv.ParseFloat(c.Param("lat"), 64)
	if err != nil {
		return err
	}
	lon, err := strconv.ParseFloat(c.Param("lon"), 64)
	if err != nil {
		return err
	}
	nearest := h.treeService.Nearest(rtreego.Point{lat, lon})
	return c.JSON(http.StatusOK, &data.NearestResponse{
		Success: true,
		Drivers: nearest,
	})
}

func (h *HTTPService) cleanByCompanyName(c echo.Context) error {
	companyName := c.Param("companyName")
	h.treeService.CleanStorageByCompanyName(companyName)
	return c.JSON(http.StatusOK, &data.DefaultResponse{
		Success: true,
		Message: "Cleaned",
	})
}
