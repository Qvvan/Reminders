package server

import (
	"Reminders/internal/database"
	"Reminders/internal/envs"
	"Reminders/internal/models"
	"go.uber.org/zap"
	"log"
)

var logger *zap.Logger

// InitLogger инициализирует логгер zap.
func InitLogger() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("Ошибка инициализации логгера: %v", err)
	}
}

// InitEnvs инициализирует переменные окружения.
func InitEnvs() error {
	err := envs.LoadEnvs()
	if err != nil {
		logger.Fatal("Ошибка инициализации ENV", zap.Error(err))
	}
	logger.Info("Инициализация ENV прошла успешно")
	return nil
}

// InitDatabase инициализирует подключение к базе данных.
func InitDatabase() error {
	err := database.InitDatabase()
	if err != nil {
		logger.Fatal("Ошибка подключения к базе данных", zap.Error(err))
	}
	logger.Info("Успешное подключение к базе данных")
	database.DB.AutoMigrate(&models.Reminder{})
	return nil
}

// InitServer инициализирует сервер, выполняя все необходимые шаги.
func InitServer() {
	InitLogger()

	if err := InitEnvs(); err != nil {
		logger.Fatal("Ошибка при инициализации окружения", zap.Error(err))
	}

	if err := InitDatabase(); err != nil {
		logger.Fatal("Ошибка при инициализации базы данных", zap.Error(err))
	}
}

// StartServer запускает сервер.
func StartServer() {
	// Инициализация роутов
	router := InitRotes()
	logger.Info("Сервер запущен")
	router.Run()
}

// GetLogger возвращает инициализированный логгер.
func GetLogger() *zap.Logger {
	return logger
}
