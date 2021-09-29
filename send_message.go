package telegram_bot

import (
	"fmt"
	tg "github.com/Syfaro/telegram-bot-api"
)

func (s *service) SendMessage(message string, chatID int64) error {
	tgMsg := tg.NewMessage(chatID, message)
	_, err := s.tgApi.Send(tgMsg)
	if err != nil {
		return fmt.Errorf("error while sending message: %v", err)
	}
	return nil
}
