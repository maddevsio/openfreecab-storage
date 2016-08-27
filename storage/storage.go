package storage

import (
	"fmt"
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

func (ds *DriverStorage) Nearest(point rtreego.Point) []data.DriverItem {
	ds.Lock()
	var drivers []data.DriverItem
	items := ds.geoIndex.NearestNeighbors(
		ds.nearestNeighbors, point,
	)
	for _, item := range items {
		driver := item.(*data.Driver)
		dr := data.DriverItem{Name: driver.CompanyName}
		dr.SetCoords(driver.Location)
		drivers = append(drivers, dr)
	}
	ds.Unlock()
	return drivers
}
func (ds *DriverStorage) RemoveDriversByCompanyName(companyName string) {
	ds.Lock()
	drivers := ds.companyDrivers[companyName]
	for _, driver := range drivers {
		deleted := ds.geoIndex.Delete(driver)
		if deleted {
			fmt.Println("Item deleted")
		}
	}
	ds.Unlock()
}
