package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rodzy/flash/db"
)
//FollowersList func
func FollowersList(w http.ResponseWriter,r *http.Request)  {
	typeUser:=r.URL.Query().Get("type")
	page:=r.URL.Query().Get("page")
	search:=r.URL.Query().Get("search")
	pagTemp,err:=strconv.Atoi(page)
	if err != nil {
		http.Error(w,"Param not found or 0",http.StatusBadRequest)
		return
	}
	pag:=int64(pagTemp)

	result,status:=db.GetFollowerList(UserID,pag,search,typeUser)
	if status==false {
		http.Error(w,"Couldn't read users info",http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}