package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCon ...
var MongoCon = Connection()

//Setting the client URI
var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")

//Connection func setting the connection values for mongo atlas cloud
func Connection() *mongo.Client {
	//getting the contextclient from DB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	//Checking the connection pinging
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}
	log.Println("Success!")
	return client
}

//CheckConnection func for checking distant connection status
func CheckConnection() int {
	//This method will allow to check the even ping for the Mongo connection established
	err := MongoCon.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
