package data

import "github.com/dhconnelly/rtreego"

type Driver struct {
	Location     rtreego.Point
	CompanyName  string
	CompanyLogo  string
	CompanyPhone string
}

func (d *Driver) Bounds() *rtreego.Rect {
	return d.Location.ToRect(0.01)
}
