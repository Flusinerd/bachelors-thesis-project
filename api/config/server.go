package config

import (
	"os"
	"strconv"
)

type ServerConfig struct {
	Port                int
	ResponseDelayFactor int
}

func LoadServerConfig() ServerConfig {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	portInt, err := strconv.Atoi(port)

	if err != nil {
		panic("Error parsing port")
	}

	responseDelayFactor := os.Getenv("RESPONSE_DELAY_FACTOR")
	if responseDelayFactor == "" {
		responseDelayFactor = "1"
	}

	responseDelayFactorInt, err := strconv.Atoi(responseDelayFactor)

	if err != nil {
		panic("Error parsing response delay factor")
	}

	return ServerConfig{
		Port:                portInt,
		ResponseDelayFactor: responseDelayFactorInt,
	}
}
