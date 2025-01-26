package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ProductID  string             `bson:"product_id"`
	CustomerID string             `bson:"customer_id"`
	Rating     float64            `bson:"rating"`
	ReviewDate time.Time          `bson:"review_date"`
	Comment    string             `bson:"comment"`
}
