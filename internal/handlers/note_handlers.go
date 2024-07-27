package handlers

import (
	"Reminders/internal/database"
	"Reminders/internal/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// Глобальная переменная логгера
var logger *zap.Logger

// SetLogger устанавливает логгер для использования в обработчиках.
func SetLogger(l *zap.Logger) {
	logger = l
}

// GetMessageByUserIDHandler godoc
// @Summary Поиск напоминаний по user_id
// @Description Получить все напоминания для конкретного пользователя
// @Tags reminders
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} models.Reminder
// @Failure 404 {object} ErrorResponse
// @Router /reminders/{user_id} [get]
func GetMessageByUserIDHandler(ctx *gin.Context) {
	start := time.Now()
	userID := ctx.Param("user_id")

	var reminders []models.Reminder

	// Поиск напоминаний по user_id
	result := database.DB.Where("user_id = ?", userID).Find(&reminders)
	if result.Error != nil {
		logger.Error("Failed to fetch reminders", zap.String("user_id", userID), zap.Error(result.Error))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reminders"})
		return
	}

	// Проверяем, если напоминания не найдены
	if result.RowsAffected == 0 {
		logger.Info("No reminders found for user_id", zap.String("user_id", userID))
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No reminders found for the given user_id"})
		return
	}

	logger.Info("Reminders fetched successfully", zap.String("user_id", userID), zap.Duration("elapsed_time", time.Since(start)))
	ctx.JSON(http.StatusOK, gin.H{"reminders": reminders})
}

// GetAllMessagesHandler godoc
// @Summary Получение всех напоминаний
// @Description Получить список всех напоминаний
// @Tags reminders
// @Accept json
// @Produce json
// @Success 200 {array} models.Reminder
// @Failure 500 {object} ErrorResponse
// @Router /reminders [get]
func GetAllMessagesHandler(ctx *gin.Context) {
	start := time.Now()
	var reminders []models.Reminder

	// Получение всех напоминаний
	result := database.DB.Find(&reminders)
	if result.Error != nil {
		logger.Error("Failed to fetch reminders", zap.Error(result.Error))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reminders"})
		return
	}

	// Проверяем, если напоминания не найдены
	if result.RowsAffected == 0 {
		logger.Info("No reminders found")
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No reminders found"})
		return
	}

	logger.Info("All reminders fetched successfully", zap.Duration("elapsed_time", time.Since(start)))
	ctx.JSON(http.StatusOK, gin.H{"reminders": reminders})
}

// DeleteMessageHandler godoc
// @Summary Удалить напоминание
// @Description Удалить напоминание по идентификатору, если оно не было отправлено
// @Tags reminders
// @Accept json
// @Produce json
// @Param id path int true "Reminder ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /reminders/{id} [delete]
func DeleteMessageHandler(ctx *gin.Context) {
	start := time.Now()
	reminderID := ctx.Param("id")

	// Поиск напоминания по ID
	var reminder models.Reminder
	result := database.DB.First(&reminder, reminderID)
	if result.Error != nil {
		logger.Error("Failed to find reminder", zap.String("reminder_id", reminderID), zap.Error(result.Error))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find reminder"})
		return
	}

	// Проверка, если напоминание не найдено
	if result.RowsAffected == 0 {
		logger.Info("No reminder found with the given ID", zap.String("reminder_id", reminderID))
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No reminder found with the given ID"})
		return
	}

	// Проверка, если напоминание уже отправлено
	if reminder.IsSent {
		logger.Info("Cannot delete reminder that has already been sent", zap.String("reminder_id", reminderID))
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Cannot delete reminder that has already been sent"})
		return
	}

	// Удаление напоминания по ID
	deleteResult := database.DB.Delete(&reminder)
	if deleteResult.Error != nil {
		logger.Error("Failed to delete reminder", zap.String("reminder_id", reminderID), zap.Error(deleteResult.Error))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete reminder"})
		return
	}

	logger.Info("Reminder deleted successfully", zap.String("reminder_id", reminderID), zap.Duration("elapsed_time", time.Since(start)))
	ctx.JSON(http.StatusOK, gin.H{"message": "Reminder deleted successfully"})
}

// UpdateMessageHandler godoc
// @Summary Обновление существующего напоминания
// @Description Обновить напоминание с указанным идентификатором
// @Tags reminders
// @Accept json
// @Produce json
// @Param id path int true "Reminder ID"
// @Param reminder body models.Reminder true "Updated reminder object"
// @Success 200 {object} models.Reminder
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /reminders/{id} [put]
func UpdateMessageHandler(ctx *gin.Context) {
	start := time.Now()
	reminderID := ctx.Param("id")
	var updatedReminder models.Reminder

	// Разбираем JSON из тела запроса
	if err := ctx.ShouldBindJSON(&updatedReminder); err != nil {
		logger.Error("Invalid request data", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Поиск напоминания по ID
	var existingReminder models.Reminder
	result := database.DB.First(&existingReminder, reminderID)
	if result.Error != nil {
		logger.Error("Failed to find reminder", zap.String("reminder_id", reminderID), zap.Error(result.Error))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find reminder"})
		return
	}

	// Проверка, если напоминание не найдено
	if result.RowsAffected == 0 {
		logger.Info("No reminder found with the given ID", zap.String("reminder_id", reminderID))
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No reminder found with the given ID"})
		return
	}

	// Проверка, если напоминание уже отправлено
	if existingReminder.IsSent {
		logger.Info("Cannot update reminder that has already been sent", zap.String("reminder_id", reminderID))
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Cannot update reminder that has already been sent"})
		return
	}

	// Обновление полей напоминания
	existingReminder.UserID = updatedReminder.UserID
	existingReminder.Message = updatedReminder.Message
	existingReminder.SendAt = updatedReminder.SendAt
	existingReminder.UpdatedAt = time.Now()

	// Сохранение обновленного напоминания в базе данных
	saveResult := database.DB.Save(&existingReminder)
	if saveResult.Error != nil {
		logger.Error("Failed to update reminder", zap.String("reminder_id", reminderID), zap.Error(saveResult.Error))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update reminder"})
		return
	}

	logger.Info("Reminder updated successfully", zap.String("reminder_id", reminderID), zap.Duration("elapsed_time", time.Since(start)), zap.Any("updated_reminder", existingReminder))
	ctx.JSON(http.StatusOK, gin.H{"message": "Reminder updated successfully", "reminder": existingReminder})
}

// CreateMessageHandler godoc
// @Summary Создать напоминание
// @Description Создать новое напоминание с предоставленными деталями
// @Tags reminders
// @Accept json
// @Produce json
// @Param reminder body models.Reminder true "Reminder object"
// @Success 200 {object} models.Reminder
// @Failure 400 {object} ErrorResponse
// @Router /reminders [post]
func CreateMessageHandler(ctx *gin.Context) {
	start := time.Now()
	var newReminder models.Reminder

	// Разбираем JSON из тела запроса
	if err := ctx.ShouldBindJSON(&newReminder); err != nil {
		logger.Error("Invalid request data", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Устанавливаем значения по умолчанию
	newReminder.CreatedAt = time.Now()
	newReminder.UpdatedAt = time.Now()
	newReminder.IsSent = false

	// Сохраняем напоминание в базе данных
	result := database.DB.Create(&newReminder)
	if result.Error != nil {
		logger.Error("Failed to create reminder", zap.Error(result.Error))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reminder"})
		return
	}

	logger.Info("Reminder created successfully", zap.Duration("elapsed_time", time.Since(start)), zap.Any("new_reminder", newReminder))
	ctx.JSON(http.StatusCreated, gin.H{"message": "Reminder created successfully", "reminder": newReminder})
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
