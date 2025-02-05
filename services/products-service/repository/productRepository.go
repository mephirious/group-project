package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/mephirious/group-project/services/products-service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	GetAllProducts(ctx context.Context) ([]domain.Product, error)
	GetProductByID(ctx context.Context, id string) (*domain.Product, error)
	CreateProduct(ctx context.Context, product *domain.Product) error
	UpdateProduct(ctx context.Context, product *domain.Product) error
	DeleteProduct(ctx context.Context, id string) error
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *productRepository {
	return &productRepository{
		collection: db.Collection("products"),
	}
}

func (p *productRepository) GetAllProducts(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product

	cursor, err := p.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &products)
	if err != nil {
		return nil, err
	}
	return products, err
}

func (p *productRepository) GetProductByID(ctx context.Context, id string) (*domain.Product, error) {
	var product domain.Product

	err := p.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}
	fmt.Println(product)

	return &product, nil
}

func (p *productRepository) CreateProduct(ctx context.Context, product *domain.Product) error {
	product.CreatedAt = fmt.Sprint(time.Now())
	product.UpdatedAt = fmt.Sprint(time.Now())

	_, err := p.collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}

	return nil
}

func (p *productRepository) UpdateProduct(ctx context.Context, product *domain.Product) error {
	product.UpdatedAt = fmt.Sprint(time.Now())

	_, err := p.collection.UpdateOne(ctx, bson.M{"_id": product.ID}, bson.M{"$set": product})
	if err != nil {
		return err
	}

	return nil
}

func (p *productRepository) DeleteProduct(ctx context.Context, id string) error {
	_, err := p.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
