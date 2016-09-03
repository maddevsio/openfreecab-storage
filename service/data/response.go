package data

import "github.com/dhconnelly/rtreego"

type DriverItem struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
type NearestResponse struct {
	Success   bool      `json:"success"`
	Companies []Company `json:"companies"`
}
type Company struct {
	Name     string       `json:"name"`
	Icon     string       `json:"icon"`
	Contacts []Contact    `json:"contacts"`
	Drivers  []DriverItem `json:"drivers"`
}

func (d *DriverItem) SetCoords(p rtreego.Point) {
	d.Lat = p[0]
	d.Lon = p[1]
}
