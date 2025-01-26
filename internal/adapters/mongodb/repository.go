package mongodb

import (
	"context"
	"fmt"

	model "github.com/mephirious/group-project/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// CreateProduct adds a new product to the "products" collection.
func (db *DB) CreateProduct(ctx context.Context, product model.Product) (*mongo.InsertOneResult, error) {
	collection := db.DB.Collection("products")
	return collection.InsertOne(ctx, product)
}

// GetAllProducts fetches all products from the "products" collection.
func (db *DB) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	collection := db.DB.Collection("products")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error fetching products: %v", err)
	}
	defer cursor.Close(ctx)

	var products []model.Product
	for cursor.Next(ctx) {
		var product model.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, fmt.Errorf("error decoding product: %v", err)
		}
		products = append(products, product)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}
	return products, nil
}

// GetProductByID fetches a product by its ID from the "products" collection.
func (db *DB) GetProductByID(ctx context.Context, id string) (*model.Product, error) {
	collection := db.DB.Collection("products")
	var product model.Product
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return nil, fmt.Errorf("error finding product: %v", err)
	}
	return &product, nil
}

// UpdateProduct updates a product's information in the "products" collection.
func (db *DB) UpdateProduct(ctx context.Context, id string, product model.Product) (*mongo.UpdateResult, error) {
	collection := db.DB.Collection("products")
	return collection.UpdateOne(ctx, bson.M{"_id": id}, bson.D{
		{Key: "$set", Value: product},
	})
}

// DeleteProduct deletes a product from the "products" collection.
func (db *DB) DeleteProduct(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	collection := db.DB.Collection("products")
	return collection.DeleteOne(ctx, bson.M{"_id": id})
}

// GetReviewsForProductByProductID fetches all reviews for a specific product from the "reviews" collection.
func (db *DB) GetReviewsForProductByProductID(ctx context.Context, productID string) ([]model.Review, error) {
	collection := db.DB.Collection("reviews")
	cursor, err := collection.Find(ctx, bson.M{"product_id": productID})
	if err != nil {
		return nil, fmt.Errorf("error fetching reviews: %v", err)
	}
	defer cursor.Close(ctx)

	var reviews []model.Review
	for cursor.Next(ctx) {
		var review model.Review
		if err := cursor.Decode(&review); err != nil {
			return nil, fmt.Errorf("error decoding review: %v", err)
		}
		reviews = append(reviews, review)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}
	return reviews, nil
}

// AddReviewForProduct adds a new review for a product to the "reviews" collection.
func (db *DB) AddReviewForProduct(ctx context.Context, productID string, review model.Review) (*mongo.InsertOneResult, error) {
	// Set the product ID in the review to associate it with the correct product
	review.ProductID = productID

	// Insert the review into the "reviews" collection
	reviewCollection := db.DB.Collection("reviews")
	insertResult, err := reviewCollection.InsertOne(ctx, review)
	if err != nil {
		return nil, fmt.Errorf("error inserting review: %v", err)
	}

	// Return the result of the insert
	return insertResult, nil
}

// CreateCategory adds a new category to the "categories" collection.
func (db *DB) CreateCategory(ctx context.Context, category model.Category) (*mongo.InsertOneResult, error) {
	collection := db.DB.Collection("categories")
	return collection.InsertOne(ctx, category)
}

// GetAllCategories fetches all categories from the "categories" collection.
func (db *DB) GetAllCategories(ctx context.Context) ([]model.Category, error) {
	collection := db.DB.Collection("categories")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error fetching categories: %v", err)
	}
	defer cursor.Close(ctx)

	var categories []model.Category
	for cursor.Next(ctx) {
		var category model.Category
		if err := cursor.Decode(&category); err != nil {
			return nil, fmt.Errorf("error decoding category: %v", err)
		}
		categories = append(categories, category)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}
	return categories, nil
}
