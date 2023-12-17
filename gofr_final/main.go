package main

import (
	"github.com/047pegasus/go-crud-mysql/handlers"
	"github.com/047pegasus/go-crud-mysql/utilities"

	"github.com/047pegasus/go-crud-mysql/controllers"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()
	// Connect to MySQL database
	db, err := utilities.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Initialize controllers
	carController := controllers.CarController{}

	// Initialize handlers
	carHandler := handlers.NewCarHandler(&carController, db)

	app.Server.ValidateHeaders = false

	// Initialize router
app.POST( "/cars", carHandler.AddCarToGarage)
	app.GET( "/cars", carHandler.GetCarsInGarage)
	app.GET( "/cars/{id}", carHandler.GetCarInGarage)
	app.PUT( "/cars/{id}", carHandler.UpdateCarInGarage)
	app.DELETE( "/cars/{id}", carHandler.RemoveCarFromGarage)

	app.Server.HTTP.Port = 3000

	app.Start()
}
