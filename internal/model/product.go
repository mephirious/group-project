package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name"`
	Brand          string             `bson:"brand"`
	CategoryID     string             `bson:"category_id"`
	Price          float64            `bson:"price"`
	Stock          int32              `bson:"stock"`
	Specifications map[string]string  `bson:"specifications"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
}
