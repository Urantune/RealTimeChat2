package repository

import (
	"RealTimeChatApplication/models"

	"gorm.io/gorm"
)

type ChatMessageRepo struct {
	db *gorm.DB
}

func NewChatMessageRepository(db *gorm.DB) *ChatMessageRepo {
	return &ChatMessageRepo{db: db}
}

func (r *ChatMessageRepo) GetChatMessageByChatRoom(chatRoomId int) ([]*models.ChatMessenger, error) {
	var chatMessages []*models.ChatMessenger
	err := r.db.Where("room_id = ?", chatRoomId).Find(&chatMessages).Error
	return chatMessages, err
}

func (r *ChatMessageRepo) Create(msg *models.ChatMessenger) error {
	return r.db.Create(msg).Error
}
