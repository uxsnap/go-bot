package packageCommander

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *PackageCommander) Get(inputMessage *tgbotapi.Message) {
	args := strings.Split(inputMessage.Text, " ")

	if len(args) < 2 {
		log.Printf("Not enough arguments! Command: Get")
		return  
	}

	id, convErr := strconv.Atoi(args[1])

	if convErr != nil {
		log.Printf("Error: %v \nCommand: Get", convErr.Error())
		return
	}

	packageItem, err := c.packageService.Get(id)

	if err != nil {
		log.Printf("Error: %v \nCommand: Get", err.Error())
		return
	} 

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, packageItem.String())

	c.bot.Send(msg)
}