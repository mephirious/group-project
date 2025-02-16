package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CustomerID primitive.ObjectID `bson:"customer_id" json:"customer_id"`
	ProductID  primitive.ObjectID `bson:"product_id" json:"product_id"`
	Content    string             `bson:"content" json:"content"`
	Rating     float64            `bson:"rating" json:"rating"`
	Verified   bool               `bson:"verified" json:"verified"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}
