package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/rodzy/flash/db"
	"github.com/rodzy/flash/models"
)
//UploadProfilePic func that handles local storage user avatars to the database by modifying the user data
func UploadProfilePic(w http.ResponseWriter,r*http.Request)  {
	file,hanler,err:=r.FormFile("avatar")
	var filetype=strings.Split(hanler.Filename,".")[1]
	var uploads string="uploads/avatar/"+UserID+"."+filetype

	//Error handleling for copying images from local storage
	f,err:=os.OpenFile(uploads,os.O_WRONLY|os.O_CREATE,0666)
	if err != nil {
		http.Error(w,"Avatar upload error",400)
		return
	}
	_,err=io.Copy(f,file)
	if err != nil {
		http.Error(w,"Copy avatar upload error"+err.Error(),400)
		return
	}

	var user models.User
	user.Avatar=UserID+"."+filetype
	
	status,err:=db.ModifyUser(user,UserID)
	if err != nil || status==false {
		http.Error(w,"Avatar upload error, not the user you're looking",http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusCreated)
}