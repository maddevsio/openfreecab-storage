package service

import (
	"github.com/dhconnelly/rtreego"
	"github.com/gen1us2k/log"
	"github.com/maddevsio/openfreecab-storage/service/data"
	"github.com/maddevsio/openfreecab-storage/storage"
)

type RtreeService struct {
	BaseService

	logger log.Logger
	os     *OpenStorage
	ds     *storage.DriverStorage
}

func (r *RtreeService) Name() string {
	return "rtree_service"
}
func (r *RtreeService) Init(os *OpenStorage) error {
	r.os = os
	r.logger = log.NewLogger(r.Name())
	r.ds = storage.NewDriverStorage(20)

	return nil
}

func (r *RtreeService) Run() error {
	return nil
}

func (r *RtreeService) AddDriver(driver *data.DriverData) {
	r.ds.AddDriver(&data.Driver{
		Location:    rtreego.Point{driver.Lat, driver.Lon},
		CompanyName: driver.CompanyName,
	})
}

func (r *RtreeService) Nearest(point rtreego.Point) []data.DriverItem {
	return r.ds.Nearest(point)
}

func (r *RtreeService) CleanStorageByCompanyName(companyName string) {
	r.ds.RemoveDriversByCompanyName(companyName)
}
