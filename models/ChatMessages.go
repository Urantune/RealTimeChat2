package models

import "time"

type ChatMessenger struct {
	ID        int       `json:"id"`
	RoomID    int       `json:"room_id" gorm:"column:room_id"`
	UserID    int       `json:"user_id" gorm:"column:user_id"`
	Content   string    `json:"content" gorm:"column:content"`
	CreatedAt time.Time `json:"created_at"`
}

func (ChatMessenger) TableName() string {
	return "chat_messages"
}
