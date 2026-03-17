package rideManagement

import (
	"API/data"
	rideModels "API/internal/ridesmanagement/models"

	"github.com/google/uuid"
)

func CreateRide(
	clientID string,
	startLat, startLng,
	endLat, endLng,
	distanceKm, price float64,
) (rideModels.Ride, error) {

	ride := rideModels.Ride{
		ID:         uuid.New().String(),
		ClientID:   clientID,
		StartLat:   startLat,
		StartLng:   startLng,
		EndLat:     endLat,
		EndLng:     endLng,
		DistanceKm: distanceKm,
		Price:      price,
		Status:     "requested",
	}

	err := data.InsertRide(data.GetDB(), ride)
	return ride, err
}
