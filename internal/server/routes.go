package server

import (
	_ "Reminders/internal/docs"
	"Reminders/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRotes() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Получение напоминания
	router.GET("/reminders/:user_id", handlers.GetMessageByUserIDHandler)
	// Получение списка всех напоминаний
	router.GET("/reminders", handlers.GetAllMessagesHandler)
	// Создание напоминания
	router.POST("/reminders", handlers.CreateMessageHandler)
	// Редактирование напоминания
	router.PUT("/reminders/:id", handlers.UpdateMessageHandler)
	// Удаление напоминания
	router.DELETE("/reminders/:id", handlers.DeleteMessageHandler)

	return router
}
