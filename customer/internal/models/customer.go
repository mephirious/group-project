package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Customer - структура клиента в MongoDB
type Customer struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password,omitempty" json:"-"` // Храним хеш
}
