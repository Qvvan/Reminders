package main

import (
	"Reminders/internal/database"
	"Reminders/internal/models"
	"Reminders/internal/server"
	"go.uber.org/zap"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const myTelegramID = 323993202

var logger *zap.Logger

func init() {
	// Инициализация логгера
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Panic("Ошибка инициализации логгера", zap.Error(err))
	}
}

func main() {
	server.InitServer()

	// Инициализация бота
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		logger.Panic("Ошибка инициализации Telegram бота", zap.Error(err))
	}

	for {
		checkAndSendReminders(bot)
		time.Sleep(1 * time.Minute)
	}
}

// checkAndSendReminders проверяет напоминания и отправляет их, если время отправки прошло.
func checkAndSendReminders(bot *tgbotapi.BotAPI) {
	var reminders []models.Reminder
	now := time.Now()

	// Запрос на получение напоминаний, которые ещё не отправлены и время отправки которых прошло
	if err := database.DB.Where("is_sent = ? AND send_at <= ?", false, now).Find(&reminders).Error; err != nil {
		logger.Error("Ошибка при получении напоминаний", zap.Error(err))
		return
	}

	for _, r := range reminders {
		if sendReminder(bot, r) {
			// Обновление статуса напоминания в базе данных
			if err := database.DB.Model(&r).Update("is_sent", true).Error; err != nil {
				logger.Error("Ошибка при обновлении статуса напоминания", zap.Int("reminder_id", r.ID), zap.Error(err))
			}
		}
	}
}

// sendReminder отправляет напоминание пользователю в Telegram.
func sendReminder(bot *tgbotapi.BotAPI, r models.Reminder) bool {
	msg := tgbotapi.NewMessage(myTelegramID, r.Message)
	_, err := bot.Send(msg)
	if err != nil {
		logger.Error("Не удалось отправить сообщение", zap.Int("reminder_id", r.ID), zap.Error(err))
		return false
	}
	logger.Info("Сообщение успешно отправлено", zap.Int("reminder_id", r.ID))
	return true
}
