package repository

import (
	"context"
	"time"

	"github.com/mephirious/group-project/services/products-service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReviewRepository interface {
	GetAllReviews(ctx context.Context, limit, skip int, sortField string, sortOrder string, verified *bool) ([]domain.Review, error)
	GetReviewByID(ctx context.Context, id primitive.ObjectID) (*domain.Review, error)
	GetReviewsByCustomerID(ctx context.Context, customerID primitive.ObjectID, limit int, verified *bool) ([]domain.Review, error)
	GetReviewsByProductID(ctx context.Context, productID primitive.ObjectID, limit int, verified *bool) ([]domain.Review, error)
	UpdateReview(ctx context.Context, id primitive.ObjectID, review *domain.Review) error
	DeleteReview(ctx context.Context, id primitive.ObjectID) error
	CreateReview(ctx context.Context, review *domain.Review) error
	GetReviewsByCustomerAndProductIDs(ctx context.Context, productID primitive.ObjectID, customerID primitive.ObjectID) ([]domain.Review, error)
	GetReviewStatsByProductID(ctx context.Context, productID primitive.ObjectID, verified *bool) (float64, error)
	GetAllReviewsCount(ctx context.Context, verified *bool) (int64, error)
}

type reviewRepository struct {
	collection *mongo.Collection
}

func NewReviewRepository(db *mongo.Database) *reviewRepository {
	return &reviewRepository{
		collection: db.Collection("reviews"),
	}
}

func (r *reviewRepository) CreateReview(ctx context.Context, review *domain.Review) error {
	review.CreatedAt = time.Now()
	review.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, review)
	if err != nil {
		return err
	}

	return nil
}

func (r *reviewRepository) GetAllReviews(ctx context.Context, limit, skip int, sortField string, sortOrder string, verified *bool) ([]domain.Review, error) {
	var reviews []domain.Review

	filter := bson.M{}
	if verified != nil {
		filter["verified"] = *verified
	}

	sortValue := 1
	if sortOrder == "desc" {
		sortValue = -1
	}
	sortOptions := bson.D{{Key: sortField, Value: sortValue}}

	opts := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(skip)).
		SetSort(sortOptions)

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *reviewRepository) GetReviewByID(ctx context.Context, id primitive.ObjectID) (*domain.Review, error) {
	var review domain.Review

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&review)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &review, nil
}

func (r *reviewRepository) GetReviewsByCustomerID(ctx context.Context, customerID primitive.ObjectID, limit int, verified *bool) ([]domain.Review, error) {
	filter := bson.M{"customer_id": customerID}
	if verified != nil {
		filter["verified"] = *verified
	}

	opts := options.Find().SetLimit(int64(limit))

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reviews []domain.Review
	if err := cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *reviewRepository) GetReviewsByCustomerAndProductIDs(ctx context.Context, productID primitive.ObjectID, customerID primitive.ObjectID) ([]domain.Review, error) {
	filter := bson.M{"product_id": productID, "customer_id": customerID}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reviews []domain.Review
	if err := cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *reviewRepository) GetReviewsByProductID(ctx context.Context, productID primitive.ObjectID, limit int, verified *bool) ([]domain.Review, error) {
	var reviews []domain.Review

	filter := bson.M{"product_id": productID}
	if verified != nil {
		filter["verified"] = *verified
	}

	opts := options.Find().
		SetLimit(int64(limit))

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *reviewRepository) GetReviewStatsByProductID(ctx context.Context, productID primitive.ObjectID, verified *bool) (float64, error) {
	filter := bson.M{"product_id": productID}
	if verified != nil {
		filter["verified"] = *verified
	}

	pipeline := bson.A{
		bson.M{"$match": filter},
		bson.M{"$group": bson.M{
			"_id":           nil,
			"averageRating": bson.M{"$avg": "$rating"},
		}},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	var result struct {
		AverageRating float64 `bson:"averageRating"`
	}
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return 0, err
		}
	}

	return result.AverageRating, nil
}

func (r *reviewRepository) GetAllReviewsCount(ctx context.Context, verified *bool) (int64, error) {
	filter := bson.M{}
	if verified != nil {
		filter["verified"] = *verified
	}
	return r.collection.CountDocuments(ctx, filter)
}

func (r *reviewRepository) UpdateReview(ctx context.Context, id primitive.ObjectID, review *domain.Review) error {
	review.UpdatedAt = time.Now()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": review})
	if err != nil {
		return err
	}

	return nil
}

func (r *reviewRepository) DeleteReview(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
