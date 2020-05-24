package db

import (
	"context"
	"log"
	"time"

	"github.com/rodzy/flash/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ReadPubs func that handles the Pubs requested
Returns an slice to get all the pubs from a specific user*/
func ReadPubs(ID string, page int64) ([]*models.ViewPub, bool) {
	con, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCon.Database("flash")
	col := db.Collection("pubs")

	var result []*models.ViewPub

	//Finding the registered user
	reg := bson.M{
		"userid": ID,
	}
	//Set the pubs matching
	op := options.Find()
	op.SetLimit(20)
	op.SetSort(bson.D{{Key: "datetime", Value: -1}})
	op.SetSkip((page - 1) * 20)

	cursor, err := col.Find(con, reg, op)

	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	/*Looping trought the cursor if this finds data related with the models.ViewPub
	appends it to the original result slice*/
	for cursor.Next(con) {
		var registered models.ViewPub
		err := cursor.Decode(&registered)
		if err != nil {
			return result, false
		}
		result = append(result, &registered)
	}
	return result, true
}
