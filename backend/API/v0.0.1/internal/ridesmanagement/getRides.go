package rideManagement

import (
	"API/data"
	rideModels "API/internal/ridesmanagement/models"
	utils "API/internal/utils"
)

func GetAllRides() ([]rideModels.Ride, error) {
	return data.GetAllRides(data.GetDB())
}

func GetRidesNearLocation(lat, lng float64) ([]rideModels.Ride, error) {
	AllRides, err := GetAllRides()
	if err != nil {
		return nil, err
	}
	nearRides := []rideModels.Ride{}

	for _, ride := range AllRides {
		if utils.IsWithinRadius(ride.StartLat, ride.StartLng, lat, lng, 5) {
			nearRides = append(nearRides, ride)
		}
	}
	return nearRides, nil
}

func GetRidesById(rideID string) (rideModels.Ride, error) {
	return data.GetRideByID(data.GetDB(), rideID)
}
