package http

import (
	"context"

	model "github.com/mephirious/group-project/internal/model"
)

// Product defines the interface for handling products.
type Product interface {
	Create(ctx context.Context, product model.Product) (model.Product, error)
	GetAll(ctx context.Context) ([]model.Product, error)
	Get(ctx context.Context, id string) (model.Product, error)
	Update(ctx context.Context, id string, product model.Product) (model.Product, error)
	Delete(ctx context.Context, id string) (model.Product, error)
}

// Review defines the interface for handling reviews.
type Review interface {
	Add(ctx context.Context, productID string, review model.Review) (model.Review, error)
	GetAll(ctx context.Context, productID string) ([]model.Review, error)
}

// Category defines the interface for handling categories.
type Category interface {
	Create(ctx context.Context, category model.Category) (model.Category, error)
	GetAll(ctx context.Context) ([]model.Category, error)
}
