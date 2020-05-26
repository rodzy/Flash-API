package routers

import (
	"net/http"

	"github.com/rodzy/flash/db"
	"github.com/rodzy/flash/models"
)

//UnfollowUser func routing to unfollow certain users
func UnfollowUser(w http.ResponseWriter,r*http.Request)  {
	ID:=r.URL.Query().Get("id")
	if len(ID)<1 {
		http.Error(w,"Id not found",http.StatusBadRequest)
		return
	}
	var follower models.Follower
	follower.UserID=UserID
	follower.UserFollowed=ID

	status,err:=db.UnfollowUser(follower)
	if err != nil {
		http.Error(w,"Unfollow relation cannot process",http.StatusBadRequest)
		return
	}
	if status==false {
		http.Error(w,"No changes been done for followers"+err.Error(),http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}