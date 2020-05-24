package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ViewPub is the struct to read the pub requested
type ViewPub struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID   string             `bson:"userid" json:"userid,omitempty"`
	Content  string             `bson:"content" json:"content,omitempty"`
	DateTime time.Time          `bson:"datetime" json:"datetime,omitempty"`
}
