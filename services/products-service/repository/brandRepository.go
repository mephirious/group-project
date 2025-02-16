package repository

import (
	"context"
	"time"

	"github.com/mephirious/group-project/services/products-service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BrandRepository interface {
	GetAllBrands(ctx context.Context) ([]domain.Brand, error)
	GetBrandByID(ctx context.Context, id primitive.ObjectID) (*domain.Brand, error)
	GetBrandByName(ctx context.Context, name string) (*domain.Brand, error)
	CreateBrand(ctx context.Context, brand *domain.Brand) error
	UpdateBrand(ctx context.Context, brand *domain.Brand) error
	DeleteBrand(ctx context.Context, id primitive.ObjectID) error
}

type brandRepository struct {
	collection *mongo.Collection
}

func NewBrandRepository(db *mongo.Database) *brandRepository {
	return &brandRepository{
		collection: db.Collection("brands"),
	}
}

func (b *brandRepository) GetAllBrands(ctx context.Context) ([]domain.Brand, error) {
	var brands []domain.Brand

	cursor, err := b.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &brands)
	if err != nil {
		return nil, err
	}

	return brands, nil
}

func (b *brandRepository) GetBrandByID(ctx context.Context, id primitive.ObjectID) (*domain.Brand, error) {
	var brand domain.Brand

	err := b.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&brand)
	if err != nil {
		return nil, err
	}

	return &brand, nil
}

func (b *brandRepository) GetBrandByName(ctx context.Context, name string) (*domain.Brand, error) {
	regexPattern := bson.M{
		"$regex":   name,
		"$options": "i",
	}

	filter := bson.M{
		"brand_name": regexPattern,
	}

	var brand domain.Brand
	err := b.collection.FindOne(ctx, filter).Decode(&brand)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &brand, nil
}

func (b *brandRepository) CreateBrand(ctx context.Context, brand *domain.Brand) error {
	brand.CreatedAt = time.Now()
	brand.UpdatedAt = time.Now()

	_, err := b.collection.InsertOne(ctx, brand)
	if err != nil {
		return err
	}

	return nil
}

func (b *brandRepository) UpdateBrand(ctx context.Context, brand *domain.Brand) error {
	brand.UpdatedAt = time.Now()

	_, err := b.collection.UpdateOne(ctx, bson.M{"_id": brand.ID}, bson.M{"$set": brand})
	if err != nil {
		return err
	}

	return nil
}

func (b *brandRepository) DeleteBrand(ctx context.Context, id primitive.ObjectID) error {
	_, err := b.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
