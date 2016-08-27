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
	return nil
}

func (h *HTTPService) Run() error {
	h.e.Run(standard.New(h.os.Config().HTTPBindAddr))
	return nil
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
