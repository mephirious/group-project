package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CategoryName string             `bson:"category_name"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}
