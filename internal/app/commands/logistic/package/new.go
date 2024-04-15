package packageCommander

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/logistic"
)

func (c *PackageCommander) New(inputMessage *tgbotapi.Message) {
	args := strings.Split(inputMessage.Text, " ")

	if len(args) < 2 {
		log.Printf("Not enough arguments! Command: New")
		return  
	}

	name := args[1]
	newPackage := logistic.Package{ Name: name }

	id, err := c.packageService.Create(newPackage)

	if err != nil {
		log.Printf("Error: %v \nCommand: New", err.Error())
		return
	} 

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Newly created package with id: %v", id))

	c.bot.Send(msg)
}