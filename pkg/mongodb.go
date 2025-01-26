package pkg

import (
	"context"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	MongoDB struct {
		URI      string `yaml:"uri" env:"MONGODB_URI"`
		Database string `yaml:"database" env:"MONGODB_DATABASE"`
	} `yaml:"mongodb"`
}

type DB struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func LoadConfig() (*Config, error) {
	var config Config
	if err := cleanenv.ReadConfig("config.yaml", &config); err != nil {
		return nil, fmt.Errorf("unable to read config: %w", err)
	}
	return &config, nil
}

func NewDB(ctx context.Context, cfg Config) (*DB, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDB.URI))
	if err != nil {
		return nil, fmt.Errorf("mongo.Connect: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("mongo.Ping: %w", err)
	}

	db := client.Database(cfg.MongoDB.Database)

	return &DB{
		Client: client,
		DB:     db,
	}, nil
}

func (db *DB) Close(ctx context.Context) error {
	if err := db.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("mongo.Disconnect: %w", err)
	}
	return nil
}
