package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = initConfig()

func initConfig() Config {

	godotenv.Load()

	return Config{
		Port: getEnv("PORT", "8080"),

		DBUser: getEnv("DB_USER", "root"),

		DBPassword: getEnv("DB_PASSWORD", "mypassword"),

		DBAddress: fmt.Sprintf(
			"%s:%s",
			getEnv("DB_HOST", "127.0.0.1"),
			getEnv("DB_PORT", "3306"),
		),

		DBName: getEnv("DB_NAME", "sms"),

		JWTSecret: getEnv(
			"JWT_SECRET",
			"super-secret-key",
		),

		JWTExpirationInSeconds: getEnvAsInt(
			"JWT_EXPIRATION_IN_SECONDS",
			86400,
		),
	}
}

func getEnv(key string, fallback string) string {

	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return value
}

func getEnvAsInt(key string, fallback int64) int64 {

	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	intValue, err := strconv.ParseInt(
		value,
		10,
		64,
	)

	if err != nil {
		return fallback
	}

	return intValue
}