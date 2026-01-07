package services

import (
	"RealTimeChatApplication/models"
	"RealTimeChatApplication/repository"
)

type UserService struct {
	UserRepo *repository.UserRepo
}

func NewUserService() *UserService {
	return &UserService{UserRepo: repository.NewUserRepository(repository.DB)}
}

func (s *UserService) GetUserByUserName(username string) (*models.User, error) {
	return s.UserRepo.GetUserByUserName(username)
}
