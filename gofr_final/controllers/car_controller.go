package controllers

import (
	"github.com/047pegasus/go-crud-mysql/models"
	"github.com/047pegasus/go-crud-mysql/utilities"
)

type CarController struct{}

// AddCarToGarage adds a car to the garage in the database
func (cc *CarController) AddCarToGarage(car models.Car) error {
	db := utilities.GetDB()
	_, err := db.Exec("INSERT INTO cars (brand, model, color, status) VALUES (?, ?, ?, ?)",
		car.Brand, car.Model, car.Color, car.Status)
	if err != nil {
		return err
	}
	return nil
}

// Other CRUD operations for cars in the garage
// UpdateCarInGarage, GetCarsInGarage, RemoveCarFromGarage, etc.

// GetCarInGarage gets a car in the garage by ID
func (cc *CarController) GetCarInGarage(id int) (models.Car, error) {
	db := utilities.GetDB()
	var car models.Car
	err := db.QueryRow("SELECT * FROM cars WHERE id = ?", id).Scan(&car.ID, &car.Brand, &car.Model, &car.Color, &car.Status)
	if err != nil {
		return car, err
	}
	return car, nil
}

// UpdateCarInGarage updates a car in the garage by ID
func (cc *CarController) UpdateCarInGarage(id int, car models.Car) error {
	db := utilities.GetDB()
	_, err := db.Exec("UPDATE cars SET brand = ?, model = ?, color = ?, status = ? WHERE id = ?",
		car.Brand, car.Model, car.Color, car.Status, id)
	if err != nil {
		return err
	}
	return nil
}

// RemoveCarFromGarage removes a car from the garage by ID
func (cc *CarController) RemoveCarFromGarage(id int) error {
	db := utilities.GetDB()
	_, err := db.Exec("DELETE FROM cars WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

// GetCarsInGarage gets all cars in the garage
func (cc *CarController) GetCarsInGarage() ([]models.Car, error) {
	db := utilities.GetDB()
	var cars []models.Car
	rows, err := db.Query("SELECT * FROM cars")
	if err != nil {
		return cars, err
	}
	defer rows.Close()
	for rows.Next() {
		var car models.Car
		err := rows.Scan(&car.ID, &car.Brand, &car.Model, &car.Color, &car.Status)
		if err != nil {
			return cars, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}
