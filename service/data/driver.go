package data

import "github.com/dhconnelly/rtreego"

type Driver struct {
	Location rtreego.Point
	Name     string
}

func (d *Driver) Bounds() *rtreego.Rect {
	return d.Location.ToRect(0.01)
}
