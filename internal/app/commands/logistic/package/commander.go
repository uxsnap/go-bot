package packageCommander

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	logisticPackage "github.com/ozonmp/omp-bot/internal/service/logistic/package"
)

type PackageCommander struct {
	bot              *tgbotapi.BotAPI
	packageService 	 *logisticPackage.DummyPackageService
}

func NewPackageCommander(
	bot *tgbotapi.BotAPI,
) *PackageCommander {
	packageService := logisticPackage.NewDummyPackageService()

	return &PackageCommander{
		bot:              bot,
		packageService: packageService,
	}
}

func (c *PackageCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("PackageCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *PackageCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "new":
		c.New(msg)
	case "get":
		c.Get(msg)
	case "list":
		c.List(msg)
	case "delete":
		c.Delete(msg)
	}
}
