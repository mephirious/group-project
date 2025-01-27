package mongo_util

import "time"

type CustomerView struct {
	ID       string `bson:"_id"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Verified bool   `bson:"verified"`
	// FirstName string    `bson:"first_name,omitempty"`
	// LastName  string    `bson:"last_name,omitempty"`
	// Phone     string    `bson:"phone,omitempty"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
}

type CustomerSchema struct {
	ID        string    `bson:"_id"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	FirstName string    `bson:"first_name,omitempty"`
	LastName  string    `bson:"last_name,omitempty"`
	Phone     string    `bson:"phone,omitempty"`
	Verified  bool      `bson:"verified"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
}
