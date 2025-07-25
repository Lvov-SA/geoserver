package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_DATABASE   string
	USER_LOGIN    string
	USER_PASSWORD string
	APP_PORT      int
	HOST          string
}

var Configs Config

func Init() error {
	err := godotenv.Load("../.env")
	if err != nil {
		return err
	}
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		return err
	}
	Configs = Config{
		os.Getenv("DB_DATABASE"),
		os.Getenv("USER_LOGIN"),
		os.Getenv("USER_PASSWORD"),
		port,
		os.Getenv("HOST"),
	}
	return nil
}
