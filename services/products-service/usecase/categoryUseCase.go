package usecase

import (
	"context"
	"errors"

	"github.com/mephirious/group-project/services/products-service/domain"
	"github.com/mephirious/group-project/services/products-service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryUseCase interface {
	GetAllCategories(ctx context.Context) ([]domain.Category, error)
	GetCategoryByID(ctx context.Context, id primitive.ObjectID) (*domain.Category, error)
	GetCategoryByName(ctx context.Context, name string) (*domain.Category, error)
	CreateCategory(ctx context.Context, category *domain.Category) error
	UpdateCategory(ctx context.Context, category *domain.Category) error
	DeleteCategory(ctx context.Context, id primitive.ObjectID) error
}

type categoryUseCase struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryUseCase(repository repository.CategoryRepository) *categoryUseCase {
	return &categoryUseCase{
		categoryRepository: repository,
	}
}

func (c *categoryUseCase) GetAllCategories(ctx context.Context) ([]domain.Category, error) {
	return c.categoryRepository.GetAllCategories(ctx)
}

func (c *categoryUseCase) GetCategoryByID(ctx context.Context, id primitive.ObjectID) (*domain.Category, error) {
	category, err := c.categoryRepository.GetCategoryByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}

func (c *categoryUseCase) GetCategoryByName(ctx context.Context, name string) (*domain.Category, error) {
	category, err := c.categoryRepository.GetCategoryByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}

func (c *categoryUseCase) CreateCategory(ctx context.Context, category *domain.Category) error {
	return c.categoryRepository.CreateCategory(ctx, category)
}

func (c *categoryUseCase) UpdateCategory(ctx context.Context, category *domain.Category) error {
	return c.categoryRepository.UpdateCategory(ctx, category)
}

func (c *categoryUseCase) DeleteCategory(ctx context.Context, id primitive.ObjectID) error {
	return c.categoryRepository.DeleteCategory(ctx, id)
}
