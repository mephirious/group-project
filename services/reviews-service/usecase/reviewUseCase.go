package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/mephirious/group-project/services/products-service/domain"
	"github.com/mephirious/group-project/services/products-service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReviewUseCase interface {
	GetAllReviews(ctx context.Context, limit, skip int, sortField string, sortOrder string, verified *bool) ([]domain.Review, error)
	GetReviewByID(ctx context.Context, id primitive.ObjectID) (*domain.Review, error)
	GetReviewsByCustomerID(ctx context.Context, customerID primitive.ObjectID, limit int, verified *bool) ([]domain.Review, error)
	GetReviewsByProductID(ctx context.Context, productID primitive.ObjectID, limit int, verified *bool) ([]domain.Review, int64, float64, error)
	UpdateReview(ctx context.Context, id primitive.ObjectID, review *domain.Review) error
	DeleteReview(ctx context.Context, id primitive.ObjectID) error
	CreateReview(ctx context.Context, review *domain.Review) error
}

type reviewUseCase struct {
	reviewRepository repository.ReviewRepository
}

func NewReviewUseCase(repository repository.ReviewRepository) *reviewUseCase {
	return &reviewUseCase{
		reviewRepository: repository,
	}
}

func (u *reviewUseCase) GetAllReviews(ctx context.Context, limit, skip int, sortField, sortOrder string, verified *bool) ([]domain.Review, error) {
	return u.reviewRepository.GetAllReviews(ctx, limit, skip, sortField, sortOrder, verified)
}
func (u *reviewUseCase) GetReviewByID(ctx context.Context, id primitive.ObjectID) (*domain.Review, error) {
	review, err := u.reviewRepository.GetReviewByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if review == nil {
		return nil, errors.New("review not found")
	}
	return review, nil
}

func (u *reviewUseCase) GetReviewsByCustomerID(ctx context.Context, customerID primitive.ObjectID, limit int, verified *bool) ([]domain.Review, error) {
	return u.reviewRepository.GetReviewsByCustomerID(ctx, customerID, limit, verified)
}

func (u *reviewUseCase) GetReviewsByProductID(ctx context.Context, productID primitive.ObjectID, limit int, verified *bool) ([]domain.Review, int64, float64, error) {
	reviews, err := u.reviewRepository.GetReviewsByProductID(ctx, productID, limit, verified)
	if err != nil {
		return nil, 0, 0, err
	}

	averageRating, ok := ratingsData[productID]
	if !ok {
		averageRating, err = u.reviewRepository.GetReviewStatsByProductID(ctx, productID, verified)
		if err != nil {
			return nil, 0, 0, err
		}
		ratingsData[productID] = averageRating
		fmt.Printf("Average rating updated: %v: %v/5\n", productID, averageRating)
	}

	totalReviews, err := u.reviewRepository.GetAllReviewsCount(ctx, verified)
	if err != nil {
		return nil, 0, 0, err
	}

	return reviews, totalReviews, averageRating, nil
}

func (u *reviewUseCase) UpdateReview(ctx context.Context, id primitive.ObjectID, review *domain.Review) error {
	existingReview, err := u.reviewRepository.GetReviewByID(ctx, id)
	if err != nil {
		return err
	}
	if existingReview != nil {
		return errors.New("product not found")
	}

	return u.reviewRepository.UpdateReview(ctx, id, review)
}

func (u *reviewUseCase) DeleteReview(ctx context.Context, id primitive.ObjectID) error {
	return u.reviewRepository.DeleteReview(ctx, id)
}

func (u *reviewUseCase) CreateReview(ctx context.Context, review *domain.Review) error {
	existingReview, err := u.reviewRepository.GetReviewsByCustomerAndProductIDs(ctx, review.ID, review.CustomerID)
	if err != nil {
		return err
	}
	if existingReview != nil {
		return errors.New("product already exists")
	}

	return u.reviewRepository.CreateReview(ctx, review)
}
