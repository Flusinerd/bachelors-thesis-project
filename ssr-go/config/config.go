package config

type Config struct {
	Server *ServerConfig
}

func LoadConfig() *Config {
	return &Config{
		Server: LoadServerConfig(),
	}
}
