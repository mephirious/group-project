package http

import (
	"context"

	model "github.com/mephirious/group-project/internal/model"
)

// ProductUsecase defines the business logic for products.
type ProductUsecase interface {
	Create(ctx context.Context, product model.Product) (*model.Product, error)
	GetAll(ctx context.Context) ([]model.Product, error)
	Get(ctx context.Context, id string) (*model.Product, error)
	Update(ctx context.Context, id string, product model.Product) (*model.Product, error)
	Delete(ctx context.Context, id string) (*model.Product, error)
}

// CategoryUsecase defines the business logic for categories.
type CategoryUsecase interface {
	Create(ctx context.Context, category model.Category) (*model.Category, error)
	GetAll(ctx context.Context) ([]model.Category, error)
}

// ReviewUsecase defines the business logic for reviews.
type ReviewUsecase interface {
	Add(ctx context.Context, productID string, review model.Review) (*model.Review, error)
	GetAll(ctx context.Context, productID string) ([]model.Review, error)
}
