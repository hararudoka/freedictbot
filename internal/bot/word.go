package bot

import (
	"github.com/hararudoka/freedictbot/internal/bot/middle"
	tele "gopkg.in/telebot.v3"
)

func (b Bot) onWord(c tele.Context) error {
	return c.Send(
		middle.GenerateMessage(c.Args()[0]),
	)
}
