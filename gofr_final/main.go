package main

import (
	"github.com/047pegasus/go-crud-mysql/handlers"
	"github.com/047pegasus/go-crud-mysql/utilities"

	"github.com/047pegasus/go-crud-mysql/controllers"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()
	db, err := utilities.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	carController := controllers.CarController{}

	carHandler := handlers.NewCarHandler(&carController, db)

	app.Server.ValidateHeaders = false

	app.POST( "/cars", carHandler.AddCarToGarage)
	app.GET( "/cars", carHandler.GetCarsInGarage)
	app.GET( "/cars/{id}", carHandler.GetCarInGarage)
	app.PUT( "/cars/{id}", carHandler.UpdateCarInGarage)
	app.DELETE( "/cars/{id}", carHandler.RemoveCarFromGarage)

	app.Server.HTTP.Port = 3000

	app.Start()
}
