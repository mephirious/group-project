package usecase

import (
	"context"

	"github.com/mephirious/group-project/services/products-service/domain"
	"github.com/mephirious/group-project/services/products-service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryUseCase interface {
	GetAllInventories(ctx context.Context) ([]domain.Inventory, error)
	GetInventoryByID(ctx context.Context, id primitive.ObjectID) (*domain.Inventory, error)
	GetInventoryByProductID(ctx context.Context, productID primitive.ObjectID) ([]domain.Inventory, error)
	GetInventoryBySerialNumber(ctx context.Context, serialNumber string) (*domain.Inventory, error)
	CreateInventory(ctx context.Context, inventory *domain.Inventory) error
	UpdateInventory(ctx context.Context, inventory *domain.Inventory) error
	DeleteInventory(ctx context.Context, id primitive.ObjectID) error
	GetProductQuantity(ctx context.Context, productID primitive.ObjectID) (int64, error)
	ReserveProducts(ctx context.Context, order domain.Order) error
	CancelReservation(ctx context.Context, order domain.Order) error
	MarkProductsAsSold(ctx context.Context, order domain.Order) error
}

type inventoryUseCase struct {
	repo repository.InventoryRepository
}

func NewInventoryUseCase(repo repository.InventoryRepository) InventoryUseCase {
	return &inventoryUseCase{repo: repo}
}

func (i *inventoryUseCase) GetAllInventories(ctx context.Context) ([]domain.Inventory, error) {
	return i.repo.GetAllInventories(ctx)
}

func (i *inventoryUseCase) GetInventoryByID(ctx context.Context, id primitive.ObjectID) (*domain.Inventory, error) {
	return i.repo.GetInventoryByID(ctx, id)
}

func (i *inventoryUseCase) GetInventoryByProductID(ctx context.Context, productID primitive.ObjectID) ([]domain.Inventory, error) {
	return i.repo.GetInventoryByProductID(ctx, productID)
}

func (i *inventoryUseCase) GetInventoryBySerialNumber(ctx context.Context, serialNumber string) (*domain.Inventory, error) {
	return i.repo.GetInventoryBySerialNumber(ctx, serialNumber)
}

func (i *inventoryUseCase) CreateInventory(ctx context.Context, inventory *domain.Inventory) error {
	return i.repo.CreateInventory(ctx, inventory)
}

func (i *inventoryUseCase) UpdateInventory(ctx context.Context, inventory *domain.Inventory) error {
	return i.repo.UpdateInventory(ctx, inventory)
}

func (i *inventoryUseCase) DeleteInventory(ctx context.Context, id primitive.ObjectID) error {
	return i.repo.DeleteInventory(ctx, id)
}

func (i *inventoryUseCase) GetProductQuantity(ctx context.Context, productID primitive.ObjectID) (int64, error) {
	return i.repo.GetProductQuantity(ctx, productID)
}

func (i *inventoryUseCase) ReserveProducts(ctx context.Context, order domain.Order) error {
	for _, product := range order.Products {
		productID, err := primitive.ObjectIDFromHex(product.ID)
		if err != nil {
			return err
		}

		_, err = i.repo.FindAndUpdateStatus(ctx, productID, "reserved", product.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *inventoryUseCase) CancelReservation(ctx context.Context, order domain.Order) error {
	for _, product := range order.Products {
		productID, err := primitive.ObjectIDFromHex(product.ID)
		if err != nil {
			return err
		}

		_, err = i.repo.FindAndUpdateStatus(ctx, productID, "in_stock", product.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *inventoryUseCase) MarkProductsAsSold(ctx context.Context, order domain.Order) error {
	for _, product := range order.Products {
		productID, err := primitive.ObjectIDFromHex(product.ID)
		if err != nil {
			return err
		}

		_, err = i.repo.FindAndUpdateStatus(ctx, productID, "sold", product.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}
