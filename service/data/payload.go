package data

type DriverData struct {
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	CompanyName string  `json:"company_name"`
}

type DefaultResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
