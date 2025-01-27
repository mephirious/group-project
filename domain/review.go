package domain

type Review struct {
	ID         string  `json:"_id" bson:"_id"`
	ProductID  string  `json:"product_id" bson:"product_id"`
	CustomerID string  `json:"customer_id" bson:"customer_id"`
	Rating     float64 `json:"rating" bson:"rating"`
	ReviewDate string  `json:"review_date" bson:"review_date"`
	Comment    string  `json:"comment" bson:"comment"`
}
