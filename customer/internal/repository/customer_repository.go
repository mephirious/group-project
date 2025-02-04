package repository

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/mephirious/group-project/services/customer/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CustomerRepository управляет клиентами в MongoDB
type CustomerRepository struct {
	Collection *mongo.Collection
}

// NewCustomerRepository создает новый репозиторий
func NewCustomerRepository(db *mongo.Database) *CustomerRepository {
	return &CustomerRepository{
		Collection: db.Collection("customers"),
	}
}

// GetCustomerByID ищет пользователя по ID
func (repo *CustomerRepository) GetCustomerByID(id int) (*models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idString := strconv.Itoa(id)

	objID, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return nil, errors.New("неверный формат ID")
	}

	var customer models.Customer
	err = repo.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
