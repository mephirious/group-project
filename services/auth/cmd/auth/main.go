package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	h "github.com/mephirious/group-project/services/auth/api/handler"
	m "github.com/mephirious/group-project/services/auth/db/mongo"
	s "github.com/mephirious/group-project/services/auth/service"
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
		Prefix: "/auth/api/v1",
	}
}

func main() {
	log.Printf("STARTING AUTH SERVICE (%v environment) ...", os.Getenv("SERVICE_ENV"))

	cfg := NewConfig()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := m.NewDB(ctx)
	if err != nil {
		log.Fatal("Error creating DB")
	}
	defer func() {
		if err := db.Close(ctx); err != nil {
			log.Fatal("Error closing MongoDB connection: %v", err)
		}
	}()

	svc := s.NewAuthService(db)
	svc = s.NewLoggingService(logger, svc)

	ApiServer := h.NewApiServer(svc)
	log.Fatal(ApiServer.Start(cfg.PORT, cfg.Prefix))

	signalChan := make(chan os.Signal, 1)
	<-signalChan
	log.Println("Shutdown signal received")

}
