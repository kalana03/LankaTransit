package main

import (
	"backend/db"
	"backend/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	db.InitDB()
	http.HandleFunc("/ping", handlers.PingHandler)

	http.HandleFunc("/api/companies/create", handlers.CreateCompanyHandler)

	http.HandleFunc("/api/companies", handlers.GetCompanies)

	http.HandleFunc("/api/buses", handlers.GetBuses)
	http.HandleFunc("/api/buses/create", handlers.AddBusHandler)

	http.HandleFunc("/api/routes", handlers.GetRoutes)
	http.HandleFunc("/api/routes/create", handlers.CreateRouteHandler)
	http.HandleFunc("/api/routes/update", handlers.UpdateRouteHandler)
	http.HandleFunc("/api/routes/delete", handlers.DeleteRouteHandler)

	http.HandleFunc("/api/tours", handlers.GetTours)
	http.HandleFunc("/api/tours/create", handlers.CreateTourHandler)
	http.HandleFunc("/api/tours/update", handlers.UpdateTourHandler)
	http.HandleFunc("/api/tours/delete", handlers.DeleteTourHandler)

	http.HandleFunc("/api/bookings", handlers.GetBookings)
	http.HandleFunc("/api/bookings/create", handlers.CreateBookingHandler)
	http.HandleFunc("/api/bookings/update", handlers.UpdateBookingHandler)
	http.HandleFunc("/api/bookings/delete", handlers.DeleteBookingHandler)

	fmt.Println("Server is booting up on port 8080...")

	c := cors.Default()
	handler := c.Handler(http.DefaultServeMux)

	err := http.ListenAndServe(":8080", handler)

	if err != nil {
		log.Fatal("Server crashed: ", err)
	}
}
