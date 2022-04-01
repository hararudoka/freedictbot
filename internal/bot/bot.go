package bot

import (
	"log"
	"strings"

	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
	"gopkg.in/telebot.v3/middleware"
)

type Bot struct {
	*tele.Bot
	*layout.Layout
}

func New(path string) (*Bot, error) {
	lt, err := layout.New(path)
	if err != nil {
		return nil, err
	}

	b, err := tele.NewBot(lt.Settings())
	if err != nil {
		return nil, err
	}

	if cmds := lt.Commands(); cmds != nil {
		if err := b.SetCommands(cmds); err != nil {
			return nil, err
		}
	}

	return &Bot{
		Bot:    b,
		Layout: lt,
	}, nil
}

func (b Bot) onQuery(c tele.Context) error {
	command, data := parseQuery(c.Data())

	log.Println("got inline rn")
	switch command {
	case "w", "word":
		return b.onInlineWord(c, data)
	default:
		return b.onHelp(c)
	}
}

func parseQuery(text string) (command, data string) {
	parts := strings.Split(text, " ")
	if len(parts) >= 1 {
		command = strings.ToLower(parts[0])
		data = strings.Join(parts[1:], " ")
	}
	return
}

func (b *Bot) Start() {
	// Middlewares
	b.Use(middleware.Logger())
	b.Use(middleware.AutoRespond())
	b.Use(b.Middleware("en"))

	// Handlers
	b.Handle("/start", b.onStart)
	b.Handle("/word", b.onWord)

	b.Handle(tele.OnQuery, b.onQuery)

	b.Bot.Start()
}
