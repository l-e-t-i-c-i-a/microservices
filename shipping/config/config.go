package config

import (
	"log"
	"os"
	"strconv"
)

func GetEnv() string {
	return getEnvironmentValue("ENV")
}

// sem o GetDataSourceURL pois Shipping não usa banco de dados

func GetApplicationPort() int {
	portStr := getEnvironmentValue("APPLICATION_PORT")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}

	return port
}

func getEnvironmentValue(key string) string {
	if os.Getenv(key) == "" {
		// Se a variável não existir, avisa e para o programa
		log.Fatalf("%s environment variable is missing", key)
	}
	return os.Getenv(key)
}