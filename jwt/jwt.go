package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/rodzy/flash/models"
)

//Spawn it's the generator for our JWt
func Spawn(u models.User) (string, error) {
	pass := []byte("YoooHelloGolang_")
	payload := jwt.MapClaims{
		"email":     u.Email,
		"name":      u.Name,
		"lastname":  u.LastName,
		"birthdate": u.BirthDate,
		"bio":       u.Bio,
		"location":  u.Location,
		"website":   u.WebSite,
		"_id":       u.ID.Hex(),
		"expire":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(pass)
	if err != nil {
		return token, err
	}
	return token, nil
}
