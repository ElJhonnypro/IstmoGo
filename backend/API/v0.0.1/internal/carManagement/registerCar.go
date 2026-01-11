package carManagement

import (
	"API/data"
	carsUseModels "API/internal/carManagement/models"

	"github.com/google/uuid"
)

func RegisterCar(Plate, Model, Color, Photo, UserId string) data.InsertCarResponse {
	id := uuid.New().String()
	car := carsUseModels.Car{
		ID:     id,
		Plate:  Plate,
		Model:  Model,
		UserId: UserId,
		Color:  Color,
		Photo:  Photo,
	}
	return data.InsertCar(data.GetDB(), car)
}
