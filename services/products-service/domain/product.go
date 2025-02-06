package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	ModelName      string             `bson:"model_name"`
	CategoryID     primitive.ObjectID `bson:"category_id"`
	BrandID        primitive.ObjectID `bson:"brand_id"`
	TypeID         primitive.ObjectID `bson:"type_id"`
	Specifications Specifications     `bson:"specifications"`
	Content        string             `bson:"content"`
	LaptopImage    []string           `bson:"laptop_image"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
}

type Specifications struct {
	CPU               string `bson:"cpu"`
	CPUCores          int    `bson:"cpu_cores"`
	OperatingSystem   string `bson:"operating_system"`
	ScreenSize        string `bson:"screen_size"`
	ScreenRefreshRate string `bson:"screen_refresh_rate"`
	ScreenBrightness  string `bson:"screen_brightness"`
	ScreenType        string `bson:"screen_type"`
	Storage           string `bson:"storage"`
	Battery           string `bson:"battery"`
	RAM               string `bson:"ram"`
	Dimensions        string `bson:"dimensions"`
	Weight            string `bson:"weight"`
}
