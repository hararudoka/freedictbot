package main

import (
	"log"

	"github.com/hararudoka/freedictbot/internal/bot"
)

func main() {
	b, err := bot.New("bot.yml")
	if err != nil {
		log.Fatal(err)
	}

	b.Start()
}
