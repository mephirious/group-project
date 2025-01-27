package usecase

import (
	"context"
	"errors"

	"github.com/mephirious/group-project/products-service/domain"
	"github.com/mephirious/group-project/products-service/repository"
)

type ProductUseCase interface {
	GetAllProducts(ctx context.Context) ([]domain.Product, error)
	GetProductByID(ctx context.Context, id string) (*domain.Product, error)
	CreateProduct(ctx context.Context, product *domain.Product) error
	UpdateProduct(ctx context.Context, product *domain.Product) error
	DeleteProduct(ctx context.Context, id string) error
}

type productUseCase struct {
	productRepository repository.ProductRepository
}

func NewProductUseCase(repository repository.ProductRepository) *productUseCase {
	return &productUseCase{
		productRepository: repository,
	}
}

func (p *productUseCase) GetAllProducts(ctx context.Context) ([]domain.Product, error) {
	return p.productRepository.GetAllProducts(ctx)
}

func (p *productUseCase) GetProductByID(ctx context.Context, id string) (*domain.Product, error) {
	product, err := p.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (p *productUseCase) CreateProduct(ctx context.Context, product *domain.Product) error {
	return p.productRepository.CreateProduct(ctx, product)
}

func (p *productUseCase) UpdateProduct(ctx context.Context, product *domain.Product) error {
	return p.productRepository.UpdateProduct(ctx, product)
}

func (p *productUseCase) DeleteProduct(ctx context.Context, id string) error {
	return p.productRepository.DeleteProduct(ctx, id)
}
