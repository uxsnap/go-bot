package packageCommander

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *PackageCommander) Delete(inputMessage *tgbotapi.Message) {
	args := strings.Split(inputMessage.Text, " ")

	if len(args) < 2 {
		log.Printf("Not enough arguments! Command: Delete")
		return  
	}

	id, convErr := strconv.Atoi(args[1])

	if convErr != nil {
		log.Printf("Error: %v \nCommand: Delete", convErr.Error())
		return
	}

	if isOutOfBoundaries(id) {
		log.Println("Out of boundaries, \nCommand: Get")
		return
	}

	removed, err := c.packageService.Remove(uint64(id))

	if err != nil || !removed {
		log.Printf("Error: %v \nCommand: Delete", err.Error())
		return
	} 

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Successfully removed: %v", id))

	c.bot.Send(msg)
}