package domain

import "time"

type Product struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"model_name" bson:"model_name"`
	Price    int64  `json:"price" bson:"price"`
	Quantity int64  `json:"quantity" bson:"quantity"`
	Currency string `json:"currency" bson:"currency"`
}

type Order struct {
	Products  []Product `bson:"products" json:"products"`
	Amount    int64     `bson:"amount" json:"amount"`
	Status    string    `bson:"status" json:"status"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}
