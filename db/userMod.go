package db

import (
	"context"
	"time"

	"github.com/rodzy/flash/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)
//ModifyUser func to modify the users information
func ModifyUser(user models.User,ID string)(bool,error)  {
	con,cancel:=context.WithTimeout(context.Background(),15*time.Second)
	defer cancel()

	db:=MongoCon.Database("flash")
	col:=db.Collection("user")

	//Mapping our found registry
	reg:=make(map[string]interface{})
	//Filling the map with the database information a clunky way to do it
	if len(user.Name)>0 {
		reg["name"]=user.Name
	}
	if len(user.LastName)>0 {
		reg["lastname"]=user.LastName
	}
	reg["birthdate"]=user.BirthDate
	if len(user.Avatar)>0 {
		reg["avatar"]=user.Avatar
	}
	if len(user.Banner)>0 {
		reg["banner"]=user.Banner
	}
	if len(user.Bio)>0 {
		reg["bio"]=user.Bio
	}
	if len(user.Location)>0 {
		reg["location"]=user.Location
	}
	if len(user.WebSite)>0 {
		reg["website"]=user.WebSite
	}
	update:=bson.M{
		"$set":reg,
	}

	_id,_:=primitive.ObjectIDFromHex(ID)
	filter:=bson.M{"_id":bson.M{"$eq":_id}}
	_,err:=col.UpdateOne(con,filter,update)
	if err != nil {
		return false,err
	}
	return true,nil
}