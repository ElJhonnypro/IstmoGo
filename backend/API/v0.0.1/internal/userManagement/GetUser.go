package userManagement

import (
	"API/data"
	userUseModels "API/internal/userManagement/models"
)

func GetALLUsers(admin bool) (users []userUseModels.User, err error) {

	if !admin {
		return []userUseModels.User{}, err
	}
	return data.AdmingetUsers(data.GetDB())

}

func GetUserByID(UserID string) (user userUseModels.User, err error) {

	return data.GetUserByID(data.GetDB(), true, UserID)

}
