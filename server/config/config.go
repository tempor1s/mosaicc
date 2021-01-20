package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// init the environment
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Did not load variables from .env file. This is normal if you are in CI/CD or production.")
	}
}

// Config represents the different variables that are needed for configuration of the application
type Config struct {
	Auth0Domain  string // auth0 api domain
	Auth0ID      string // auth0 api ID
	AppEnv       string // the environment that the application is running in (env, prod, etc)
	DbURI        string // database URI
	DbUsername   string // database username
	DbPassword   string // database password
	GCPProjectID string // google cloud platform project ID
}

// GetConfig will return the current config
func GetConfig() *Config {
	config := &Config{
		AppEnv:       os.Getenv("APP_ENV"),
		Auth0Domain:  os.Getenv("AUTH0_DOMAIN"),
		Auth0ID:      os.Getenv("AUTH0_API_ID"),
		DbURI:        os.Getenv("DB_URI"),
		DbPassword:   os.Getenv("DB_PASSWORD"),
		DbUsername:   os.Getenv("DB_USER"),
		GCPProjectID: os.Getenv("GCP_PROJECT_ID"),
	}

	return config
}
