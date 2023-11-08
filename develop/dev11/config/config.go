package config

type Config struct {
	HttpServerPort string
}

// можно было использовать viper и сделать файл yaml
func GetConfig() Config {
	return Config{
		HttpServerPort: ":8080",
	}
}
