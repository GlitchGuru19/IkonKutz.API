package initializers

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	DBURL             string
	JWTSecret         string
	JWTExpiresInHours int
	ClientOrigin      string
	GinMode           string

	AdminName     string
	AdminEmail    string
	AdminPassword string
}

var AppConfig Config

func LoadConfig() {
	jwtHours := 24

	if raw := os.Getenv("JWT_EXPIRES_IN_HOURS"); raw != "" {
		parsed, err := strconv.Atoi(raw)
		if err != nil || parsed <= 0 {
			log.Fatal("JWT_EXPIRES_IN_HOURS must be a positive integer")
		}
		jwtHours = parsed
	}

	AppConfig = Config{
		DBURL:             os.Getenv("URI"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		JWTExpiresInHours: jwtHours,
		ClientOrigin:      os.Getenv("CLIENT_ORIGIN"),
		GinMode:           os.Getenv("GIN_MODE"),

		AdminName:     os.Getenv("ADMIN_NAME"),
		AdminEmail:    os.Getenv("ADMIN_EMAIL"),
		AdminPassword: os.Getenv("ADMIN_PASSWORD"),
	}

	if AppConfig.DBURL == "" {
		log.Fatal("DB_URL is required")
	}

	if AppConfig.JWTSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	if AppConfig.ClientOrigin == "" {
		log.Fatal("CLIENT_ORIGIN is required")
	}
}
