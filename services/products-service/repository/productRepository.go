package repository

import (
	"context"
	"regexp"
	"time"

	"github.com/mephirious/group-project/services/products-service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository interface {
	GetAllProducts(ctx context.Context, limit, skip int, sortField, sortOrder, search string) ([]domain.Product, error)
	GetProductByID(ctx context.Context, id primitive.ObjectID) (*domain.Product, error)
	GetProductByName(ctx context.Context, name string) (*domain.Product, error)
	CreateProduct(ctx context.Context, product *domain.Product) error
	UpdateProduct(ctx context.Context, product *domain.Product) error
	DeleteProduct(ctx context.Context, id primitive.ObjectID) error
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *productRepository {
	return &productRepository{
		collection: db.Collection("products"),
	}
}

func (p *productRepository) GetAllProducts(ctx context.Context, limit, skip int, sortField, sortOrder, search string) ([]domain.Product, error) {
	var products []domain.Product

	filter := bson.M{}
	if search != "" {
		filter["$text"] = bson.M{
			"$search": search,
		}
	}

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(skip))

	if sortField != "" {
		sortOrderValue := 1
		if sortOrder == "desc" {
			sortOrderValue = -1
		}
		findOptions.SetSort(bson.D{{Key: sortField, Value: sortOrderValue}})
	}

	cursor, err := p.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productRepository) GetProductByID(ctx context.Context, id primitive.ObjectID) (*domain.Product, error) {
	var product domain.Product

	err := p.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productRepository) GetProductByName(ctx context.Context, name string) (*domain.Product, error) {
	escapedName := regexp.QuoteMeta(name)
	regexPattern := bson.M{
		"$regex":   escapedName,
		"$options": "i",
	}

	filter := bson.M{
		"model_name": regexPattern,
	}

	var product domain.Product
	err := p.collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func (p *productRepository) CreateProduct(ctx context.Context, product *domain.Product) error {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	_, err := p.collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}

	return nil
}

func (p *productRepository) UpdateProduct(ctx context.Context, product *domain.Product) error {
	product.UpdatedAt = time.Now()

	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": product.ID}, bson.M{"$set": product})
	if err != nil {
		return err
	}

	return nil
}

func (p *productRepository) DeleteProduct(ctx context.Context, id primitive.ObjectID) error {
	_, err := p.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
