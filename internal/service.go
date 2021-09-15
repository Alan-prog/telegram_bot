package internal

import (
	tg "github.com/Syfaro/telegram-bot-api"
)

type Implementation interface {
	StartBot()
	StopBot()
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

	go func() {
		srv.SendMessage()
	}()

	return &srv, err
}
