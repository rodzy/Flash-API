package middlewares

import (
	"net/http"
	"github.com/rodzy/flash/db"
)

//CheckDataBase method to check the database http
func CheckDataBase(next http.HandlerFunc) http.HandlerFunc  {
	return func (w http.ResponseWriter,r *http.Request)  {
		//Checking for errors in my Mongo DB connection
		if db.CheckConnection()==0 {
			http.Error(w,"Couldn't find any connection",500)
			return
		}
		next.ServeHTTP(w,r)
	}
}