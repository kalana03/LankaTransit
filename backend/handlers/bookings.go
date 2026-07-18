package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"backend/db"
	"backend/models"
)

func CreateBookingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid JSON Data", http.StatusBadRequest)
		return
	}

	seatPayload, err := json.Marshal(booking.Seats)
	if err != nil {
		http.Error(w, "Invalid seat data", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO bookings (tour_id, name, nic, start_location, destination, date, time, seats, status, total_price, email, phone, address, bus_company, bus_id, bus_model, driver_id, assistant_id) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	result, err := db.DB.Exec(query, booking.TourID, booking.Name, booking.Nic, booking.Start, booking.Destination, booking.Date, booking.Time, string(seatPayload), booking.Status, booking.TotalPrice, booking.Email, booking.Phone, booking.Address, booking.BusCompany, booking.BusID, booking.BusModel, booking.DriverID, booking.AssistantID)
	if err != nil {
		http.Error(w, "Failed to save booking", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	booking.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
	fmt.Println("New booking added:", booking.Name)
}

func GetBookings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.DB.Query(`SELECT id, tour_id, name, nic, start_location, destination, date, time, seats, status, total_price, email, phone, address, bus_company, bus_id, bus_model, driver_id, assistant_id FROM bookings ORDER BY id DESC`)
	if err != nil {
		http.Error(w, "Failed to fetch bookings", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var booking models.Booking
		var seatsText string
		rows.Scan(&booking.ID, &booking.TourID, &booking.Name, &booking.Nic, &booking.Start, &booking.Destination, &booking.Date, &booking.Time, &seatsText, &booking.Status, &booking.TotalPrice, &booking.Email, &booking.Phone, &booking.Address, &booking.BusCompany, &booking.BusID, &booking.BusModel, &booking.DriverID, &booking.AssistantID)
		json.Unmarshal([]byte(seatsText), &booking.Seats)
		bookings = append(bookings, booking)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookings)
}

func UpdateBookingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid JSON Data", http.StatusBadRequest)
		return
	}
	if booking.ID == 0 {
		http.Error(w, "Booking ID is required", http.StatusBadRequest)
		return
	}

	seatPayload, err := json.Marshal(booking.Seats)
	if err != nil {
		http.Error(w, "Invalid seat data", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec(`UPDATE bookings SET tour_id=?, name=?, nic=?, start_location=?, destination=?, date=?, time=?, seats=?, status=?, total_price=?, email=?, phone=?, address=?, bus_company=?, bus_id=?, bus_model=?, driver_id=?, assistant_id=? WHERE id=?`, booking.TourID, booking.Name, booking.Nic, booking.Start, booking.Destination, booking.Date, booking.Time, string(seatPayload), booking.Status, booking.TotalPrice, booking.Email, booking.Phone, booking.Address, booking.BusCompany, booking.BusID, booking.BusModel, booking.DriverID, booking.AssistantID, booking.ID)
	if err != nil {
		http.Error(w, "Failed to update booking", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(booking)
}

func DeleteBookingHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Valid booking ID is required", http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec(`DELETE FROM bookings WHERE id=?`, id)
	if err != nil {
		http.Error(w, "Failed to delete booking", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Booking deleted"})
}
