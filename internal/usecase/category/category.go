package usecase

import (
	"context"

	model "github.com/mephirious/group-project/internal/model"
)

// Category defines the structure for handling category-related operations.
type Category struct {
	categoryRepo CategoryRepo
}

// NewCategory creates a new instance of the Category use case.
func NewCategory(categoryRepo CategoryRepo) *Category {
	return &Category{
		categoryRepo: categoryRepo,
	}
}

// Create creates a new category.
func (uc *Category) Create(ctx context.Context, category model.Category) (*model.Category, error) {
	return uc.categoryRepo.CreateCategory(ctx, category)
}

// GetAll retrieves all categories.
func (uc *Category) GetAll(ctx context.Context) ([]model.Category, error) {
	return uc.categoryRepo.GetAllCategories(ctx)
}
