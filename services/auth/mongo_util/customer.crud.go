package mongo_util

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	CreateCustomerInput struct {
		Email     string
		Password  string
		FirstName *string
		LastName  *string
		Phone     *string
	}
	GetCustomersInput struct {
		Limit   int64
		Offset  int64
		OrderBy *string

		ID        *string
		FirstName *string
		LastName  *string
		Email     *string
		Phone     *string
		CreatedAt *time.Time
		UpdatedAt *time.Time
		Verified  *bool
	}
)

func (c GetCustomersInput) buildFilter() bson.M {
	filter := bson.M{}

	if c.ID != nil {
		filter["_id"] = *c.ID
	}
	if c.FirstName != nil {
		filter["first_name"] = *c.FirstName
	}
	if c.LastName != nil {
		filter["last_name"] = *c.LastName
	}
	if c.Email != nil {
		filter["email"] = *c.Email
	}
	if c.Phone != nil {
		filter["phone"] = *c.Phone
	}
	if c.CreatedAt != nil {
		filter["created_at"] = *c.CreatedAt
	}
	if c.UpdatedAt != nil {
		filter["updated_at"] = *c.UpdatedAt
	}
	if c.Verified != nil {
		filter["verified"] = *c.Verified
	}

	return filter
}

func (db *DB) CreateCustomer(ctx context.Context, input CreateCustomerInput) (*CustomerSchema, error) {
	// Generate a unique ID for the new customer
	newCustomer := &CustomerSchema{
		ID:        primitive.NewObjectID().Hex(),
		Email:     input.Email,
		Password:  input.Password,
		Verified:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Set optional fields if provided
	if input.Phone != nil && *input.Phone != "" {
		newCustomer.Phone = *input.Phone
	}
	if input.FirstName != nil && *input.FirstName != "" {
		newCustomer.FirstName = *input.FirstName
	}
	if input.LastName != nil && *input.LastName != "" {
		newCustomer.LastName = *input.LastName
	}

	// Insert the new customer into the "customers" collection
	collection := db.DB.Collection("customers")
	_, err := collection.InsertOne(ctx, newCustomer)
	if err != nil {
		return nil, fmt.Errorf("failed to insert new customer: %w", err)
	}

	return newCustomer, nil
}

func (db *DB) GetCustomersOne(ctx context.Context, input GetCustomersInput) (*CustomerSchema, error) {
	collection := db.DB.Collection("customers")

	filter := input.buildFilter()

	var customer CustomerSchema
	err := collection.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (db *DB) GetCustomersMany(ctx context.Context, input GetCustomersInput) (*List[CustomerSchema], error) {
	collection := db.DB.Collection("customers")

	filter := input.buildFilter()

	var customers List[CustomerSchema]
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var customer CustomerSchema
		if err := cursor.Decode(&customer); err != nil {
			return nil, err
		}
		customers.Elements = append(customers.Elements, customer)
		customers.Total++
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &customers, nil
}
