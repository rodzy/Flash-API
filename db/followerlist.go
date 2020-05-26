package db

import (
	"context"
	"time"

	"github.com/rodzy/flash/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)
//GetFollowerList func
func GetFollowerList(ID string,page int64,search string,type string)([]*models.User,bool){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCon.Database("flash")
	col := db.Collection("user")

	var results []*models.User

	findOptions:=options.Find()
	findOptions.SetSkip((page -1)*20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}
	
	cursor,err:=col.Find(ctx,query,findOptions)
	if err != nil {
		return results,false
	}

	var found,add bool

	for cursor.Next(ctx){
		var u models.User
		err:=cursor.Decode(&u)
		if err != nil {
			return results,false
		}
		var rel models.Follower
		rel.UserID=ID
		rel.UserFollowed=u.ID.Hex()
		
		add=false

		found,err=CheckFollower(rel)
		if type=="new" && found==true {
			add=true		
		}
		if type=="follow" && found==false {
			add=true
		}
		if rel.UserFollowed==ID{
			add=false
		}
		//Setting blank the properties that we don't need to show in the actual view
		if add==true {
			u.Password=""
			u.Bio=""
			u.Email=""
			u.WebSite=""
			u.Banner=""
			u.Location=""
			results=append(results,&u)
		}
	}
	err=cursor.Err()
	if err != nil {
		return results,false
	}
	cursor.Close(ctx)
	return results,true
	
}
