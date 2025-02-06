package cart

import (
	"context"
	"errors"

	"github.com/mephirious/group-project/services/customer/internal/database"
	"github.com/mephirious/group-project/services/customer/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CartRepository - репозиторий для работы с корзиной
type CartRepository struct {
	Collection string
}

// NewCartRepository создает новый репозиторий корзины
func NewCartRepository() *CartRepository {
	return &CartRepository{
		Collection: "carts",
	}
}

// GetCart возвращает корзину по userID
func (r *CartRepository) GetCart(ctx context.Context, userID string) (*Cart, error) {
	logger.Log.Infof("Fetching cart for user_id: %s", userID)
	collection := database.GetCollection("cart_service", r.Collection)
	var cart Cart
	err := collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&cart)
	if err != nil {
		logger.Log.Warnf("Cart not found for user_id: %s", userID)
		return nil, err
	}
	return &cart, nil
}

// AddToCart добавляет товар в корзину
func (r *CartRepository) AddToCart(ctx context.Context, userID, productID string, amount int) error {
	logger.Log.Infof("Adding product %s (amount: %d) to cart of user %s", productID, amount, userID)
	collection := database.GetCollection("cart_service", r.Collection)
	cart, _ := r.GetCart(ctx, userID)

	if cart == nil {
		newCart := Cart{
			ID:     primitive.NewObjectID(),
			UserID: userID,
			Products: []CartItem{
				{ID: primitive.NewObjectID(), ProductID: productID, Amount: amount},
			},
		}
		_, err := collection.InsertOne(ctx, newCart)
		if err != nil {
			logger.Log.Errorf("Failed to create new cart for user_id: %s, error: %v", userID, err)
		}
		return err
	}

	// Проверяем, есть ли уже этот продукт в корзине
	for i, item := range cart.Products {
		if item.ProductID == productID {
			cart.Products[i].Amount += amount
			_, err := collection.UpdateOne(ctx, bson.M{"user_id": userID}, bson.M{"$set": bson.M{"products": cart.Products}})
			if err != nil {
				logger.Log.Errorf("Failed to update product %s in cart of user %s", productID, userID)
			}
			return err
		}
	}

	// Добавляем новый товар
	newItem := CartItem{
		ID:        primitive.NewObjectID(),
		ProductID: productID,
		Amount:    amount,
	}
	cart.Products = append(cart.Products, newItem)

	_, err := collection.UpdateOne(ctx, bson.M{"user_id": userID}, bson.M{"$set": bson.M{"products": cart.Products}})
	if err != nil {
		logger.Log.Errorf("Failed to add product %s to cart of user %s", productID, userID)
	}
	return err
}

// RemoveFromCart удаляет товар по его ID в корзине
func (r *CartRepository) RemoveFromCart(ctx context.Context, userID, cartItemID string) error {
	logger.Log.Infof("Removing item %s from cart of user %s", cartItemID, userID)
	collection := database.GetCollection("cart_service", r.Collection)
	cart, err := r.GetCart(ctx, userID)
	if err != nil {
		logger.Log.Warnf("Cart not found for user_id: %s", userID)
		return err
	}

	itemID, _ := primitive.ObjectIDFromHex(cartItemID)

	updatedProducts := []CartItem{}
	for _, item := range cart.Products {
		if item.ID != itemID {
			updatedProducts = append(updatedProducts, item)
		}
	}

	_, err = collection.UpdateOne(ctx, bson.M{"user_id": userID}, bson.M{"$set": bson.M{"products": updatedProducts}})
	if err != nil {
		logger.Log.Errorf("Failed to remove item %s from cart of user %s", cartItemID, userID)
	}
	return err
}

// ClearCart очищает корзину
func (r *CartRepository) ClearCart(ctx context.Context, userID string) error {
	logger.Log.Infof("Clearing cart of user %s", userID)
	collection := database.GetCollection("cart_service", r.Collection)
	_, err := collection.UpdateOne(ctx, bson.M{"user_id": userID}, bson.M{"$set": bson.M{"products": []CartItem{}}})
	if err != nil {
		logger.Log.Errorf("Failed to clear cart for user %s", userID)
	}
	return err
}

// UpdateCartItem обновляет количество товара или удаляет его, если amount < 1
func (r *CartRepository) UpdateCartItem(ctx context.Context, userID, cartItemID string, newAmount int) (string, error) {
	logger.Log.Infof("Updating item %s in cart of user %s to amount %d", cartItemID, userID, newAmount)
	collection := database.GetCollection("cart_service", r.Collection)
	cart, err := r.GetCart(ctx, userID)
	if err != nil {
		logger.Log.Warnf("Cart not found for user_id: %s", userID)
		return "", err
	}

	itemID, err := primitive.ObjectIDFromHex(cartItemID)
	if err != nil {
		logger.Log.Errorf("Invalid item ID format: %s", cartItemID)
		return "", errors.New("invalid item ID format")
	}

	// Если количество меньше 1, удаляем товар и возвращаем специальное сообщение
	if newAmount < 1 {
		logger.Log.Infof("Amount is %d, removing item %s from cart of user %s", newAmount, cartItemID, userID)
		err := r.RemoveFromCart(ctx, userID, cartItemID)
		if err != nil {
			return "", err
		}
		return "Item removed from cart because amount was less than 1", nil
	}

	// Обновляем количество товара
	for i, item := range cart.Products {
		if item.ID == itemID {
			cart.Products[i].Amount = newAmount
			_, err := collection.UpdateOne(ctx, bson.M{"user_id": userID}, bson.M{"$set": bson.M{"products": cart.Products}})
			if err != nil {
				logger.Log.Errorf("Failed to update item %s in cart of user %s", cartItemID, userID)
				return "", err
			}
			return "Item amount updated successfully", nil
		}
	}

	logger.Log.Warnf("Item %s not found in cart of user %s", cartItemID, userID)
	return "", errors.New("item not found in cart")
}
