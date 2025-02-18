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

type BlogPostRepository interface {
	GetAllBlogPosts(ctx context.Context, limit, skip int, sortField, sortOrder string) ([]domain.BlogPost, error)
	GetBlogPostByID(ctx context.Context, id primitive.ObjectID) (*domain.BlogPost, error)
	GetBlogPostByTitle(ctx context.Context, title string) (*domain.BlogPost, error)
	CreateBlogPost(ctx context.Context, post *domain.BlogPost) error
	UpdateBlogPost(ctx context.Context, post *domain.BlogPost) error
	DeleteBlogPost(ctx context.Context, id primitive.ObjectID) error
}

type blogPostRepository struct {
	collection *mongo.Collection
}

func NewBlogPostRepository(db *mongo.Database) *blogPostRepository {
	return &blogPostRepository{
		collection: db.Collection("blog_posts"),
	}
}

func (r *blogPostRepository) GetAllBlogPosts(ctx context.Context, limit, skip int, sortField, sortOrder string) ([]domain.BlogPost, error) {
	var posts []domain.BlogPost

	sortValue := 1
	if sortOrder == "desc" {
		sortValue = -1
	}

	sortOptions := bson.D{{Key: sortField, Value: sortValue}}

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(skip))
	findOptions.SetSort(sortOptions)

	cursor, err := r.collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *blogPostRepository) GetBlogPostByID(ctx context.Context, id primitive.ObjectID) (*domain.BlogPost, error) {
	var post domain.BlogPost

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &post, nil
}

func (r *blogPostRepository) GetBlogPostByTitle(ctx context.Context, title string) (*domain.BlogPost, error) {
	regexPattern := bson.M{
		"$regex":   title,
		"$options": "i",
	}

	filter := bson.M{
		"title": regexPattern,
	}

	var post domain.BlogPost
	err := r.collection.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &post, nil
}

func (r *blogPostRepository) CreateBlogPost(ctx context.Context, post *domain.BlogPost) error {
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, post)
	if err != nil {
		return err
	}

	return nil
}

func (r *blogPostRepository) UpdateBlogPost(ctx context.Context, post *domain.BlogPost) error {
	post.UpdatedAt = time.Now()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": post.ID}, bson.M{"$set": post})
	if err != nil {
		return err
	}

	return nil
}

func (r *blogPostRepository) DeleteBlogPost(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
