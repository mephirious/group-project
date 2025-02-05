package mongo_util

import "time"

type VerificationCodeSchema struct {
	ID        string    `bson:"_id"`
	UserID    string    `bson:"userId"`
	Type      string    `bson:"type"`
	ExpiresAt time.Time `bson:"expires_at"`
	CreatedAt time.Time `bson:"created_at"`
}
