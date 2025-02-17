package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mephirious/group-project/services/payment-service/config"
	"github.com/mephirious/group-project/services/payment-service/domain"
)

type Handler struct {
	MongoClient *mongo.Client
	Config      *config.Config
}

func NewHandler(mongoClient *mongo.Client, config *config.Config) *Handler {
	return &Handler{
		MongoClient: mongoClient,
		Config:      config,
	}
}

func (h *Handler) CreateCheckoutSession(c *gin.Context) {
	var products []domain.Product
	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		return
	}

	var totalAmount int64
	for _, product := range products {
		totalAmount += product.Price * product.Quantity
	}

	var lineItems []*stripe.CheckoutSessionLineItemParams
	for _, product := range products {
		lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String(product.Currency),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name: stripe.String(product.Name),
				},
				UnitAmount: stripe.Int64(product.Price),
			},
			Quantity: stripe.Int64(product.Quantity),
		})
	}

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems:          lineItems,
		Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:         stripe.String("http://localhost:3000/success"),
		CancelURL:          stripe.String("http://localhost:3000/cancel"),
	}

	s, err := session.New(params)
	if err != nil {
		log.Printf("Error creating checkout session: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create checkout session"})
		return
	}

	order := domain.Order{
		Products:  products,
		Amount:    totalAmount,
		Status:    "pending",
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	collection := h.MongoClient.Database(h.Config.Database.Name).Collection("orders")
	_, err = collection.InsertOne(context.Background(), order)
	if err != nil {
		log.Printf("Failed to save order: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"url": s.URL})
}

func (h *Handler) HandleWebhook(c *gin.Context) {
	var event stripe.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch event.Type {
	case "checkout.session.completed":
		var session stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &session)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse session"})
			return
		}

		collection := h.MongoClient.Database(h.Config.Database.Name).Collection("orders")
		filter := bson.M{"status": "pending"}
		update := bson.M{"$set": bson.M{"status": "paid"}}
		_, err = collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			log.Printf("Failed to update order status: %v", err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
