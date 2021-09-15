package main

import (
	"fmt"
	"github.com/telegram_bot/internal"
	"log"
)

func main() {
	tgImpl, err := internal.NewService("1991633298:AAEs5NDiuT2hSPZ45pZ6qdWrr1kquR39sFE")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tgImpl)
	tgImpl.SendMessage()
}
