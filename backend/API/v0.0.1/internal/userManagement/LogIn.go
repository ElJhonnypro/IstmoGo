// Process to get Token
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

	tryHashPss, _ := bcrypt.GenerateFromPassword(
		[]byte(Password),
		bcrypt.DefaultCost,
	)

	if user.Password == string(tryHashPss) {
		token, err := utils.GenerateToken(user.ID)
		if err != nil {
			return false, ""
		}
		return true, token

	}

	return false, ""

}
