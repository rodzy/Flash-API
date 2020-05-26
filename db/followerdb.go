package db

import (
	"context"
	"time"

	"github.com/rodzy/flash/models"
)

//InsertFollower func charges the new follower to the database
func InsertFollower(f models.Follower)(bool,error)  {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCon.Database("flash")
	col := db.Collection("follower")
	
	_, err := col.InsertOne(ctx, f)
	if err != nil {
		return false, err
	}
	return true, nil
}

//UnfollowUser func to unfollow user as disguised
func UnfollowUser(f models.Follower)(bool,error)  {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCon.Database("flash")
	col := db.Collection("follower")
	
	_, err := col.DeleteOne(ctx, f)
	if err != nil {
		return false, err
	}
	return true, nil
}