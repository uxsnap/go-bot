package packageCommander

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *PackageCommander) Get(inputMessage *tgbotapi.Message) {
	text := "Here're the available commands: \n\n"

	for _, v := range AVAILABLE_COMMANDS {
		text += fmt.Sprintf("/%v\n", v)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)

	sentMsg, err := c.bot.Send(msg)

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Printf("Help command invoked for chat %v with message %v", inputMessage.Chat.Title, sentMsg.Text)
	}
}