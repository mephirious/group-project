package mongo_util

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client *mongo.Client
	DB     *mongo.Database
}

type List[T any] struct {
	Total    int64
	Elements []T
}

func NewDB(ctx context.Context) (*DB, error) {
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

	return &DB{Client: client, DB: client.Database(databaseName)}, nil
}

func (db *DB) Close(ctx context.Context) error {
	if err := db.Client.Disconnect(ctx); err != nil {
		return err
	}
	return nil
}
