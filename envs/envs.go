package envs

import (
	"github.com/joho/godotenv"
	"os"
)

// Хранение данных значений ENV
var ServerEnvs Envs

// Структура для хранения значений ENV
type Envs struct {
	POSTGRES_PASSWORD string
	POSTGRES_USER     string
	POSTGRES_PORT     string
	POSTGRES_NAME     string
	POSTGRES_HOST     string
}

// / Инициализация значений ENV
func LoadEnvs() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	ServerEnvs.POSTGRES_USER = os.Getenv("DB_USER")
	ServerEnvs.POSTGRES_PASSWORD = os.Getenv("DB_PASSWORD")
	ServerEnvs.POSTGRES_PORT = os.Getenv("DB_PORT")
	ServerEnvs.POSTGRES_NAME = os.Getenv("DB_NAME")
	ServerEnvs.POSTGRES_HOST = os.Getenv("DB_HOST")

	return nil
}
