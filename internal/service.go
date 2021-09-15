package internal

import (
	tg "github.com/Syfaro/telegram-bot-api"
)

type Implementation interface {
	SendMessage()
}

type service struct {
	tgApi *tg.BotAPI
}

func NewService(token string) (Implementation, error) {
	var err error

	srv := service{}
	srv.tgApi, err = tg.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &srv, err
}
