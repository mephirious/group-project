package domain

type Product struct {
	ID             string         `json:"_id" bson:"_id"`
	Name           string         `json:"name" bson:"name"`
	Brand          string         `json:"brand" bson:"brand"`
	CategoryID     string         `json:"category_id" bson:"category_id"`
	Price          float64        `json:"price" bson:"price"`
	Stock          int            `json:"stock" bson:"stock"`
	Specifications Specifications `json:"specifications" bson:"specifications"`
	CreatedAt      string         `json:"created_at" bson:"created_at"`
	UpdatedAt      string         `json:"updated_at" bson:"updated_at"`
}

type Specifications struct {
	Processor  string `json:"processor" bson:"processor"`
	RAM        string `json:"RAM" bson:"RAM"`
	Storage    string `json:"storage" bson:"storage"`
	GPU        string `json:"GPU" bson:"GPU"`
	ScreenSize string `json:"screen_size" bson:"screen_size"`
	OS         string `json:"OS" bson:"OS"`
}
