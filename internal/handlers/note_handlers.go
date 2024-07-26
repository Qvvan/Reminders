package handlers

import (
	"Reminders/internal/database"
	"Reminders/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

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
	userID := ctx.Param("user_id")

	var reminders []models.Reminder

	// Поиск напоминаний по user_id
	result := database.DB.Where("user_id = ?", userID).Find(&reminders)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reminders"})
		return
	}

	// Проверяем, если напоминания не найдены
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No reminders found for the given user_id"})
		return
	}

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
	var reminders []models.Reminder

	// Получение всех напоминаний
	result := database.DB.Find(&reminders)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reminders"})
		return
	}

	// Проверяем, если напоминания не найдены
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No reminders found"})
		return
	}

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
	reminderID := ctx.Param("id")

	// Поиск напоминания по ID
	var reminder models.Reminder
	result := database.DB.First(&reminder, reminderID)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find reminder"})
		return
	}

	// Проверка, если напоминание не найдено
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No reminder found with the given ID"})
		return
	}

	// Проверка, если напоминание уже отправлено
	if reminder.IsSent {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Cannot delete reminder that has already been sent"})
		return
	}

	// Удаление напоминания по ID
	deleteResult := database.DB.Delete(&reminder)
	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete reminder"})
		return
	}

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
	reminderID := ctx.Param("id")
	var updatedReminder models.Reminder

	// Разбираем JSON из тела запроса
	if err := ctx.ShouldBindJSON(&updatedReminder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Поиск напоминания по ID
	var existingReminder models.Reminder
	result := database.DB.First(&existingReminder, reminderID)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find reminder"})
		return
	}

	// Проверка, если напоминание не найдено
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No reminder found with the given ID"})
		return
	}

	// Проверка, если напоминание уже отправлено
	if existingReminder.IsSent {
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update reminder"})
		return
	}

	// Отправляем успешный ответ
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
	var newReminder models.Reminder

	// Разбираем JSON из тела запроса
	if err := ctx.ShouldBindJSON(&newReminder); err != nil {
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reminder"})
		return
	}

	// Отправляем успешный ответ
	ctx.JSON(http.StatusCreated, gin.H{"message": "Reminder created successfully", "reminder": newReminder})
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
