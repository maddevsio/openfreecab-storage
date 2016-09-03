package storage

import (
	"sync"

	"github.com/dhconnelly/rtreego"
	"github.com/maddevsio/openfreecab-storage/service/data"
)

type DriverStorage struct {
	sync.RWMutex
	companyDrivers   map[string][]*data.Driver
	geoIndex         *rtreego.Rtree
	nearestNeighbors int
}

func NewDriverStorage(nearestNeighbors int) *DriverStorage {
	ds := new(DriverStorage)
	ds.companyDrivers = make(map[string][]*data.Driver)
	ds.geoIndex = rtreego.NewTree(2, 25, 50)
	ds.nearestNeighbors = nearestNeighbors
	return ds
}

func (ds *DriverStorage) AddDriver(driver *data.Driver) {
	ds.Lock()
	ds.geoIndex.Insert(driver)
	drivers := ds.companyDrivers[driver.CompanyName]
	drivers = append(drivers, driver)
	ds.companyDrivers[driver.CompanyName] = drivers
	ds.Unlock()
}

func (ds *DriverStorage) Nearest(point rtreego.Point) map[string][]data.DriverItem {
	ds.Lock()
	drivers := make(map[string][]data.DriverItem)
	items := ds.geoIndex.NearestNeighbors(
		ds.nearestNeighbors, point,
	)
	for _, item := range items {
		if item == nil {
			continue
		}
		driver := item.(*data.Driver)
		if driver == nil {
			continue
		}

		currentDrivers := drivers[driver.CompanyName]

		dr := data.DriverItem{Name: driver.CompanyName}
		dr.SetCoords(driver.Location)
		currentDrivers = append(currentDrivers, dr)
		drivers[driver.CompanyName] = currentDrivers
	}
	ds.Unlock()
	return drivers
}
func (ds *DriverStorage) RemoveDriversByCompanyName(companyName string) {
	ds.Lock()
	drivers := ds.companyDrivers[companyName]
	for _, driver := range drivers {
		ds.geoIndex.Delete(driver)
	}
	ds.Unlock()
}
