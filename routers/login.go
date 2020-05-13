package routers

import (
	"encoding/json"
	"net/http"

	"github.com/rodzy/flash/db"
	"github.com/rodzy/flash/jwt"
	"github.com/rodzy/flash/models"
)

//Login func for the http endpoint
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Email or password invalid"+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w, "Email invalid", 400)
		return
	}
	us, exist := db.LoginValidation(user.Email, user.Password)
	if exist == false {
		http.Error(w, "Server error", 400)
		return
	}
	//Using the jwt model for the token usage
	key, err := jwt.Spawn(us)
	if err != nil {
		http.Error(w, "Token generation error"+err.Error(), 400)
		return
	}
	response := models.LoginResponse{
		Token: key,
	}
	//Setting the token to the Json and the frontend
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
