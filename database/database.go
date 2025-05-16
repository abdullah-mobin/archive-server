package database

import (
	"archive-server/config"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectMongoDB(config *config.Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.MongoDB.URL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	DB = client

	log.Println("Connected to MongoDB successfully")
	return client, nil

}
