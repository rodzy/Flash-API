package routers

import (
	"encoding/json"
	"net/http"

	"github.com/rodzy/flash/db"
)

//ViewProfile is the routing method to display the user profile information
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Param is required", http.StatusBadRequest)
		return
	}
	profile, err := db.FindProfile(ID)
	if err != nil {
		http.Error(w, "Error ocurred"+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
