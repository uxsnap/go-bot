package packageCommander

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *PackageCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	cursor, cursorErr := strconv.Atoi(callbackPath.CallbackData)

	if cursorErr != nil {
		log.Printf("Error: %v \nCommand: List", cursorErr.Error())
		return
	} 

	c.sendList(callback.Message.Chat.ID, cursor)
}