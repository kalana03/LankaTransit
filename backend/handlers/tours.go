package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"backend/db"
	"backend/models"
)

func CreateTourHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var tour models.Tour
	if err := json.NewDecoder(r.Body).Decode(&tour); err != nil {
		http.Error(w, "Invalid JSON Data", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO tours (route_id, bus_id, start_location, destination, date, time, bus_company, driver_name, driver_phone, status) VALUES (?,?,?,?,?,?,?,?,?,?)`
	result, err := db.DB.Exec(query, tour.RouteID, tour.BusID, tour.Start, tour.Destination, tour.Date, tour.Time, tour.BusCompany, tour.DriverName, tour.DriverPhone, tour.Status)
	if err != nil {
		http.Error(w, "Failed to save tour", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	tour.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tour)
	fmt.Println("New tour added:", tour.Start, "->", tour.Destination)
}

func GetTours(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.DB.Query(`SELECT id, route_id, bus_id, start_location, destination, date, time, bus_company, driver_name, driver_phone, status FROM tours ORDER BY id DESC`)
	if err != nil {
		http.Error(w, "Failed to fetch tours", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tours []models.Tour
	for rows.Next() {
		var tour models.Tour
		rows.Scan(&tour.ID, &tour.RouteID, &tour.BusID, &tour.Start, &tour.Destination, &tour.Date, &tour.Time, &tour.BusCompany, &tour.DriverName, &tour.DriverPhone, &tour.Status)
		tours = append(tours, tour)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tours)
}

func UpdateTourHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var tour models.Tour
	if err := json.NewDecoder(r.Body).Decode(&tour); err != nil {
		http.Error(w, "Invalid JSON Data", http.StatusBadRequest)
		return
	}
	if tour.ID == 0 {
		http.Error(w, "Tour ID is required", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec(`UPDATE tours SET route_id=?, bus_id=?, start_location=?, destination=?, date=?, time=?, bus_company=?, driver_name=?, driver_phone=?, status=? WHERE id=?`, tour.RouteID, tour.BusID, tour.Start, tour.Destination, tour.Date, tour.Time, tour.BusCompany, tour.DriverName, tour.DriverPhone, tour.Status, tour.ID)
	if err != nil {
		http.Error(w, "Failed to update tour", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tour)
}

func DeleteTourHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		var payload struct {
			ID int `json:"id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Invalid JSON Data", http.StatusBadRequest)
			return
		}
		idStr = strconv.Itoa(payload.ID)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Valid tour ID is required", http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec(`DELETE FROM tours WHERE id=?`, id)
	if err != nil {
		http.Error(w, "Failed to delete tour", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Tour not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Tour deleted"})
}
