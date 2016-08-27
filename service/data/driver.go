package data

import "github.com/dhconnelly/rtreego"

type Driver struct {
	Location    rtreego.Point
	CompanyName string
}

func (d *Driver) Bounds() *rtreego.Rect {
	return d.Location.ToRect(0.01)
}
