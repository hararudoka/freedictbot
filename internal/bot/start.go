package bot

import (
	tele "gopkg.in/telebot.v3"
)

func (b Bot) onStart(c tele.Context) error {
	return c.Send(b.Text(c, "start", c.Sender()))
}

// are you ill?
// answer: https://www.youtube.com/watch?v=8zKjg-X9dzY