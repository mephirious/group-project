package domain

type Product struct {
	ID             string         `json:"_id"`
	Name           string         `json:"name"`
	Brand          string         `json:"brand"`
	CategoryID     string         `json:"category_id"`
	Price          float64        `json:"price"`
	Stock          int            `json:"stock"`
	Specifications Specifications `json:"specifications"`
	CreatedAt      string         `json:"created_at"`
	UpdatedAt      string         `json:"updated_at"`
}

type Specifications struct {
	Processor  string `json:"processor"`
	RAM        string `json:"RAM"`
	Storage    string `json:"storage"`
	GPU        string `json:"GPU"`
	ScreenSize string `json:"screen_size"`
	OS         string `json:"OS"`
}
