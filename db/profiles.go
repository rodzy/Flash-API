package db

import (
	"context"
	"fmt"
	"time"

	"github.com/rodzy/flash/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FindProfile finds the user by ID, then it returns the found user
func FindProfile(ID string)(models.User,error){
	con,cancel:=context.WithTimeout(context.Background(),time.Second*15)
	defer cancel()
	db:=MongoCon.Database("flash")
	col:=db.Collection("user")

	var profile models.User
	_id,_:=primitive.ObjectIDFromHex(ID)

	//Search for the ID
	search:=bson.M{
		"_id":_id,
	}
	err:=col.FindOne(con,search).Decode(&profile)
	//Setting the password blank we don't need it
	profile.Password=""
	if err != nil {
		fmt.Println("User not found"+err.Error())
		return profile,err
	}
	return profile,nil
}