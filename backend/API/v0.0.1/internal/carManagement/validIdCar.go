package carManagement

import (
	"API/data"
	"fmt"
)

func ValidIdCar(Id string) bool {
	fmt.Println(Id)
	if car, _ := data.GetCardById(data.GetDB(), Id); car.ID == "" {
		return false
	}

	return true

}

// INFO
