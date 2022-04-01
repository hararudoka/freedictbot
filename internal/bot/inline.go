package bot

import (
	"strings"

	"github.com/hararudoka/freedictbot/internal/bot/middle"
	tele "gopkg.in/telebot.v3"
)

func (b Bot) onInlineWord(c tele.Context, a string) error {
	results := tele.Results{}

	args := strings.Split(a, " ")

	{ // here we can add some logic on args, etc. error handling
		res := b.Result(c, "word", map[string]interface{}{
			"Title": "title",
		})
		res.SetContent(&tele.InputTextMessageContent{Text: middle.GenerateMessage(a)}) // TODO: errors
		_ = args
		results = append(results, res)
	}

	return c.Answer(&tele.QueryResponse{
		Results:   results,
		CacheTime: -1,
	})
}

// you are warm
// answer:
