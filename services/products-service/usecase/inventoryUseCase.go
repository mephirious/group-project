package usecase

import (
	"context"
	"errors"

	"github.com/mephirious/group-project/services/products-service/domain"
	"github.com/mephirious/group-project/services/products-service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryUseCase interface {
	GetAllInventories(ctx context.Context) ([]domain.Inventory, error)
	GetInventoryByProductID(ctx context.Context, productID primitive.ObjectID) (*domain.Inventory, error)
	CreateInventory(ctx context.Context, inventory *domain.Inventory) error
	UpdateInventory(ctx context.Context, inventory *domain.Inventory) error
	DeleteInventory(ctx context.Context, id primitive.ObjectID) error
}

type inventoryUseCase struct {
	inventoryRepository repository.InventoryRepository
}

func NewInventoryUseCase(repository repository.InventoryRepository) *inventoryUseCase {
	return &inventoryUseCase{
		inventoryRepository: repository,
	}
}

func (i *inventoryUseCase) GetAllInventories(ctx context.Context) ([]domain.Inventory, error) {
	return i.inventoryRepository.GetAllInventories(ctx)
}

func (i *inventoryUseCase) GetInventoryByProductID(ctx context.Context, productID primitive.ObjectID) (*domain.Inventory, error) {
	inventory, err := i.inventoryRepository.GetInventoryByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	if inventory == nil {
		return nil, errors.New("inventory not found")
	}
	return inventory, nil
}

func (i *inventoryUseCase) CreateInventory(ctx context.Context, inventory *domain.Inventory) error {
	return i.inventoryRepository.CreateInventory(ctx, inventory)
}

func (i *inventoryUseCase) UpdateInventory(ctx context.Context, inventory *domain.Inventory) error {
	return i.inventoryRepository.UpdateInventory(ctx, inventory)
}

func (i *inventoryUseCase) DeleteInventory(ctx context.Context, id primitive.ObjectID) error {
	return i.inventoryRepository.DeleteInventory(ctx, id)
}
