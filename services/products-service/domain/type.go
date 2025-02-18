package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Type struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TypeName  string             `bson:"type_name"  json:"type_name"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
