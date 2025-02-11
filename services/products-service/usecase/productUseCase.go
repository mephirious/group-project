package usecase

import (
	"context"
	"errors"

	"github.com/mephirious/group-project/services/products-service/domain"
	"github.com/mephirious/group-project/services/products-service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductUseCase interface {
	GetAllProducts(ctx context.Context, limit, skip int, sortField, sortOrder, search string) ([]domain.Product, error)
	GetProductByID(ctx context.Context, id primitive.ObjectID) (*domain.Product, error)
	GetProductByName(ctx context.Context, name string) (*domain.Product, error)
	CreateProduct(ctx context.Context, product *domain.Product) error
	UpdateProduct(ctx context.Context, product *domain.Product) error
	DeleteProduct(ctx context.Context, id primitive.ObjectID) error
}

type productUseCase struct {
	productRepository repository.ProductRepository
}

func NewProductUseCase(repository repository.ProductRepository) *productUseCase {
	return &productUseCase{
		productRepository: repository,
	}
}

func (p *productUseCase) GetAllProducts(ctx context.Context, limit, skip int, sortField, sortOrder, search string) ([]domain.Product, error) {
	products, err := p.productRepository.GetAllProducts(ctx, limit, skip, sortField, sortOrder, search)
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, errors.New("no products found")
	}
	return products, nil
}

func (p *productUseCase) GetProductByID(ctx context.Context, id primitive.ObjectID) (*domain.Product, error) {
	product, err := p.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (p *productUseCase) GetProductByName(ctx context.Context, name string) (*domain.Product, error) {
	product, err := p.productRepository.GetProductByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (p *productUseCase) CreateProduct(ctx context.Context, product *domain.Product) error {
	existingProduct, err := p.productRepository.GetProductByName(ctx, product.ModelName)
	if err != nil {
		return err
	}
	if existingProduct != nil {
		return errors.New("product already exists")
	}

	return p.productRepository.CreateProduct(ctx, product)
}

func (p *productUseCase) UpdateProduct(ctx context.Context, product *domain.Product) error {
	existingProduct, err := p.productRepository.GetProductByID(ctx, product.ID)
	if err != nil {
		return err
	}
	if existingProduct == nil {
		return errors.New("product not found")
	}

	return p.productRepository.UpdateProduct(ctx, product)
}

func (p *productUseCase) DeleteProduct(ctx context.Context, id primitive.ObjectID) error {
	product, err := p.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("product not found")
	}

	return p.productRepository.DeleteProduct(ctx, id)
}
