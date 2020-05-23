package models

//Pub content for the selected Pub
type Pub struct {
	Content string `bson:"content" json:"content,omitempty"`
}