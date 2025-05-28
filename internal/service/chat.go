package service

import (
	"corpChat/internal/models"
	"corpChat/internal/repository"
	"errors"
)

type ChatService struct {
	chatRepo *repository.ChatRepository
}

func NewChatService(chatRepo *repositories.ChatRepository) *ChatService {
	return &ChatService{
		chatRepo: chatRepo,
	}
}

func (s *ChatService) CreateChat(name string, creatorID int) (*models.Chat, error) {
	chat := &models.Chat{
		Name:      name,
		CreatedBy: creatorID,
	}
	return s.chatRepo.Create(chat)
}

func (s *ChatService) InviteUser(chatID, userID, inviterID int) error {
	// Проверяем, что приглашающий — создатель чата
	chat, err := s.chatRepo.GetByID(chatID)
	if err != nil {
		return err
	}

	if chat.CreatedBy != inviterID {
		return errors.New("only chat creator can invite users")
	}

	return s.chatRepo.AddMember(chatID, userID, "member")
}
