package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/rodzy/flash/db"
	"github.com/rodzy/flash/models"
)

//Email global var
var Email string

//UserID global var
var UserID string

//ProcessJWT to auth header, and checking the credentials from the generated token
func ProcessJWT(token string) (*models.Claim, bool, string, error) {
	pass := []byte("YoooHelloGolang_")
	claim := &models.Claim{}

	manageToken := strings.Split(token, "Bearer")
	if len(manageToken) != 2 {
		return claim, false, string(""), errors.New("invalid token synxtax")
	}
	token = strings.TrimSpace(manageToken[1])

	ftoken, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return pass, nil
	})
	if err == nil {
		_, found, _ := db.UserFound(claim.Email)
		if found == true {
			Email = claim.Email
			UserID = claim.ID.Hex()
		}
		return claim, found, UserID, nil
	}
	if !ftoken.Valid {
		return claim, false, string(""), errors.New("Invalid token credentials")
	}
	return claim, false, string(""), err
}
