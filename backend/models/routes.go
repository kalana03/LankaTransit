package models

type Route struct {
	ID              int     `json:"id"`
	Start           string  `json:"start"`
	Destination     string  `json:"destination"`
	DistanceKm      float64 `json:"distance_km"`
	DurationMinutes int     `json:"duration_minutes"`
	Fare            float64 `json:"fare"`
	Status          string  `json:"status"`
}
