package database

import (
	"awesomeProject1/envs"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// / Объявление переменной DB, хранящей ссылку на экземпляр базы данных
var DB *gorm.DB

// / Функция инициализации базы данных
func InitDatabase() error {
	env := envs.ServerEnvs

	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		env.POSTGRES_HOST, env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_NAME, env.POSTGRES_PORT)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return err
	} else {
		DB = db
		return nil
	}
}
