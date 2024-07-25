package server

import (
	_ "awesomeProject1/docs" // импортируйте ваш файл docs.go
	"awesomeProject1/handlers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // swagger middleware
)

func InitRotes() *gin.Engine {
	router := gin.Default()

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Создание заметки
	router.POST("/reminders", handlers.CreateMessageHandler)
	// Удаление заметки
	router.DELETE("/reminders/:id", handlers.DeleteMessageHandler)
	// Получение заметки
	router.GET("/reminders/:user_id", handlers.GetMessageByUserIDHandler)
	// Редактирование заметки
	router.PUT("/reminders/:id", handlers.UpdateMessageHandler)
	// Получение списка всех заметок
	router.GET("/reminders", handlers.GetAllMessagesHandler)

	return router
}
