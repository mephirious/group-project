package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mephirious/group-project/services/customer/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongoDB(cfg *config.Config) {
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Faield to connect to MongoDB: %v", err)
	}

	MongoClient = client
	fmt.Println("Connected to MongoDB!")
}

func GetCollection(dbName, collectionName string) *mongo.Collection {
	return MongoClient.Database(dbName).Collection(collectionName)
}
