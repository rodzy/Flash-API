package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/rodzy/flash/db"
)
//ShowBanner func that gets the banner from the database
func ShowBanner(w http.ResponseWriter,r*http.Request)  {
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
	file,err:=os.Open("uploads/avatar/"+profile.Banner)
	if err != nil {
		http.Error(w,"No banner found",400)
		return
	}
	_,err=io.Copy(w,file)
	if err != nil {
		http.Error(w,"Error showing the banner",400)
		return
	}
}