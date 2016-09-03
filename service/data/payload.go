package data

type DriverData struct {
	Lat          float64 `json:"lat"`
	Lon          float64 `json:"lon"`
	CompanyName  string  `json:"company_name"`
	CompanyPhone string  `json:"company_phone"`
	CompanyLogo  string  `json:"company_logo"`
}

type DefaultResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
