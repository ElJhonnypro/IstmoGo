package rideManagement

import (
	"API/data"
	rideModels "API/internal/ridesmanagement/models"
)

func GetRideByUser(userID string) ([]rideModels.Ride, error) {
	return data.GetActiveRidesByUserID(data.GetDB(), userID)
}
