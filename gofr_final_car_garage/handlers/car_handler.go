package handlers

import (
	"database/sql"
	"encoding/json"
	"strconv"

	"github.com/047pegasus/go-crud-mysql/controllers"
	"github.com/047pegasus/go-crud-mysql/models"
	"gofr.dev/pkg/gofr"
)

type CarHandler struct {
	carController *controllers.CarController
}

func NewCarHandler(cc *controllers.CarController, db *sql.DB) *CarHandler {

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS cars (id SERIAL PRIMARY KEY, brand TEXT, model TEXT, color TEXT, status TEXT )")

	if err != nil {

		// fmt.Println(err);
		panic(err)

	}
	return &CarHandler{carController: cc}
}

func (ch *CarHandler) AddCarToGarage(ctx *gofr.Context) (interface{}, error) {
	var car models.Car
	err:=ctx.BindStrict(&car)
	if err !=nil {
		return nil, err
	}

	if err := ch.carController.AddCarToGarage(car); err != nil {
		return nil, err
	}
	return "Car added successfully!", nil
}

func (ch *CarHandler) GetCarsInGarage(ctx *gofr.Context) (interface{}, error) {
	cars, err := ch.carController.GetCarsInGarage()
	b, err := json.Marshal(cars)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

func (ch *CarHandler) GetCarInGarage(ctx *gofr.Context) (interface{}, error) {
	id,err:=strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}
	car, err := ch.carController.GetCarInGarage(id)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (ch *CarHandler) UpdateCarInGarage(ctx *gofr.Context) (interface{}, error) {
	id,err:=strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}
	var car models.Car
	err=ctx.BindStrict(&car)
	if err !=nil {
		return nil, err
	}
	updated_car,err := ch.carController.UpdateCarInGarage(id, car)
	if err != nil {
		return nil, err
	}
	return updated_car, nil
}

func (ch *CarHandler) RemoveCarFromGarage(ctx *gofr.Context) (interface{}, error) {

	id,err:=strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}

	if err := ch.carController.RemoveCarFromGarage(id); err != nil {
		return nil, err
	}
	return nil, nil
}
