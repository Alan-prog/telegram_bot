package internal

import (
	"fmt"
	tg "github.com/Syfaro/telegram-bot-api"
)

func (s *service) StartBot() {
	u := tg.NewUpdate(0)
	updates, err := s.tgApi.GetUpdatesChan(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	for update := range updates {
		msg := tg.NewMessage(update.Message.Chat.ID, "some message")
		s.tgApi.Send(msg)
	}
}

func (s *service) StopBot() {
	u := tg.NewUpdate(0)
	updates, err := s.tgApi.GetUpdatesChan(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	for update := range updates {
		msg := tg.NewMessage(update.Message.Chat.ID, "some message")
		s.tgApi.Send(msg)
	}
}
