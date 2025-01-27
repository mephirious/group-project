package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/mephirious/group-project/products-service/config"
	"github.com/mephirious/group-project/products-service/db"
	"github.com/mephirious/group-project/products-service/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("Error occured while loading config: %s", err))
		os.Exit(1)
	}

	client, err := db.ConnectToMongoDB(ctx, cfg.Database.URI)
	if err != nil {
		slog.Error(fmt.Sprintf("Error occured while loading config: %s", err))
		os.Exit(1)
	}
	defer db.DisconnectFromMongoDB(ctx, client)

	var product domain.Product

	db := client.Database("laptopStore")
	collection := db.Collection("products")

	err = collection.FindOne(ctx, bson.M{"_id": "1001"}).Decode(&product)
	if err != nil {
		slog.Error(fmt.Sprintf("Error occured while finding record in db: %s", err))
	}

	fmt.Println(product)
}
