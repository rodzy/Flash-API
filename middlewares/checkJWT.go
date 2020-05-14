package middlewares

import (
	"net/http"

	"github.com/rodzy/flash/routers"
)

//CheckJWT invoke a routine from routes
func CheckJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessJWT(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w,"cant find token "+err.Error(),http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
