package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Brand struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	BrandName string             `bson:"brand_name"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
