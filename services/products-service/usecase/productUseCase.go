package usecase

import (
	"context"
	"errors"

	"github.com/mephirious/group-project/services/products-service/domain"
	"github.com/mephirious/group-project/services/products-service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductUseCase interface {
	GetAllProducts(ctx context.Context, limit, skip int, sortField, sortOrder, search string) ([]domain.ProductView, error)
	GetProductByID(ctx context.Context, id primitive.ObjectID) (*domain.ProductView, error)
	GetProductByName(ctx context.Context, name string) (*domain.ProductView, error)
	CreateProduct(ctx context.Context, product *domain.Product) error
	UpdateProduct(ctx context.Context, product *domain.Product) error
	DeleteProduct(ctx context.Context, id primitive.ObjectID) error
}

type productUseCase struct {
	productRepository  repository.ProductRepository
	categoryRepository repository.CategoryRepository
	brandRepository    repository.BrandRepository
	typeRepository     repository.TypeRepository
}

func NewProductUseCase(productRepository repository.ProductRepository, categoryRepository repository.CategoryRepository, brandRepository repository.BrandRepository, typeRepository repository.TypeRepository) *productUseCase {
	return &productUseCase{
		productRepository:  productRepository,
		categoryRepository: categoryRepository,
		brandRepository:    brandRepository,
		typeRepository:     typeRepository,
	}
}

func (p *productUseCase) GetAllProducts(ctx context.Context, limit, skip int, sortField, sortOrder, search string) ([]domain.ProductView, error) {
	products, err := p.productRepository.GetAllProducts(ctx, limit, skip, sortField, sortOrder, search)
	if err != nil {
		return nil, err
	}

	// convert product to product view
	productViews := make([]domain.ProductView, len(products))

	// set category, brand, type names
	for i, product := range products {
		Category, err := p.categoryRepository.GetCategoryByID(ctx, product.CategoryID)
		if err != nil {
			Category = &domain.Category{CategoryName: "Unknown"}
		}
		Brand, err := p.brandRepository.GetBrandByID(ctx, product.BrandID)
		if err != nil {
			Brand = &domain.Brand{BrandName: "Unknown"}
		}
		Type, err := p.typeRepository.GetTypeByID(ctx, product.TypeID)
		if err != nil {
			Type = &domain.Type{TypeName: "Unknown"}
		}
		productViews[i] = domain.ProductView{
			ID:             product.ID,
			ModelName:      product.ModelName,
			Price:          product.Price,
			Category:       Category.CategoryName,
			Brand:          Brand.BrandName,
			Type:           Type.TypeName,
			Specifications: product.Specifications,
			Content:        product.Content,
			LaptopImage:    product.LaptopImage,
			CreatedAt:      product.CreatedAt,
			UpdatedAt:      product.UpdatedAt,
		}
	}
	return productViews, nil
}

func (p *productUseCase) GetProductByID(ctx context.Context, id primitive.ObjectID) (*domain.ProductView, error) {
	product, err := p.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	// Get category, brand, type names
	Category, err := p.categoryRepository.GetCategoryByID(ctx, product.CategoryID)
	if err != nil {
		Category = &domain.Category{CategoryName: "Unknown"}
	}
	Brand, err := p.brandRepository.GetBrandByID(ctx, product.BrandID)
	if err != nil {
		Brand = &domain.Brand{BrandName: "Unknown"}
	}
	Type, err := p.typeRepository.GetTypeByID(ctx, product.TypeID)
	if err != nil {
		Type = &domain.Type{TypeName: "Unknown"}
	}

	return &domain.ProductView{
		ID:             product.ID,
		ModelName:      product.ModelName,
		Price:          product.Price,
		Category:       Category.CategoryName,
		Brand:          Brand.BrandName,
		Type:           Type.TypeName,
		Specifications: product.Specifications,
		Content:        product.Content,
		LaptopImage:    product.LaptopImage,
		CreatedAt:      product.CreatedAt,
		UpdatedAt:      product.UpdatedAt,
	}, nil
}

func (p *productUseCase) GetProductByName(ctx context.Context, name string) (*domain.ProductView, error) {
	product, err := p.productRepository.GetProductByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	// Get category, brand, type names
	Category, err := p.categoryRepository.GetCategoryByID(ctx, product.CategoryID)
	if err != nil {
		Category = &domain.Category{CategoryName: "Unknown"}
	}
	Brand, err := p.brandRepository.GetBrandByID(ctx, product.BrandID)
	if err != nil {
		Brand = &domain.Brand{BrandName: "Unknown"}
	}
	Type, err := p.typeRepository.GetTypeByID(ctx, product.TypeID)
	if err != nil {
		Type = &domain.Type{TypeName: "Unknown"}
	}

	return &domain.ProductView{
		ID:             product.ID,
		ModelName:      product.ModelName,
		Price:          product.Price,
		Category:       Category.CategoryName,
		Brand:          Brand.BrandName,
		Type:           Type.TypeName,
		Specifications: product.Specifications,
		Content:        product.Content,
		LaptopImage:    product.LaptopImage,
		CreatedAt:      product.CreatedAt,
		UpdatedAt:      product.UpdatedAt,
	}, nil
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
