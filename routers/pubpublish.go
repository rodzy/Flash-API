package routers

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/rodzy/flash/db"
	"github.com/rodzy/flash/models"
)
//PublishPub router
func PublishPub(w http.ResponseWriter,r *http.Request)  {
	var content models.Pub
	err:=json.NewDecoder(r.Body).Decode(&content)

	reg:=models.InsertPub{
		//Pub user: Global user
		UserID: UserID,
		Content: content.Content,
		DateTime: time.Now(),
	}
	_,status,err:=db.PublishPub(reg)
	if err != nil {
		http.Error(w,"Register error"+err.Error(),400)
		return
	}
	if status==false {
		http.Error(w,"Bad status code",400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}