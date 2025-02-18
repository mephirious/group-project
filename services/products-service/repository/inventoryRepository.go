package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/mephirious/group-project/services/products-service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InventoryRepository interface {
	GetAllInventories(ctx context.Context) ([]domain.Inventory, error)
	GetInventoryByID(ctx context.Context, id primitive.ObjectID) (*domain.Inventory, error)
	GetInventoryByProductID(ctx context.Context, productID primitive.ObjectID) ([]domain.Inventory, error)
	GetInventoryBySerialNumber(ctx context.Context, serialNumber string) (*domain.Inventory, error)
	CreateInventory(ctx context.Context, inventory *domain.Inventory) error
	UpdateInventory(ctx context.Context, inventory *domain.Inventory) error
	DeleteInventory(ctx context.Context, id primitive.ObjectID) error
	GetProductQuantity(ctx context.Context, productID primitive.ObjectID) (int64, error)
	FindAndUpdateStatus(ctx context.Context, productID primitive.ObjectID, status string, limit int64) ([]domain.Inventory, error)
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

func (i *inventoryRepository) GetInventoryByID(ctx context.Context, id primitive.ObjectID) (*domain.Inventory, error) {
	var inventory domain.Inventory

	err := i.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&inventory)
	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (i *inventoryRepository) GetInventoryByProductID(ctx context.Context, productID primitive.ObjectID) ([]domain.Inventory, error) {
	var inventories []domain.Inventory

	cursor, err := i.collection.Find(ctx, bson.M{"product_id": productID})
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

func (i *inventoryRepository) GetInventoryBySerialNumber(ctx context.Context, serialNumber string) (*domain.Inventory, error) {
	var inventory domain.Inventory

	err := i.collection.FindOne(ctx, bson.M{"serial_number": serialNumber}).Decode(&inventory)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
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

func (i *inventoryRepository) GetProductQuantity(ctx context.Context, productID primitive.ObjectID) (int64, error) {
	pipeline := bson.A{
		bson.M{"$match": bson.M{"product_id": productID, "status": "in_stock"}},
		bson.M{"$group": bson.M{
			"_id":   nil,
			"total": bson.M{"$sum": 1},
		}},
	}

	cursor, err := i.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	var result struct {
		Total int64 `bson:"total"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return 0, err
		}
	}

	return result.Total, nil
}

func (i *inventoryRepository) FindAndUpdateStatus(ctx context.Context, productID primitive.ObjectID, status string, limit int64) ([]domain.Inventory, error) {
	filterStatus := "in_stock"
	if status == "in_stock" {
		filterStatus = "reserved"
	}

	filter := bson.M{"product_id": productID, "status": filterStatus}
	update := bson.M{"$set": bson.M{"status": status}}

	opts := options.Find().SetLimit(limit)
	cursor, err := i.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var inventories []domain.Inventory
	if err := cursor.All(ctx, &inventories); err != nil {
		return nil, err
	}

	var ids []primitive.ObjectID
	for _, inventory := range inventories {
		ids = append(ids, inventory.ID)
	}

	if len(ids) == 0 {
		return nil, fmt.Errorf("no inventory items found to update for product %s with status %s", productID.Hex(), filterStatus)
	}

	_, err = i.collection.UpdateMany(ctx, bson.M{"_id": bson.M{"$in": ids}}, update)
	if err != nil {
		return nil, err
	}

	return inventories, nil
}
