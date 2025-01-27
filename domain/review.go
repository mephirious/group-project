package domain

type Review struct {
	ID         string  `json:"_id"`
	ProductID  string  `json:"product_id"`
	CustomerID string  `json:"customer_id"`
	Rating     float64 `json:"rating"`
	ReviewDate string  `json:"review_date"`
	Comment    string  `json:"comment"`
}
