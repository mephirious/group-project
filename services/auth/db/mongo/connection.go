package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mephirious/group-project/services/auth/db/mongo/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDB(ctx context.Context) (*repository.DB, error) {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set in .env file")
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	databaseName := os.Getenv("MONGO_DB_NAME")
	if databaseName == "" {
		log.Fatal("MONGO_DB_NAME not set in .env file")
	}
	log.Println("connected to MongoDB")

	return &repository.DB{Client: client, DB: client.Database(databaseName)}, nil
}
