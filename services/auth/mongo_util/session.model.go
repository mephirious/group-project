package mongo_util

import "time"

type SessionSchema struct {
	ID        string    `bson:"_id"`
	UserID    string    `bson:"userId"`
	UserAgent string    `bson:"user_agent"`
	ExpiresAt time.Time `bson:"expires_at"`
	CreatedAt time.Time `bson:"created_at"`
}
