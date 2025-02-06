package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func (db *DB) Close(ctx context.Context) error {
	if err := db.Client.Disconnect(ctx); err != nil {
		return err
	}
	return nil
}
