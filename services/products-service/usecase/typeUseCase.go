package usecase

import (
	"context"
	"errors"

	"github.com/mephirious/group-project/services/products-service/domain"
	"github.com/mephirious/group-project/services/products-service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TypeUseCase interface {
	GetAllTypes(ctx context.Context) ([]domain.Type, error)
	GetTypeByID(ctx context.Context, id primitive.ObjectID) (*domain.Type, error)
	GetTypeByName(ctx context.Context, name string) (*domain.Type, error)
	CreateType(ctx context.Context, typeEntity *domain.Type) error
	UpdateType(ctx context.Context, typeEntity *domain.Type) error
	DeleteType(ctx context.Context, id primitive.ObjectID) error
}

type typeUseCase struct {
	typeRepository repository.TypeRepository
}

func NewTypeUseCase(repository repository.TypeRepository) *typeUseCase {
	return &typeUseCase{
		typeRepository: repository,
	}
}

func (t *typeUseCase) GetAllTypes(ctx context.Context) ([]domain.Type, error) {
	return t.typeRepository.GetAllTypes(ctx)
}

func (t *typeUseCase) GetTypeByID(ctx context.Context, id primitive.ObjectID) (*domain.Type, error) {
	typeEntity, err := t.typeRepository.GetTypeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if typeEntity == nil {
		return nil, errors.New("type not found")
	}
	return typeEntity, nil
}

func (t *typeUseCase) GetTypeByName(ctx context.Context, name string) (*domain.Type, error) {
	typeEntity, err := t.typeRepository.GetTypeByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if typeEntity == nil {
		return nil, errors.New("type not found")
	}
	return typeEntity, nil
}

func (t *typeUseCase) CreateType(ctx context.Context, typeEntity *domain.Type) error {
	existingType, err := t.typeRepository.GetTypeByName(ctx, typeEntity.TypeName)
	if err != nil {
		return err
	}
	if existingType != nil {
		return errors.New("type already exists")
	}

	return t.typeRepository.CreateType(ctx, typeEntity)
}

func (t *typeUseCase) UpdateType(ctx context.Context, typeEntity *domain.Type) error {
	return t.typeRepository.UpdateType(ctx, typeEntity)
}

func (t *typeUseCase) DeleteType(ctx context.Context, id primitive.ObjectID) error {
	return t.typeRepository.DeleteType(ctx, id)
}
