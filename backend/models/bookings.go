package models

type Booking struct {
	ID          int      `json:"id"`
	TourID      int      `json:"tour_id"`
	Name        string   `json:"name"`
	Nic         string   `json:"nic"`
	Start       string   `json:"start"`
	Destination string   `json:"destination"`
	Date        string   `json:"date"`
	Time        string   `json:"time"`
	Seats       []string `json:"seats"`
	Status      string   `json:"status"`
	TotalPrice  float64  `json:"totalPrice"`
	Email       string   `json:"email"`
	Phone       string   `json:"phone"`
	Address     string   `json:"address"`
	BusCompany  string   `json:"busCompany"`
	BusID       string   `json:"busId"`
	BusModel    string   `json:"busModel"`
	DriverID    string   `json:"driverId"`
	AssistantID string   `json:"assistantId"`
}
