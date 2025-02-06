package usecase

import (
	"context"
	"errors"

	"github.com/mephirious/group-project/services/products-service/domain"
	"github.com/mephirious/group-project/services/products-service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BrandUseCase interface {
	GetAllBrands(ctx context.Context) ([]domain.Brand, error)
	GetBrandByID(ctx context.Context, id primitive.ObjectID) (*domain.Brand, error)
	GetBrandByName(ctx context.Context, name string) (*domain.Brand, error)
	CreateBrand(ctx context.Context, brand *domain.Brand) error
	UpdateBrand(ctx context.Context, brand *domain.Brand) error
	DeleteBrand(ctx context.Context, id primitive.ObjectID) error
}

type brandUseCase struct {
	brandRepository repository.BrandRepository
}

func NewBrandUseCase(repository repository.BrandRepository) *brandUseCase {
	return &brandUseCase{
		brandRepository: repository,
	}
}

func (b *brandUseCase) GetAllBrands(ctx context.Context) ([]domain.Brand, error) {
	return b.brandRepository.GetAllBrands(ctx)
}

func (b *brandUseCase) GetBrandByID(ctx context.Context, id primitive.ObjectID) (*domain.Brand, error) {
	brand, err := b.brandRepository.GetBrandByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if brand == nil {
		return nil, errors.New("brand not found")
	}
	return brand, nil
}

func (b *brandUseCase) GetBrandByName(ctx context.Context, name string) (*domain.Brand, error) {
	brand, err := b.brandRepository.GetBrandByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if brand == nil {
		return nil, errors.New("brand not found")
	}
	return brand, nil
}

func (b *brandUseCase) CreateBrand(ctx context.Context, brand *domain.Brand) error {
	return b.brandRepository.CreateBrand(ctx, brand)
}

func (b *brandUseCase) UpdateBrand(ctx context.Context, brand *domain.Brand) error {
	return b.brandRepository.UpdateBrand(ctx, brand)
}

func (b *brandUseCase) DeleteBrand(ctx context.Context, id primitive.ObjectID) error {
	return b.brandRepository.DeleteBrand(ctx, id)
}
