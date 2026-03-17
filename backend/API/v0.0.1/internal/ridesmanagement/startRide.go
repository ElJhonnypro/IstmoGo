package rideManagement

import (
	"API/data"
	rideModels "API/internal/ridesmanagement/models"
)

func StartRide(rideID string) (rideModels.Ride, error) {
	ride, err := data.GetRideByID(data.GetDB(), rideID)
	if err != nil {
		return rideModels.Ride{}, err
	}
	ride.Status = "started"
	err = data.UpdateRide(data.GetDB(), ride)
	return ride, err
}
