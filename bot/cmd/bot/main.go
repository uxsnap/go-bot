package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/uxsnap/go-bot/internal/app/commands"
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
	commander := commands.NewCommander(bot, productService)

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		commander.HandleUpdate(&update)

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}

