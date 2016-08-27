package service

import (
	"github.com/dhconnelly/rtreego"
	"github.com/gen1us2k/log"
	"github.com/maddevsio/openfreecab-storage/service/data"
)

type RtreeService struct {
	BaseService

	logger log.Logger
	os     *OpenStorage
	tree   *rtreego.Rtree
}

func (r *RtreeService) Name() string {
	return "rtree_service"
}
func (r *RtreeService) Init(os *OpenStorage) error {
	r.os = os
	r.logger = log.NewLogger(r.Name())
	r.tree = rtreego.NewTree(2, 25, 50)

	return nil
}

func (r *RtreeService) Run() error {
	r.tree.Insert(&data.Driver{rtreego.Point{42.872198, 74.584931}, "First"})
	r.tree.Insert(&data.Driver{rtreego.Point{42.871789, 74.583901}, "Second"})
	r.tree.Insert(&data.Driver{rtreego.Point{42.872505, 74.584952}, "Third"})
	r.tree.Insert(&data.Driver{rtreego.Point{42.872819, 74.582002}, "Fourth"})
	r.tree.Insert(&data.Driver{rtreego.Point{42.867044, 74.563705}, "Hided"})
	return nil
}

func (r *RtreeService) AddDriver() {
	// TODO: Add code
}

func (r *RtreeService) Nearest(point rtreego.Point) []data.DriverItem {
	var drivers []data.DriverItem
	items := r.tree.NearestNeighbors(3, point)
	for _, item := range items {
		driver := item.(*data.Driver)
		dr := data.DriverItem{Name: driver.Name}
		dr.SetCoords(driver.Location)
		drivers = append(drivers, dr)
	}
	return drivers
}
