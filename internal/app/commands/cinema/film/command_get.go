package film

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *FilmCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	film, err := c.filmService.Get(idx)
	sendMsg := ""
	if err != nil {
		sendMsg = fmt.Sprintf("fail to get film with idx %d: %v", idx, err)
		log.Printf(sendMsg)
	} else {
		sendMsg = film.String()
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		sendMsg,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("FilmCommander.Get: error sending reply message to chat - %v", err)
	}
}
