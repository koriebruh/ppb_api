package conf

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Server struct {
	Host string
	Port string
}

type DataBase struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

type Config struct {
	Server
	DataBase
}

// <-- CONSTRUCTOR --> //

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using default static configuration")
	}

	// Return Config with default fallback values
	return &Config{
		Server: Server{
			Host: getEnv("SERVER_HOST", "localhost"),
			Port: getEnv("SERVER_PORT", "3000"),
		},
		DataBase: DataBase{
			User: getEnv("DB_USER", "admin"),
			Pass: getEnv("DB_PASS", "korie123"),
			Host: getEnv("DB_HOST", "13.210.50.68"),
			Port: getEnv("DB_PORT", "3306"),
			Name: getEnv("DB_NAME", "uas_ppb"),
		},
	}

}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
