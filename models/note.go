package models

import (
	"time"
)

type Reminder struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id"`
	Message   string    `json:"message"`
	SendAt    time.Time `json:"send_at"`
	IsSent    bool      `json:"is_sent"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
