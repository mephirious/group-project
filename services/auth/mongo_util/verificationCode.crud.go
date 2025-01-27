package mongo_util

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	CreateVerificationCodeInput struct {
		UserID    string
		Type      string
		ExpiresAt time.Time
	}
)

const EmailVerification string = "email_verification"

func (db *DB) CreateVerificationCode(ctx context.Context, input CreateVerificationCodeInput) (VerificationCodeSchema, error) {
	collection := db.DB.Collection("verification_codes")

	newCode := VerificationCodeSchema{
		ID:        primitive.NewObjectID().Hex(),
		UserID:    input.UserID,
		Type:      input.Type,
		ExpiresAt: input.ExpiresAt,
		CreatedAt: time.Now(),
	}

	_, err := collection.InsertOne(ctx, newCode)
	if err != nil {
		return VerificationCodeSchema{}, err
	}

	return newCode, nil
}
