package services

import (
	"encoding/json"
	"fmt"
	"time"

	"RealTimeChatApplication/models"
	"RealTimeChatApplication/repository"
)

type ChatMessengeService struct {
	ChatRepo *repository.ChatMessageRepo
}

func NewChatMessengeService() *ChatMessengeService {
	return &ChatMessengeService{ChatRepo: repository.NewChatMessageRepository(repository.DB)}
}

func (r *ChatMessengeService) GetChatByRoomId(roomId int) ([]*models.ChatMessenger, error) {
	return r.ChatRepo.GetChatMessageByChatRoom(roomId)
}

func (r *ChatMessengeService) GetChatByRoomIdCached(roomId int) ([]*models.ChatMessenger, error) {
	key := fmt.Sprintf("chat:room:%d", roomId)

	if b, err := repository.GetJSON(key); err == nil {
		var cached []*models.ChatMessenger
		if json.Unmarshal(b, &cached) == nil {
			return cached, nil
		}
	}

	listChat, err := r.GetChatByRoomId(roomId)
	if err != nil {
		return nil, err
	}

	if b, err := json.Marshal(listChat); err == nil {
		_ = repository.SetJSON(key, b, 15*time.Second)
	}

	return listChat, nil
}

func (r *ChatMessengeService) InvalidateRoomCache(roomId int) {
	key := fmt.Sprintf("chat:room:%d", roomId)
	_ = repository.Del(key)
}

func (r *ChatMessengeService) AddMessage(roomId int, userId int, content string) error {
	msg := &models.ChatMessenger{
		RoomID:    roomId,
		UserID:    userId,
		Content:   content,
		CreatedAt: time.Now(),
	}
	if err := r.ChatRepo.Create(msg); err != nil {
		return err
	}
	r.InvalidateRoomCache(roomId)
	return nil
}
