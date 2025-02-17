package main

import (
	"context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mephirious/group-project/services/payment-service/adapter/mongo"
	"github.com/mephirious/group-project/services/payment-service/api/http/handler"
	"github.com/mephirious/group-project/services/payment-service/config"
	"github.com/stripe/stripe-go/v76"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	stripe.Key = cfg.StripeSecretKey

	ctx := context.Background()
	mongoClient, err := mongo.ConnectToMongoDB(ctx, cfg.Database.URI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongo.DisconnectFromMongoDB(ctx, mongoClient)

	r := gin.Default()

	h := handler.NewHandler(mongoClient, cfg)

	r.POST("/create-checkout-session", h.CreateCheckoutSession)
	r.POST("/webhook", h.HandleWebhook)

	serverAddr := ":" + strconv.Itoa(cfg.Server.Port)
	log.Printf("Backend running on port %s...\n", serverAddr)
	r.Run(serverAddr)
}
