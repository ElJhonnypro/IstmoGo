package userManagement

import (
	"API/data"
	userUseModels "API/internal/userManagement/models"
	"API/internal/utils"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

func RegisterUser(Name string, Phone string, Role string, RID *string, RIDPhoto *string, CarID *string, Email string, Password string, Location string, Birthdate *string) data.InsertUserResponse {
	id := uuid.New().String()

	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(Password),
		bcrypt.DefaultCost,
	)

	if utils.HasSymbol(Name) == true {
		return data.InsertUserResponse{
			Message: "Password must contain at least one symbol",
			Success: false,
		}
	}

	if utils.HasSymbol(Phone) == true {
		return data.InsertUserResponse{
			Message: "Phone must not contain symbols",
			Success: false,
		}
	}

	if utils.HasSymbol(Password) == false {
		return data.InsertUserResponse{
			Message: "Password must contain at least one symbol",
			Success: false,
		}
	}

	CarID = utils.EmptyToNil(CarID)
	RID = utils.EmptyToNil(RID)
	RIDPhoto = utils.EmptyToNil(RIDPhoto)

	user := userUseModels.User{
		ID:        id,
		Name:      Name,
		Phone:     Phone,
		Role:      Role,
		RID:       RID,
		RIDPhoto:  RIDPhoto,
		CarID:     CarID,
		Email:     Email,
		Password:  string(hashedPassword),
		Location:  Location,
		Birthdate: Birthdate,
	}
	return data.InsertUser(data.GetDB(), user)
}
