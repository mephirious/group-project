package config

type Config struct {
	Service struct {
		Name     string `yaml:"name" env:"SERVICE_NAME"`
		Port     int    `yaml:"port" env:"SERVICE_PORT"`
		LogLevel string `yaml:"log_level" env:"SERVICE_LOG_LEVEL"`
	} `yaml:"service"`
	MongoDB struct {
		URI         string `yaml:"uri" env:"MONGODB_URI"`
		Database    string `yaml:"database" env:"MONGODB_DATABASE"`
		Collections struct {
			Products   string `yaml:"products" env:"MONGODB_PRODUCTS"`
			Reviews    string `yaml:"reviews" env:"MONGODB_REVIEWS"`
			Categories string `yaml:"categories" env:"MONGODB_CATEGORIES"`
		} `yaml:"collections"`
	} `yaml:"mongodb"`
}
