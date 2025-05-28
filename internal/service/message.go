package service

import (
	"corpChat/internal/models"
	"corpChat/internal/repository"
)

type MessageService struct {
	messageRepo *repository.MessageRepository
}

func NewMessageService(messageRepo *repository.MessageRepository) *MessageService {
	return &MessageService{
		messageRepo: messageRepo,
	}
}

func (s *MessageService) SendMessage(chatID, userID int, content string) error {
	msg := &models.Message{
		ChatID:  chatID,
		UserID:  userID,
		Content: content,
	}
	return s.messageRepo.Create(msg)
}

func (s *MessageService) GetMessages(chatID, limit int) ([]models.Message, error) {
	return s.messageRepo.GetByChatID(chatID, limit)
}
