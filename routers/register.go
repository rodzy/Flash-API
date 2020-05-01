package routers

import (
	"encoding/json"
	"net/http"
	"github.com/rodzy/flash/db"
	"github.com/rodzy/flash/models"
)

//Register func to create an user in our MongoDB
func Register(w http.ResponseWriter,r *http.Request)  {
	var t models.User
	//Streaming the json file
	err:=json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w,"Error on streaming data"+err.Error(),400)
		return
	}
	if len(t.Email)==0 {
		http.Error(w,"Email required",400)
		return
	}
	if len(t.Password)< 8 {
		http.Error(w,"Password must be 8 caracters",400)
		return
	}
	//Finding the user in DB
	_,userfound,_:= db.UserFound(t.Email)
	if userfound==true {
		http.Error(w,"User is already registered",400)
		return
	}
	//Registering user
	_,status,err:=db.InsertUser(t)
	if err != nil {
		http.Error(w,"Error trying to register the user",400)
		return
	}
	if status==false {
		http.Error(w,"Couldn't register the user",400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}