package server

import (
	"awesomeProject1/database"
	"awesomeProject1/envs"
	"awesomeProject1/models"
	"log"
)

func InitServer() {
	// Инициализация внешних значений ENV
	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		// Вывод сообщения об ошибке
		log.Fatal("Ошибка инициализации ENV: ", errEnvs)
	} else {
		log.Println("Инициализация ENV прошла успешно")
	}

	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatal("Ошибка подключения к базе данных: ", errDatabase)
	} else {
		log.Println("Успешное подключение к базе данных")
		database.DB.AutoMigrate(&models.Reminder{})
	}
}

func StartServer() {
	// Инициализация роутов
	router := InitRotes()
	router.Run()
}
