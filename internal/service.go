package internal

import (
	tg "github.com/Syfaro/telegram-bot-api"
)

type Implementation interface {
	StartBot() (chan []byte, chan error, error)
	StopBot() error
	SendMessage()
}

type service struct {
	tgApi          *tg.BotAPI
	updatesChannel tg.UpdatesChannel
	stopChan       chan bool
	alreadyStarted bool
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
