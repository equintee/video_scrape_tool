package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Content struct {
	Id          bson.ObjectID `bson:"_id"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
	Source      string        `bson:"source"`
	ContentUrl  string        `bson:"content_url"`
	Tags        []string      `bson:"tags"`
	Song        Song
}

type Song struct {
	Name   string `bson:"name"`
	Artist string `bson:"artist"`
}
