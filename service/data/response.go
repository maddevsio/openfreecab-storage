package data

import "github.com/dhconnelly/rtreego"

type DriverItem struct {
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Name string  `json:"name"`
}
type NearestResponse struct {
	Success   bool      `json:"success"`
	Companies []Company `json:"companies"`
}
type Company struct {
	Name    string       `json:"name"`
	Drivers []DriverItem `json:"drivers"`
}

func (d *DriverItem) SetCoords(p rtreego.Point) {
	d.Lat = p[0]
	d.Lon = p[1]
}
