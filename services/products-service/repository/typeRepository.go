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

type TypeRepository interface {
	GetAllTypes(ctx context.Context) ([]domain.Type, error)
	GetTypeByID(ctx context.Context, id primitive.ObjectID) (*domain.Type, error)
	GetTypeByName(ctx context.Context, name string) (*domain.Type, error)
	CreateType(ctx context.Context, Type *domain.Type) error
	UpdateType(ctx context.Context, Type *domain.Type) error
	DeleteType(ctx context.Context, id primitive.ObjectID) error
}

type typeRepository struct {
	collection *mongo.Collection
}

func NewTypeRepository(db *mongo.Database) *typeRepository {
	return &typeRepository{
		collection: db.Collection("types"),
	}
}

func (t *typeRepository) GetAllTypes(ctx context.Context) ([]domain.Type, error) {
	var types []domain.Type

	cursor, err := t.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &types)
	if err != nil {
		return nil, err
	}

	return types, nil
}

func (t *typeRepository) GetTypeByID(ctx context.Context, id primitive.ObjectID) (*domain.Type, error) {
	var typeEntity domain.Type

	err := t.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&typeEntity)
	if err != nil {
		return nil, err
	}

	fmt.Println(typeEntity)
	return &typeEntity, nil
}

func (t *typeRepository) GetTypeByName(ctx context.Context, name string) (*domain.Type, error) {
	regexPattern := bson.M{
		"$regex":   name,
		"$options": "i",
	}

	filter := bson.M{
		"type_name": regexPattern, // Adjusted field name here
	}

	var typeEntity domain.Type
	err := t.collection.FindOne(ctx, filter).Decode(&typeEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &typeEntity, nil
}

func (t *typeRepository) CreateType(ctx context.Context, typeEntity *domain.Type) error {
	typeEntity.CreatedAt = time.Now()
	typeEntity.UpdatedAt = time.Now()

	_, err := t.collection.InsertOne(ctx, typeEntity)
	if err != nil {
		return err
	}

	return nil
}

func (t *typeRepository) UpdateType(ctx context.Context, typeEntity *domain.Type) error {
	typeEntity.UpdatedAt = time.Now()

	_, err := t.collection.UpdateOne(ctx, bson.M{"_id": typeEntity.ID}, bson.M{"$set": typeEntity})
	if err != nil {
		return err
	}

	return nil
}

func (t *typeRepository) DeleteType(ctx context.Context, id primitive.ObjectID) error {
	_, err := t.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
