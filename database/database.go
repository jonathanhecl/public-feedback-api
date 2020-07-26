package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataStore struct {
	client   *mongo.Client
	mongoDb  string
	messages *mongo.Collection
	groups   *mongo.Collection
}

func InitDatabase(mongoUri string, mongoDb string) *DataStore {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri).SetRetryWrites(false))
	if err != nil {
		log.Fatal("Database->InitDatabase", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Database->InitDatabase", err)
	}

	return &DataStore{
		client:   client,
		mongoDb:  mongoDb,
		messages: client.Database(mongoDb).Collection("messages"),
		groups:   client.Database(mongoDb).Collection("groups"),
	}
}

func CloseDatabase(db *DataStore) {
	err := db.client.Disconnect(context.Background())
	if err != nil {
		log.Fatal("Database->CloseDatabase", err)
	}
}
