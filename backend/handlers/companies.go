package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"backend/db"
	"backend/models"
)

func CreateCompanyHandler(w http.ResponseWriter, r *http.Request) {
	var NewCompany models.Companies

	err := json.NewDecoder(r.Body).Decode(&NewCompany)
	if err != nil {
		http.Error(w, "Invalid JSON Data", http.StatusBadRequest)
		return
	}

	tx, err := db.DB.Begin()
	if err != nil {
		http.Error(w, "Failed to start database transaction", http.StatusInternalServerError)
		return
	}

	result, err := tx.Exec("INSERT INTO COMPANIES (Name, Owner, ContactNo, Email, Address) VALUES (?,?,?,?,?)", NewCompany.Name, NewCompany.Owner, NewCompany.ContactNo, NewCompany.Email, NewCompany.Address)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to save company to database", http.StatusInternalServerError)
		return
	}

	companyID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to read inserted company id", http.StatusInternalServerError)
		return
	}

	username, password := buildDefaultCompanyCredentials(int(companyID))
	_, err = tx.Exec("INSERT INTO company_logins (company_id, username, password) VALUES (?,?,?)", companyID, username, password)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to save company login credentials", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to commit company creation", http.StatusInternalServerError)
		return
	}

	NewCompany.ID = int(companyID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(NewCompany)

	fmt.Println("New company created:", NewCompany.Name)
}

func buildDefaultCompanyCredentials(companyID int) (string, string) {
	username := "company_" + strconv.Itoa(companyID)
	password := "company@" + strconv.Itoa(companyID)
	return username, password
}

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM COMPANIES"
	rows, err := db.DB.Query(query)

	if err != nil {
		http.Error(w, "Failed to fetch companies", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var CompanyList []models.Companies

	for rows.Next() {
		var company models.Companies

		rows.Scan(&company.ID, &company.Name, &company.Owner, &company.ContactNo, &company.Email, &company.Address)

		CompanyList = append(CompanyList, company)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CompanyList)
}
