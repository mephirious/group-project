package usecase

import (
	"context"

	model "github.com/mephirious/group-project/internal/model"
)

// Product defines the structure for handling product-related operations.
type Product struct {
	productRepo ProductRepo
}

// NewProduct creates a new instance of the Product use case.
func NewProduct(productRepo ProductRepo) *Product {
	return &Product{
		productRepo: productRepo,
	}
}

// Create creates a new product.
func (uc *Product) Create(ctx context.Context, product model.Product) (*model.Product, error) {
	return uc.productRepo.CreateProduct(ctx, product)
}

// GetAll retrieves all products.
func (uc *Product) GetAll(ctx context.Context) ([]model.Product, error) {
	return uc.productRepo.GetAllProducts(ctx)
}

// Get retrieves a product by its ID.
func (uc *Product) Get(ctx context.Context, id string) (*model.Product, error) {
	return uc.productRepo.GetProductByID(ctx, id)
}

// Update updates an existing product.
func (uc *Product) Update(ctx context.Context, id string, product model.Product) (*model.Product, error) {
	return uc.productRepo.UpdateProduct(ctx, id, product)
}

// Delete removes a product.
func (uc *Product) Delete(ctx context.Context, id string) (*model.Product, error) {
	return uc.productRepo.DeleteProduct(ctx, id)
}
