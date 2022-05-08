package film

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *FilmCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__cinema__film - help\n"+
			"/list__cinema__film - list films\n"+
			"/get__cinema__film - get film\n"+
			"/delete__cinema__film - delete film\n"+
			"/new__cinema__film - new film\n"+
			"/edit__cinema__film - edit film",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}
