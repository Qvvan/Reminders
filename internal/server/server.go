package server

import (
	"Reminders/internal/database"
	"Reminders/internal/envs"
	"Reminders/internal/models"
	"log"
)

// InitEnvs инициализирует переменные окружения.
func InitEnvs() error {
	err := envs.LoadEnvs()
	if err != nil {
		log.Fatal("Ошибка инициализации ENV: ", err)
	}
	log.Println("Инициализация ENV прошла успешно")
	return nil
}

// InitDatabase инициализирует подключение к базе данных.
func InitDatabase() error {
	err := database.InitDatabase()
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных: ", err)
	}
	log.Println("Успешное подключение к базе данных")
	database.DB.AutoMigrate(&models.Reminder{})
	return nil
}

// InitServer инициализирует сервер, выполняя все необходимые шаги.
func InitServer() {
	if err := InitEnvs(); err != nil {
		log.Fatal(err)
	}

	if err := InitDatabase(); err != nil {
		log.Fatal(err)
	}
}

// StartServer запускает сервер.
func StartServer() {
	// Инициализация роутов
	router := InitRotes()
	router.Run()
}
