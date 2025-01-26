package usecase

import (
	"context"

	model "github.com/mephirious/group-project/internal/model"
)

// Review defines the structure for handling review-related operations.
type Review struct {
	reviewRepo ReviewRepo
}

// NewReview creates a new instance of the Review use case.
func NewReview(reviewRepo ReviewRepo) *Review {
	return &Review{
		reviewRepo: reviewRepo,
	}
}

// Add adds a new review for a product.
func (uc *Review) Add(ctx context.Context, productID string, review model.Review) (*model.Review, error) {
	_, err := uc.reviewRepo.AddReviewForProduct(ctx, productID, review)
	if err != nil {
		return nil, err
	}
	return &review, nil
}

// GetAll retrieves all reviews for a specific product.
func (uc *Review) GetAll(ctx context.Context, productID string) ([]model.Review, error) {
	return uc.reviewRepo.GetReviewsForProductByProductID(ctx, productID)
}
