package cart

import "go.mongodb.org/mongo-driver/bson/primitive"

// Cart представляет корзину пользователя
type Cart struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID   string             `bson:"user_id" json:"user_id"`
	Products []CartItem         `bson:"products" json:"products"`
}

// CartItem представляет товар в корзине (с уникальным ID)
type CartItem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProductID string             `bson:"product_id" json:"product_id"`
	Amount    int                `bson:"amount" json:"amount"`
}
