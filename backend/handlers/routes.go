package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"backend/db"
	"backend/models"
)

func CreateRouteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var route models.Route
	if err := json.NewDecoder(r.Body).Decode(&route); err != nil {
		http.Error(w, "Invalid JSON Data", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO routes (start_location, destination, distance_km, duration_minutes, fare, status) VALUES (?,?,?,?,?,?)`
	result, err := db.DB.Exec(query, route.Start, route.Destination, route.DistanceKm, route.DurationMinutes, route.Fare, route.Status)
	if err != nil {
		http.Error(w, "Failed to save route", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	route.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(route)
	fmt.Println("New route added:", route.Start, "->", route.Destination)
}

func GetRoutes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.DB.Query(`SELECT id, start_location, destination, distance_km, duration_minutes, fare, status FROM routes ORDER BY id DESC`)
	if err != nil {
		http.Error(w, "Failed to fetch routes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var routes []models.Route
	for rows.Next() {
		var route models.Route
		rows.Scan(&route.ID, &route.Start, &route.Destination, &route.DistanceKm, &route.DurationMinutes, &route.Fare, &route.Status)
		routes = append(routes, route)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(routes)
}

func UpdateRouteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var route models.Route
	if err := json.NewDecoder(r.Body).Decode(&route); err != nil {
		http.Error(w, "Invalid JSON Data", http.StatusBadRequest)
		return
	}
	if route.ID == 0 {
		http.Error(w, "Route ID is required", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec(`UPDATE routes SET start_location=?, destination=?, distance_km=?, duration_minutes=?, fare=?, status=? WHERE id=?`, route.Start, route.Destination, route.DistanceKm, route.DurationMinutes, route.Fare, route.Status, route.ID)
	if err != nil {
		http.Error(w, "Failed to update route", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(route)
}

func DeleteRouteHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Valid route ID is required", http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec(`DELETE FROM routes WHERE id=?`, id)
	if err != nil {
		http.Error(w, "Failed to delete route", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Route not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Route deleted"})
}
