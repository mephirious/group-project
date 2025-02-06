package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Inventory struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProductID primitive.ObjectID `bson:"product_id"`
	Quantity  int64              `bson:"quantity"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
