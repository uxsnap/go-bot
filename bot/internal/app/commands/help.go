package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Help(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, `
	/help - help
/list - list products`)

	c.bot.Send(msg)
}

func init() {
	registeredCommands["help"] = (*Commander).Help;
}