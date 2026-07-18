package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"backend/db"
	"backend/models"
)

func AddBusHandler(w http.ResponseWriter, r *http.Request) {
	var NewBus models.Buses

	err := json.NewDecoder(r.Body).Decode(&NewBus)
	if err != nil {
		http.Error(w, "Invalid JSON Data", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO buses (numberplate, make, model, seats, company_id) VALUES (?,?,?,?,?)"

	result, err := db.DB.Exec(query, NewBus.Numberplate, NewBus.Make, NewBus.Model, NewBus.Seats, NewBus.CompanyId)
	if err != nil {
		http.Error(w, "Failed to save to database", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	NewBus.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(NewBus)

	fmt.Println("New bus added:", NewBus.Numberplate)
}

func GetBuses(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM buses"
	rows, err := db.DB.Query(query)
	if err != nil {
		http.Error(w, "Failed to fetch buses", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var busList []models.Buses

	for rows.Next() {
		var bus models.Buses
		rows.Scan(&bus.ID, &bus.Numberplate, &bus.Make, &bus.Model, &bus.Seats, &bus.CompanyId)
		busList = append(busList, bus)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(busList)
}
