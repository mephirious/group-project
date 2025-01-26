package usecase

import (
	"context"

	model "github.com/mephirious/group-project/internal/model"
)

type CategoryRepo interface {
	CreateCategory(ctx context.Context, category model.Category) (*model.Category, error)
	GetAllCategories(ctx context.Context) ([]model.Category, error)
}
