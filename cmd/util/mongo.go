package util

import (
	"github.com/u2takey/go-utils/env"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type mongoConnection struct {
	connection *mongo.Client
}

var database mongoConnection

func init() {
	connectionUri := env.GetEnvAsStringOrFallback("MONGO_CONNECTION_URI", "")
	if connectionUri == "" {
		panic("MONGO_CONNECTION_URI not defined in env.")
	}
	opts := options.Client()
	opts.ApplyURI(connectionUri)

	client, err := mongo.Connect(opts)

	if err != nil {
		panic(err)
	}

	database = mongoConnection{
		connection: client,
	}
}

func GetDatabase(document string) *mongo.Collection {
	return database.connection.Database("content-service").Collection(document)
}
