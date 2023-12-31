package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Database, error) {
	clientOptions := options.Client()
	hostURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	clientOptions = clientOptions.ApplyURI(hostURI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	db := client.Database(dbName)
	_, err = db.Collection("users").Indexes().CreateMany(context.Background(), []mongo.IndexModel{{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}, {

		Keys:    bson.D{{Key: "email", Value: 2}},
		Options: options.Index().SetUnique(true),
	}})
	if err != nil {
		return nil, err
	}
	log.Println("Connected to database.")
	return db, nil
}
