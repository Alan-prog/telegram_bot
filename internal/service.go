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
	tgApi    *tg.BotAPI
	stopChan chan bool
}

func NewService(token string) (Implementation, error) {
	var err error

	srv := service{}
	srv.tgApi, err = tg.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	go srv.StartBot()

	return &srv, err
}
