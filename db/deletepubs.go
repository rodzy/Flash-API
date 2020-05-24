package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DeletePub deletes the registry from the pub collection
func DeletePub(ID string,UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCon.Database("flash")
	col := db.Collection("pubs")

	objID,_:=primitive.ObjectIDFromHex(ID)

	reg := bson.M{
		"_id":   objID,
		"userid":  UserID,
	}

	_,err:=col.DeleteOne(ctx,reg)
	return err
}