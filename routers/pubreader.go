package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rodzy/flash/db"
)

//PubReader func to show pubs for people request like
func PubReader(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "No id found", 400)
		return
	}
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "No page found", 400)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Page not found or ==0", 400)
		return
	}
	p := int64(page)
	response, status := db.ReadPubs(ID, p)
	if status == false {
		http.Error(w, "Error reading pubs", 400)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
