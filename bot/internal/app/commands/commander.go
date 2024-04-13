package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/uxsnap/go-bot/internal/service/product"
)

var registeredCommands = map[string]func(c *Commander, message *tgbotapi.Message){};

type Commander struct {
	bot *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{ bot: bot,  productService: productService }
}

func (c *Commander) HandleUpdate(update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	command, ok := registeredCommands[update.Message.Command()] 

	if ok {
		command(c, update.Message)
	} else {
		c.Default(update.Message)
	}
}