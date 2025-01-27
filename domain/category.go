package domain

type Category struct {
	ID          string `json:"_id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	CreatedAt   string `json:"created_at" bson:"created_at"`
	UpdatedAt   string `json:"updated_at" bson:"updated_at"`
}
