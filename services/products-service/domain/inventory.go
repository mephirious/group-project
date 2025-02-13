package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Inventory struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProductID primitive.ObjectID `bson:"product_id" json:"product_id"`
	Quantity  uint32             `bson:"quantity" json:"quantity"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
