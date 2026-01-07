package services

import (
	"RealTimeChatApplication/models"
	"RealTimeChatApplication/repository"
)

type ChatRoomService struct {
	ChatRoomRepo *repository.ChatRoomRepo
}

func NewChatRoomService() *ChatRoomService {
	return &ChatRoomService{ChatRoomRepo: repository.NewChatRoomRepository(repository.DB)}
}

func (r *ChatRoomService) GetAllChat() ([]*models.ChatRoom, error) {

	return r.ChatRoomRepo.GetChatRoom()
}
