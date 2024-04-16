package packageCommander

import (
	"errors"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const DEFAULT_CURSOR_VALUE = 0

func prepareCursor(inputMessage *tgbotapi.Message) (int, error) {
	args := strings.Split(inputMessage.Text, " ")
	cursor := DEFAULT_CURSOR_VALUE

	if len(args) > 1 {
		converted, convErr := strconv.Atoi(args[1])
		
		if convErr != nil {
			return -1, errors.New("convert error")
		}

		cursor = converted
	}

	return cursor, nil
}

func (c *PackageCommander) List(inputMessage *tgbotapi.Message) {
	cursor, cursorErr := prepareCursor(inputMessage)

	if cursorErr != nil {
		log.Printf("Error: %v \nCommand: List", cursorErr.Error())
		return
	} 

	c.sendList(inputMessage.Chat.ID, cursor)
}