package routers

import (
	"encoding/json"
	"net/http"

	"github.com/rodzy/flash/db"
	"github.com/rodzy/flash/models"

)

//ModifyUserInfo our method to env the new user info
func ModifyUserInfo(w http.ResponseWriter,r *http.Request)  {
	var user models.User
	err:=json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w,"Incorrect data "+err.Error(),400)
		return
	}
	//Setting a status w/ the global var from the logged user
	var status bool
	status,err=db.ModifyUser(user,UserID)
	if err != nil {
		http.Error(w,"Error trying to insert the data "+err.Error(),400)
		return
	}
	if status==false {
		http.Error(w,"Register not bound in the database ",400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}