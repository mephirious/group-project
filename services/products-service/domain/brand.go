package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Brand struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BrandName string             `bson:"brand_name" json:"brand_name"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
