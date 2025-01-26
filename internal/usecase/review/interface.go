package usecase

import (
	"context"

	model "github.com/mephirious/group-project/internal/model"
)

// ReviewRepo defines the methods for interacting with review data.
type ReviewRepo interface {
	AddReviewForProduct(ctx context.Context, productID string, review model.Review) (*model.Review, error)
	GetReviewsForProductByProductID(ctx context.Context, productID string) ([]model.Review, error)
}
