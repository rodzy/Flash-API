package models

//Follower is the encharge of defining the follower
type Follower struct {
	UserID       string `bson:"user_id" json:"user_id,omitempty"`
	UserFollowed string `bson:"user_followed" json:"user_followed,omitempty"`
}