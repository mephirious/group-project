package mongo_util

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	CreateSessionInput struct {
		UserID    string
		UserAgent string
	}
	GetSessionsInput struct {
		Limit   int64
		Offset  int64
		OrderBy *string

		ID        *string
		UserID    *string
		UserAgent *string
		ExpiresAt *time.Time
		CreatedAt *time.Time
	}
)

func (c GetSessionsInput) buildFilter() bson.M {
	filter := bson.M{}

	if c.ID != nil {
		filter["_id"] = *c.ID
	}
	if c.UserID != nil {
		filter["userId"] = *c.UserID
	}
	if c.UserAgent != nil {
		filter["user_agent"] = *c.UserAgent
	}
	if c.ExpiresAt != nil {
		filter["expires_at"] = *c.ExpiresAt
	}
	if c.CreatedAt != nil {
		filter["created_at"] = *c.CreatedAt
	}

	return filter
}

func (db *DB) CreateSession(ctx context.Context, input CreateSessionInput) (*SessionSchema, error) {
	collection := db.DB.Collection("sessions")

	newSession := SessionSchema{
		ID:        primitive.NewObjectID().Hex(),
		UserID:    input.UserID,
		UserAgent: input.UserAgent,
		ExpiresAt: time.Now().AddDate(0, 0, 30),
		CreatedAt: time.Now(),
	}

	_, err := collection.InsertOne(ctx, newSession)
	if err != nil {
		return nil, err
	}

	return &newSession, nil
}

func (db *DB) GetSessionOne(ctx context.Context, input GetSessionsInput) (*SessionSchema, error) {
	collection := db.DB.Collection("sessions")

	filter := input.buildFilter()

	var session SessionSchema
	err := collection.FindOne(ctx, filter).Decode(&session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}
