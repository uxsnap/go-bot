package packageCommander

import (
	"fmt"
	"log"
	"math"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/logistic"
)

const ITEMS_LIMIT = 5

func prepareKeyboard(msg *tgbotapi.MessageConfig, itemsLen int) {
	rowCount := math.Ceil(float64(itemsLen) / ITEMS_LIMIT)
	var rows []tgbotapi.InlineKeyboardButton
	
	for rowInd := 0; rowInd < int(rowCount); rowInd++ {
		rows = append(rows, tgbotapi.NewInlineKeyboardButtonData(
			fmt.Sprintf("%v", rowInd + 1), 
			fmt.Sprintf("logistic_package_list_%v", rowInd),
		))
	}

	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(rows...),
	)

	msg.ReplyMarkup = numericKeyboard
}

func listString(packageItems []logistic.Package) string {
	text := "Here's the items: \n\n"

	for _, item := range packageItems {
		text += item.String() + "\n"
	}

	return text
}

func (c *PackageCommander) sendList(chatID int64, cursor int) {
	packageItems, err := c.packageService.List(uint64(cursor * ITEMS_LIMIT), ITEMS_LIMIT)
	allItemsCount, countErr := c.packageService.Count()

	if err != nil {
		log.Printf("Error: %v \nCommand: List", err.Error())
		return
	} 

	if countErr != nil {
		log.Printf("Error: %v \nCommand: List", countErr.Error())
		return
	} 

	msg := tgbotapi.NewMessage(chatID, listString(packageItems))

	prepareKeyboard(&msg, allItemsCount)

	if _, err := c.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}