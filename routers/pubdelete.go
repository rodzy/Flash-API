package routers

import (
	"net/http"

	"github.com/rodzy/flash/db"
)

//DeletePub routes the delete callback in db
func DeletePub(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "No id found", 400)
		return
	}
	err := db.DeletePub(ID, UserID)
	if err != nil {
		http.Error(w, "Pub error trial", 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
