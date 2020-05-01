package db

import (
	"context"
	"time"

	"github.com/rodzy/flash/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertUser inserting users
func InsertUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCon.Database("flash")
	col := db.Collection("user")
	u.Password, _ = EncryptPassword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	//Obtain the result ID value
	oID, _ := result.InsertedID.(primitive.ObjectID)
	return oID.String(), true, nil
}
