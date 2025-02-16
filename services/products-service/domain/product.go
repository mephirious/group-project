package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ModelName      string             `bson:"model_name" json:"model_name"`
	Price          float64            `bson:"price" json:"price"`
	CategoryID     primitive.ObjectID `bson:"category_id" json:"category_id"`
	BrandID        primitive.ObjectID `bson:"brand_id" json:"brand_id"`
	TypeID         primitive.ObjectID `bson:"type_id" json:"type_id"`
	Specifications Specifications     `bson:"specifications" json:"specifications"`
	Content        string             `bson:"content" json:"content"`
	Images         []string           `bson:"laptop_image" json:"images"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}

type Specifications struct {
	CPU               string `bson:"cpu" json:"cpu"`
	CPUCores          int    `bson:"cpu_cores" json:"cpu_cores"`
	OperatingSystem   string `bson:"operating_system" json:"operating_system"`
	ScreenSize        string `bson:"screen_size" json:"screen_size"`
	ScreenRefreshRate string `bson:"screen_refresh_rate" json:"screen_refresh_rate"`
	ScreenBrightness  string `bson:"screen_brightness" json:"screen_brightness"`
	ScreenType        string `bson:"screen_type" json:"screen_type"`
	Storage           string `bson:"storage" json:"storage"`
	Battery           string `bson:"battery" json:"battery"`
	RAM               string `bson:"ram" json:"ram"`
	Dimensions        string `bson:"dimensions" json:"dimensions"`
	Weight            string `bson:"weight" json:"weight"`
}

type ProductView struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ModelName      string             `bson:"model_name" json:"model_name"`
	Price          float64            `bson:"price" json:"price"`
	Category       string             `bson:"category" json:"category"`
	Brand          string             `bson:"brand" json:"brand"`
	Type           string             `bson:"type" json:"type"`
	Specifications Specifications     `bson:"specifications" json:"specifications"`
	Content        string             `bson:"content" json:"content"`
	Images         []string           `bson:"laptop_image" json:"images"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}
