package cinema

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/cinema/film"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type CinemaCommander struct {
	bot           *tgbotapi.BotAPI
	FilmCommander Commander
}

func NewCinemaCommander(
	bot *tgbotapi.BotAPI,
) *CinemaCommander {
	return &CinemaCommander{
		bot: bot,
		// CinemaCommander
		FilmCommander: film.NewFilmCommander(bot),
	}
}

func (c *CinemaCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "film":
		c.FilmCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("CinemaCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *CinemaCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "film":
		c.FilmCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("CinemaCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
