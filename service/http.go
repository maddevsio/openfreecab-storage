package service

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	_ "net/http/pprof"

	"github.com/dhconnelly/rtreego"
	"github.com/gen1us2k/log"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	h.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	h.e.GET("/nearest/:lat/:lon", h.nearestNeighbors)
	h.e.Static("/static", "static")
	h.e.POST("/add/", h.addData)
	h.e.POST("/clean/:companyName/", h.cleanByCompanyName)
	return nil
}

func (h *HTTPService) Run() error {
	h.e.Start(h.os.Config().HTTPBindAddr)
	return nil
}

func (h *HTTPService) addData(c echo.Context) error {
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
	var companies []data.Company

	nearest := h.treeService.Nearest(rtreego.Point{lat, lon})
	for key, value := range nearest {
		info := data.Companies[key]
		companies = append(companies, data.Company{
			Name:     key,
			Icon:     fmt.Sprintf("%s%s", h.os.Config().BaseURL, info.Icon),
			Contacts: info.Contacts,
			Drivers:  value,
		})
	}
	return c.JSON(http.StatusOK, &data.NearestResponse{
		Success:   true,
		Companies: companies,
	})
}

func (h *HTTPService) cleanByCompanyName(c echo.Context) error {
	companyName := c.Param("companyName")
	companyName, err := url.QueryUnescape(companyName)
	if err != nil {
		h.logger.Errorf("Got error decoding companyName", err)
	}
	h.logger.Infof("Cleaning data for %s", companyName)
	h.treeService.CleanStorageByCompanyName(companyName)
	return c.JSON(http.StatusOK, &data.DefaultResponse{
		Success: true,
		Message: "Cleaned",
	})
}
