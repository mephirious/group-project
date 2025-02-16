package repository

import (
	"context"
	"time"

	"github.com/mephirious/group-project/services/products-service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryRepository interface {
	GetAllInventories(ctx context.Context) ([]domain.Inventory, error)
	GetInventoryByProductID(ctx context.Context, productID primitive.ObjectID) (*domain.Inventory, error)
	CreateInventory(ctx context.Context, inventory *domain.Inventory) error
	UpdateInventory(ctx context.Context, inventory *domain.Inventory) error
	DeleteInventory(ctx context.Context, id primitive.ObjectID) error
}

type inventoryRepository struct {
	collection *mongo.Collection
}

func NewInventoryRepository(db *mongo.Database) *inventoryRepository {
	return &inventoryRepository{
		collection: db.Collection("inventories"),
	}
}

func (i *inventoryRepository) GetAllInventories(ctx context.Context) ([]domain.Inventory, error) {
	var inventories []domain.Inventory

	cursor, err := i.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &inventories)
	if err != nil {
		return nil, err
	}

	return inventories, nil
}

func (i *inventoryRepository) GetInventoryByProductID(ctx context.Context, productID primitive.ObjectID) (*domain.Inventory, error) {
	var inventory domain.Inventory

	err := i.collection.FindOne(ctx, bson.M{"product_id": productID}).Decode(&inventory)
	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (i *inventoryRepository) CreateInventory(ctx context.Context, inventory *domain.Inventory) error {
	inventory.CreatedAt = time.Now()
	inventory.UpdatedAt = time.Now()

	_, err := i.collection.InsertOne(ctx, inventory)
	if err != nil {
		return err
	}

	return nil
}

func (i *inventoryRepository) UpdateInventory(ctx context.Context, inventory *domain.Inventory) error {
	inventory.UpdatedAt = time.Now()

	_, err := i.collection.UpdateOne(ctx, bson.M{"_id": inventory.ID}, bson.M{"$set": inventory})
	if err != nil {
		return err
	}

	return nil
}

func (i *inventoryRepository) DeleteInventory(ctx context.Context, id primitive.ObjectID) error {
	_, err := i.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
