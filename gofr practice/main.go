package main

import (
	// "database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type Car struct {
	LicensePlate string
	EntryTime    time.Time
	Status       string
}

func createCar(w http.ResponseWriter, r *http.Request) {
	var car Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert car into database
	err = insertCar(&car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

func getCars(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status") // optional filter by status

	cars, err := getCars(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cars)
}

func updateCarStatus(w http.ResponseWriter, r *http.Request) {
	licensePlate := mux.Vars(r)["license_plate"]
	newStatus := r.FormValue("status")

	err := updateCarStatus(licensePlate, newStatus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
	licensePlate := mux.Vars(r)["license_plate"]

	err := deleteCar(licensePlate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	// ... database connection and initialization
	http.HandleFunc("/cars", createCar)
	http.HandleFunc("/cars", getCars)
	http.HandleFunc("/cars/{license_plate}", updateCarStatus)
	http.HandleFunc("/cars/{license_plate}", deleteCar)

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
