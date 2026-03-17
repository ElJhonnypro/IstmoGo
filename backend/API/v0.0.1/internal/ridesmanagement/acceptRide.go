package rideManagement

import (
	"API/data"
	rideModels "API/internal/ridesmanagement/models"
)

func AcceptRide(rideID, driverID string) (rideModels.Ride, error) {
	ride, err := data.GetRideByID(data.GetDB(), rideID)
	if err != nil {
		return rideModels.Ride{}, err
	}

	if ride.Status != "requested" {
		return rideModels.Ride{}, data.ErrRideNotAvailable
	}
	ride.Status = "accepted"
	ride.UberID = &driverID

	err = data.UpdateRide(data.GetDB(), ride)
	return ride, err
}
