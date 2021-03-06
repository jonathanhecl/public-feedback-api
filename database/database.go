package database

import (
	"context"
	"log"
	"net/url"
	"path"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataStore struct {
	googleCert       string
	groupSpreadsheet string
	client           *mongo.Client
	mongoDb          string
	messages         *mongo.Collection
	moderation       *mongo.Collection
	tracking         *mongo.Collection
	feedback         *mongo.Collection
}

func InitDatabase(mongoUri string, googleCert string, groupSpreadsheet string) *DataStore {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri).SetRetryWrites(false))
	if err != nil {
		log.Fatal("Database->InitDatabase: ", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Database->InitDatabase: ", err)
	}

	mongoDb, err := url.Parse(mongoUri)
	if err != nil {
		log.Fatal("Database->InitDatabase: ", err)
	}

	return &DataStore{
		googleCert:       googleCert,
		groupSpreadsheet: groupSpreadsheet,
		client:           client,
		mongoDb:          path.Base(mongoDb.Path),
		messages:         client.Database(path.Base(mongoDb.Path)).Collection("messages"),
		moderation:       client.Database(path.Base(mongoDb.Path)).Collection("moderation"),
		tracking:         client.Database(path.Base(mongoDb.Path)).Collection("tracking"),
		feedback:         client.Database(path.Base(mongoDb.Path)).Collection("feedback"),
	}
}

func CloseDatabase(db *DataStore) {
	err := db.client.Disconnect(context.Background())
	if err != nil {
		log.Fatal("Database->CloseDatabase: ", err)
	}
}
