package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/mephirious/group-project/services/products-service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository interface {
	GetAllCategories(ctx context.Context) ([]domain.Category, error)
	GetCategoryByID(ctx context.Context, id primitive.ObjectID) (*domain.Category, error)
	GetCategoryByName(ctx context.Context, name string) (*domain.Category, error)
	CreateCategory(ctx context.Context, category *domain.Category) error
	UpdateCategory(ctx context.Context, category *domain.Category) error
	DeleteCategory(ctx context.Context, id primitive.ObjectID) error
}

type categoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository(db *mongo.Database) *categoryRepository {
	return &categoryRepository{
		collection: db.Collection("categories"),
	}
}

func (c *categoryRepository) GetAllCategories(ctx context.Context) ([]domain.Category, error) {
	var categories []domain.Category

	cursor, err := c.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *categoryRepository) GetCategoryByID(ctx context.Context, id primitive.ObjectID) (*domain.Category, error) {
	var category domain.Category

	err := c.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&category)
	if err != nil {
		return nil, err
	}

	fmt.Println(category)
	return &category, nil
}

func (c *categoryRepository) GetCategoryByName(ctx context.Context, name string) (*domain.Category, error) {
	regexPattern := bson.M{
		"$regex":   name,
		"$options": "i",
	}

	filter := bson.M{
		"category_name": regexPattern,
	}

	var category domain.Category
	err := c.collection.FindOne(ctx, filter).Decode(&category)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &category, nil
}

func (c *categoryRepository) CreateCategory(ctx context.Context, category *domain.Category) error {
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	_, err := c.collection.InsertOne(ctx, category)
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryRepository) UpdateCategory(ctx context.Context, category *domain.Category) error {
	category.UpdatedAt = time.Now()

	_, err := c.collection.UpdateOne(ctx, bson.M{"_id": category.ID}, bson.M{"$set": category})
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryRepository) DeleteCategory(ctx context.Context, id primitive.ObjectID) error {
	_, err := c.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
