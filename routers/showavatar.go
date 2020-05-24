package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/rodzy/flash/db"
)
//ShowAvatar func that gets the avatar from the database
func ShowAvatar(w http.ResponseWriter,r*http.Request)  {
	ID:=r.URL.Query().Get("id")
	if len(ID)<1 {
		http.Error(w,"No id found",400)
		return
	}
	profile,err:=db.FindProfile(ID)
	if err != nil {
		http.Error(w,"User not found",400)
		return
	}
	file,err:=os.Open("uploads/avatar/"+profile.Avatar)
	if err != nil {
		http.Error(w,"No avatar found",400)
		return
	}
	_,err=io.Copy(w,file)
	if err != nil {
		http.Error(w,"Error showing the avatar",400)
		return
	}
}