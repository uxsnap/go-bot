package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/uxsnap/go-bot/internal/service/product"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	productService := product.NewService()

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			helpCommand(bot, &update)
		case "list":
			listCommand(bot, &update, productService)
		default:
			defaultBehavior(bot, &update)
		}


		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}

func helpCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, `
	/help - help
/list - list products`)

	bot.Send(msg)
}


func defaultBehavior(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update, productService *product.Service) {
	outputMsg := "Here all products: \n\n"
	products := productService.List()

	for _, p := range products {
		outputMsg += p.Title
		outputMsg += "\n"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, outputMsg)

	bot.Send(msg)
}


