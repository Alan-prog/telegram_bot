package internal

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func (s *service) SendMessage() {
	u := tgbotapi.NewUpdate(0)
	updates, err := s.tgApi.GetUpdatesChan(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	for update := range updates {
		fmt.Println(update.Message.Chat.ID)
		msg := tgbotapi.NewMessage(496869421, "some message")
		s.tgApi.Send(msg)
	}
}
