package usecase

import (
	"context"

	model "github.com/mephirious/group-project/internal/model"
)

type ProductRepo interface {
	CreateProduct(ctx context.Context, product model.Product) (*model.Product, error)
	GetAllProducts(ctx context.Context) ([]model.Product, error)
	GetProductByID(ctx context.Context, id string) (*model.Product, error)
	UpdateProduct(ctx context.Context, id string, product model.Product) (*model.Product, error)
	DeleteProduct(ctx context.Context, id string) (*model.Product, error)
}
