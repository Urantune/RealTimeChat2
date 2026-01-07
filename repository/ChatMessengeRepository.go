package repository

import "gorm.io/gorm"

type ChatMessageRepo struct {
	db *gorm.DB
}

func NewChatMessageRepository(db *gorm.DB) *ChatMessageRepo {
	return &ChatMessageRepo{db: db}
}
