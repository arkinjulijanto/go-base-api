package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	APP_HOST    string
	APP_PORT    string
	API_VERSION string
	APP_URL     string
	APP_ENV     string

	DB_HOST       string
	DB_PORT       string
	DB_USER       string
	DB_PASS       string
	DB_NAME       string
	DB_CHARSET    string
	DB_PARSE_TIME string
	DB_LOCAL      string

	REDIS_HOST string
	REDIS_PORT string
	REDIS_PASS string

	JWT_EXPIRED string
	JWT_ISSUER  string
	JWT_SECRET  string

	LOG_STDOUT bool
	LOG_LEVEL  string
	LOG_PATH   string
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
}

func GetEnv() Config {
	logStdout, err := strconv.ParseBool(os.Getenv("LOG_STDOUT"))
	if err != nil {
		log.Println(err)
	}

	cfg := Config{
		APP_HOST:      os.Getenv("APP_HOST"),
		APP_PORT:      os.Getenv("APP_PORT"),
		API_VERSION:   os.Getenv("API_VERSION"),
		APP_URL:       os.Getenv("APP_URL"),
		APP_ENV:       os.Getenv("APP_ENV"),
		DB_HOST:       os.Getenv("DB_HOST"),
		DB_PORT:       os.Getenv("DB_PORT"),
		DB_USER:       os.Getenv("DB_USER"),
		DB_PASS:       os.Getenv("DB_PASS"),
		DB_NAME:       os.Getenv("DB_NAME"),
		DB_CHARSET:    os.Getenv("DB_CHARSET"),
		DB_PARSE_TIME: os.Getenv("DB_PARSE_TIME"),
		DB_LOCAL:      os.Getenv("DB_LOCAL"),
		REDIS_HOST:    os.Getenv("REDIS_HOST"),
		REDIS_PORT:    os.Getenv("REDIS_PORT"),
		REDIS_PASS:    os.Getenv("REDIS_PASS"),
		JWT_EXPIRED:   os.Getenv("JWT_EXPIRED"),
		JWT_ISSUER:    os.Getenv("JWT_ISSUER"),
		JWT_SECRET:    os.Getenv("JWT_SECRET"),
		LOG_STDOUT:    logStdout,
		LOG_LEVEL:     os.Getenv("LOG_LEVEL"),
		LOG_PATH:      os.Getenv("LOG_PATH"),
	}

	return cfg
}
