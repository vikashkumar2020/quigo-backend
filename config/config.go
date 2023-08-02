package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port              string
	ServerApiPrefixV1 string
	JwtSecret         string
	BasePath          string
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
}

type JWTConfig struct {
	AccessTokenPrivateKey string
}

type EmailConfig struct {
	EmailFrom string
	EmailHost string
	EmailPort string
	EmailUser string
	EmailPass string
}

// NewServerConfig returns a pointer to a new ServerConfig struct initialized with values from environment variables.
func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:              os.Getenv("APP_PORT"),
		ServerApiPrefixV1: os.Getenv("SERVER_API_PREFIX_V1"),
		JwtSecret:         os.Getenv("JWT_SECRET"),
		BasePath:          os.Getenv("SERVER_BASE_PATH"),
	}
}

// NewDBConfig returns a pointer to a new DBConfig struct initialized with values from environment variables.
func NewDBConfig() *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
	}
}

func NewJWTConfig() *JWTConfig {
	return &JWTConfig{
		AccessTokenPrivateKey: os.Getenv("TOKEN_SECRET"),
	}
}

func NewEmailConfig() *EmailConfig {
	return &EmailConfig{
		EmailFrom: os.Getenv("EMAIL_FROM"),
		EmailHost: os.Getenv("SMTP_HOST"),
		EmailPort: os.Getenv("SMTP_PORT"),
		EmailUser: os.Getenv("SMTP_USER"),
		EmailPass: os.Getenv("SMTP_PASS"),
	}
}

// LoadEnv loads environment variables from the .env file in the current directory.
func LoadEnv() {

	loadEnvError := godotenv.Load(".env")
	if loadEnvError != nil {
		LogFatal(loadEnvError)
	}
}
