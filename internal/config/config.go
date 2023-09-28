package config

type Config struct {
	Swagger SwaggerConfig
}
type SwaggerConfig struct {
	Enabled bool
	DirPath string
	URL     string
}
