package repository

import (
	"RealTimeChatApplication/models"

	"gorm.io/gorm"
)

type ChatRoomRepo struct {
	db *gorm.DB
}

func NewChatRoomRepository(db *gorm.DB) *ChatRoomRepo {
	return &ChatRoomRepo{db: db}
}

func (r *ChatRoomRepo) GetChatRoom() ([]*models.ChatRoom, error) {
	var chatRooms []*models.ChatRoom
	err := r.db.Find(&chatRooms).Error
	return chatRooms, err
}
