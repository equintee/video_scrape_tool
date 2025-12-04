package util

import (
	"reflect"

	"github.com/u2takey/go-utils/env"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func GenerateUpdate(old any, update any) bson.D {
	oldType := reflect.TypeOf(old)
	updateType := reflect.TypeOf(update)
	updates := bson.D{}
	for i := 0; i < updateType.NumField(); i++ {
		fieldName := updateType.Field(i).Name
		_, exists := oldType.FieldByName(fieldName)
		if !exists {
			continue
		}

		oldValue := reflect.ValueOf(old).FieldByName(fieldName)
		newValue := reflect.ValueOf(update).FieldByName(fieldName)

		if oldValue != newValue {
			updates = append(updates, bson.E{Key: fieldName, Value: newValue})
		}
	}
	return updates
}
