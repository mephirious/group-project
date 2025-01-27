package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/mephirious/group-project/services/authorization/mongo_util"
)

type Config struct {
	PORT   string
	Prefix string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file")
	}

	return &Config{
		PORT:   ":5001",
		Prefix: "/api/v1",
	}
}

func main() {
	log.Printf("STARTING AUTH SERVICE (%v environment) ...", os.Getenv("SERVICE_ENV"))

	cfg := NewConfig()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := mongo_util.NewDB(ctx)
	if err != nil {
		log.Fatal("Error creating DB")
	}
	defer func() {
		if err := db.Close(ctx); err != nil {
			log.Fatal("Error closing MongoDB connection: %v", err)
		}
	}()

	svc := NewAuthService(db)
	svc = NewLoggingService(logger, svc)

	ApiServer := NewApiServer(svc)
	log.Fatal(ApiServer.Start(cfg.PORT, cfg.Prefix))

	signalChan := make(chan os.Signal, 1)
	<-signalChan
	log.Println("Shutdown signal received")

}
