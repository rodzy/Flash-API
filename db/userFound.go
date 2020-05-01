package db

import (
	"context"
	"time"

	"github.com/rodzy/flash/models"
	"go.mongodb.org/mongo-driver/bson"
)

//UserFound finds for specific user througt email
func UserFound(email string) (models.User,bool,string) {
	ctx,cancel := context.WithTimeout(context.Background(),15*time.Second)
	defer cancel()

	db:=MongoCon.Database("flash")
	col:=db.Collection("user")

	quest := bson.M{"email":email}

	var result models.User

	err:=col.FindOne(ctx,quest).Decode(&result)
	ID:=result.ID.Hex()
	if err != nil {
		return result,false,ID
	}
	return result,true,ID
}

