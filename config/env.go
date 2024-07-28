package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// Musr      = "admin"
	// Mpwd      = "password"
	// Mhost     = "mongo"
	// Mport     = "27017"
	PublicHost             string
	Port                   string
	MDBName                string
	MDBURI                 string
	JWTExpirationInSeconds int64
	JWTRefreshExpirationInSeconds int64
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		// Musr      = "admin"
		// Mpwd      = "password"
		// Mhost     = "mongo"
		// MPort     = "27017"
		Port:                   getEnv("POST", "8080"),
		MDBName:                getEnv("DB_NAME", "didis-comp-bk-api"),
		MDBURI:                 getEnv("DB_URI", "mongodb://localhost:27018,localhost:27019,localhost:27020/admin?authSource=admin&replicaSet=didi"),
		JWTExpirationInSeconds: getEnvsAsInt("JWT_EXP", 3600*24),
		JWTRefreshExpirationInSeconds: getEnvsAsInt("JWT_EXP", 3600*24*7*4),
		JWTSecret:              getEnv("JWT_SECRET", "non-secret-secret-anymore?"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvsAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
