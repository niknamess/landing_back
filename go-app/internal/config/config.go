package config

import "os"

type ConfigSwagger struct {
	Swagger SwaggerConfig
}
type SwaggerConfig struct {
	Enabled bool   `env:"ENABLED"`
	DirPath string `env:"DIRPATH"`
	URL     string `env:"URL"`
}

const (
	prod = "production"
)

// Config object
type Config struct {
	Env           string        `env:"ENV"`
	MongoDB       MongoDBConfig `json:"mongodb"`
	Host          string        `env:"APP_HOST"`
	Port          string        `env:"APP_PORT"`
	ConfigSwagger SwaggerConfig `json:"swagger"`
}

// IsProd Checks if env is production
func (c Config) IsProd() bool {
	return c.Env == prod
}

// GetConfig gets all config for the application
func GetConfig() Config {
	return Config{
		Env:           os.Getenv("ENV"),
		MongoDB:       GetMongoDBConfig(),
		Host:          os.Getenv("APP_HOST"),
		Port:          os.Getenv("APP_PORT"),
		ConfigSwagger: GetSwaggerConfig(),
	}
}

// MongoDBConfig object
type MongoDBConfig struct {
	URI string `env:"MONGO_URI"` // i.e. "mongodb://localhost:27017"
}

// GetMongoDBConfig returns MongoDBConfig object
func GetMongoDBConfig() MongoDBConfig {
	return MongoDBConfig{
		URI: os.Getenv("MONGO_URI"),
	}
}

// GetMongoDBConfig returns MongoDBConfig object
func GetSwaggerConfig() SwaggerConfig {
	return SwaggerConfig{
		Enabled: false,
		DirPath: os.Getenv("DIRPATH"),
		URL:     os.Getenv("URL"),
	}
}
