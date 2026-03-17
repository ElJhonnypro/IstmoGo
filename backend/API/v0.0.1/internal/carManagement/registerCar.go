package carManagement

import (
	data "API/data"
	carsUseModels "API/internal/carmanagement/models"

	"github.com/google/uuid"
)

func RegisterCar(Plate, Model, Color, Photo string) data.InsertCarResponse {
	id := uuid.New().String()
	car := carsUseModels.Car{
		ID:    id,
		Plate: Plate,
		Model: Model,
		Color: Color,
		Photo: Photo,
	}
	return data.InsertCar(data.GetDB(), car)
}
