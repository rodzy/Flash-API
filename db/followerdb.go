package db

import (
	"context"
	"time"

	"github.com/rodzy/flash/models"
	"go.mongodb.org/mongo-driver/bson"
)

//InsertFollower func charges the new follower to the database
func InsertFollower(f models.Follower) (bool, error) {
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
func UnfollowUser(f models.Follower) (bool, error) {
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

//CheckFollower func to see specific follower information like the profile
func CheckFollower(f models.Follower) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCon.Database("flash")
	col := db.Collection("follower")

	reg := bson.M{
		"user_id":       f.UserID,
		"user_followed": f.UserFollowed,
	}

	var result models.Follower
	err := col.FindOne(ctx, reg).Decode(&result)
	if err != nil {
		return false, err
	}
	return true, nil
}
