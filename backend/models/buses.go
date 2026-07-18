package models

type Buses struct {
	ID          int    `json:"id"`
	Numberplate string `json:"numberplate"`
	Make        string `json:"make"`
	Model       string `json:"model"`
	Seats       string `json:"seats"`
	CompanyId   string `json:"company_id"`
}
