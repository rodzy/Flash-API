package models

import (
	"time"
)

//InsertPub struct the pub into the db
type InsertPub struct {
	UserID   string    `bson:"userid" json:"userid,omitempty"`
	Content  string    `bson:"content" json:"content,omitempty"`
	DateTime time.Time `bson:"datetime" json:"datetime,omitempty"`
}
