package userManagement

import (
	"API/data"
	"API/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

func LogIn(Identifier, Password string) (bool, string) {
	user, err := data.VerifyLogIn(data.GetDB(), Identifier)
	if err != nil {
		return false, ""
	}

	//fmt.Println(user)

	try := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(Password),
	)

	if try == nil {
		token, err := utils.GenerateToken(user.ID)
		if err != nil {

			return false, ""
		}
		return true, token

	}

	return false, ""

}
