package db

import (
	"context"
	"time"

	"github.com/rodzy/flash/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//PublishPub is the func that manages the Insertion of the PUB
func PublishPub(p models.InsertPub) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCon.Database("flash")
	col := db.Collection("pubs")

	reg := bson.M{
		"userid":   p.UserID,
		"content":  p.Content,
		"datetime": p.DateTime,
	}
	result, err := col.InsertOne(ctx, reg)
	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
