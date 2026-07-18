package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "./lanka-transit-db")
	if err != nil {
		log.Fatal("Could not open database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	fmt.Println("✅ SQLite Database Connected!")

	createTables()

}

func createTables() {
	companyQuery := `
	CREATE TABLE IF NOT EXISTS companies (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		owner TEXT NOT NULL,
		contactno TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		address TEXT NOT NULL
	);`

	_, err := DB.Exec(companyQuery)
	if err != nil {
		log.Fatal("Failed to create companies table: ", err)
	}

	companyLoginQuery := `
	CREATE TABLE IF NOT EXISTS company_logins (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		company_id INTEGER NOT NULL UNIQUE,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		FOREIGN KEY(company_id) REFERENCES companies(id)
	);`

	_, err = DB.Exec(companyLoginQuery)
	if err != nil {
		log.Fatal("Failed to create company_logins table: ", err)
	}

	busQuery := `
	CREATE TABLE IF NOT EXISTS buses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		numberplate TEXT NOT NULL UNIQUE,
		make TEXT NOT NULL,
		model TEXT NOT NULL,
		seats INTEGER NOT NULL,
		company_id INTEGER NOT NULL,
		FOREIGN KEY(company_id) REFERENCES companies(id)
	);`

	_, err = DB.Exec(busQuery)
	if err != nil {
		log.Fatal("Failed to create buses table: ", err)
	}

	routeQuery := `
	CREATE TABLE IF NOT EXISTS routes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		start_location TEXT NOT NULL,
		destination TEXT NOT NULL,
		distance_km REAL NOT NULL,
		duration_minutes INTEGER NOT NULL,
		fare REAL NOT NULL,
		status TEXT NOT NULL DEFAULT 'active'
	);`

	_, err = DB.Exec(routeQuery)
	if err != nil {
		log.Fatal("Failed to create routes table: ", err)
	}

	tourQuery := `
	CREATE TABLE IF NOT EXISTS tours (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		route_id INTEGER NOT NULL,
		bus_id INTEGER NOT NULL,
		start_location TEXT NOT NULL,
		destination TEXT NOT NULL,
		date TEXT NOT NULL,
		time TEXT NOT NULL,
		bus_company TEXT NOT NULL,
		driver_name TEXT NOT NULL,
		driver_phone TEXT NOT NULL,
		status TEXT NOT NULL DEFAULT 'scheduled'
	);`

	_, err = DB.Exec(tourQuery)
	if err != nil {
		log.Fatal("Failed to create tours table: ", err)
	}

	bookingQuery := `
	CREATE TABLE IF NOT EXISTS bookings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tour_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		nic TEXT NOT NULL,
		start_location TEXT NOT NULL,
		destination TEXT NOT NULL,
		date TEXT NOT NULL,
		time TEXT NOT NULL,
		seats TEXT NOT NULL,
		status TEXT NOT NULL DEFAULT 'Active',
		total_price REAL NOT NULL,
		email TEXT NOT NULL,
		phone TEXT NOT NULL,
		address TEXT NOT NULL,
		bus_company TEXT NOT NULL,
		bus_id TEXT NOT NULL,
		bus_model TEXT NOT NULL,
		driver_id TEXT NOT NULL,
		assistant_id TEXT NOT NULL
	);`

	_, err = DB.Exec(bookingQuery)
	if err != nil {
		log.Fatal("Failed to create bookings table: ", err)
	}

	fmt.Println("✅ Tables created successfully!")
}
