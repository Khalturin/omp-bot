package film

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/cinema/film"
	"log"
)

type FilmCommander struct {
	bot         *tgbotapi.BotAPI
	filmService *film.DummyFilmService
}

func NewFilmCommander(
	bot *tgbotapi.BotAPI,
) *FilmCommander {
	filmService := film.NewDummyFilmService()

	return &FilmCommander{
		bot:         bot,
		filmService: filmService,
	}
}

func (c *FilmCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("FilmCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *FilmCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "new":
		c.New(msg)
	case "delete":
		c.Delete(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
