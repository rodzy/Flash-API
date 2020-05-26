package routers

import (
	"encoding/json"
	"net/http"

	"github.com/rodzy/flash/db"
	"github.com/rodzy/flash/models"
)

//AskFollower func
func AskFollower(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var follower models.Follower
	follower.UserID = UserID
	follower.UserFollowed = ID

	//Setting the status for the followed user
	var response models.FollowPlusUser

	status, err := db.CheckFollower(follower)

	if err != nil || status == false {
		response.Status = false
	} else {
		response.Status = true
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
