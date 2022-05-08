package film

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *FilmCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	number, err := strconv.Atoi(args)
	if err != nil {
		log.Println("Can't convert arg!")
		return
	}

	c.filmService.Remove(uint64(number))

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("film with id: %d deleted", number))
	_, err2 := c.bot.Send(msg)
	if err2 != nil {
		log.Println("Can't send message from Delete")
	}

	return
}
