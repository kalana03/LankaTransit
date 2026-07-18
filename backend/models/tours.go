package models

type Tour struct {
	ID          int    `json:"id"`
	RouteID     int    `json:"route_id"`
	BusID       int    `json:"bus_id"`
	Start       string `json:"start"`
	Destination string `json:"destination"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	BusCompany  string `json:"busCompany"`
	DriverName  string `json:"driverName"`
	DriverPhone string `json:"driverPhone"`
	Status      string `json:"status"`
}
