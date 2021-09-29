package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	tg "github.com/Syfaro/telegram-bot-api"
)

func (s *service) StartBot() (chan []byte, chan error, error) {
	var (
		responseChan = make(chan []byte)
		errorChan    = make(chan error)
	)

	if s.alreadyStarted {
		return nil, nil, errors.New("the chan is already working, you cant make several channels")
	}

	s.stopChan = make(chan bool)

	if s.updatesChannel == nil {
		var err error
		u := tg.NewUpdate(0)
		s.updatesChannel, err = s.tgApi.GetUpdatesChan(u)
		if err != nil {
			return nil, nil, fmt.Errorf("error while creating updates chan: %v", err)
		}
	}

	go func(infoChan chan []byte, errChan chan error) {
		for update := range s.updatesChannel {
			select {
			case <-s.stopChan:
				s.alreadyStarted = false
				close(infoChan)
				close(errChan)
				return
			default:
				var err error
				updateByte, err := json.Marshal(&update)
				if err != nil {
					errChan <- fmt.Errorf("json marshal error: %v", err)
				} else {
					infoChan <- updateByte
				}
			}
		}
	}(responseChan, errorChan)
	s.alreadyStarted = true
	return responseChan, errorChan, nil
}

func (s *service) StopBot() error {
	if !s.alreadyStarted {
		return errors.New("was already stopped")
	}
	s.stopChan <- true
	return nil
}
