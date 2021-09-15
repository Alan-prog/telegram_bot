package internal

import (
	"fmt"
	tg "github.com/Syfaro/telegram-bot-api"
)

func (s *service) StartBot() {
	s.stopChan = make(chan bool)

	u := tg.NewUpdate(0)
	updates, err := s.tgApi.GetUpdatesChan(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	for update := range updates {
		select {
		case <-s.stopChan:
			return
		default:
			msg := tg.NewMessage(update.Message.Chat.ID, "some message")
			fmt.Println(update.Message.Chat.ID)
			s.tgApi.Send(msg)

		}

	}
}

func (s *service) StopBot() {
	s.stopChan <- true
}
