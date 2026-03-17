package userManagement

import (
	"API/data"
	userUseModels "API/internal/usermanagement/models"
	"API/internal/utils"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

func RegisterUser(
	Name string,
	Phone string,
	Role string,
	RID *string,
	RIDPhoto *string,
	CarID *string,
	Email string,
	Password string,
	Location string,
	Birthdate *string,
) data.InsertUserResponse {

	if utils.HasSymbol(Name) {
		return data.InsertUserResponse{
			Message: "Name must not contain symbols",
			Success: false,
		}
	}

	if !utils.IsPhoneNumber(Phone) {
		fmt.Print(Phone)
		return data.InsertUserResponse{
			Message: "Phone must contain only digits",
			Success: false,
		}
	}

	if !utils.HasSymbol(Password) {
		return data.InsertUserResponse{
			Message: "Password must contain at least one symbol",
			Success: false,
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return data.InsertUserResponse{
			Message: "Error hashing password",
			Success: false,
		}
	}

	CarID = utils.EmptyToNil(CarID)
	RID = utils.EmptyToNil(RID)
	RIDPhoto = utils.EmptyToNil(RIDPhoto)

	user := userUseModels.User{
		ID:        uuid.New().String(),
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
