package film

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
	"log"
	"strconv"
)

func (c *FilmCommander) Edit(inputMessage *tgbotapi.Message) {
	film := cinema.Film{}
	args := inputMessage.CommandArguments()
	err := json.Unmarshal([]byte(args), &film)
	if err != nil {
		outputMsg := "Can't Unmarshal data, err: "
		outputMsg += fmt.Sprintf("%s\n", err)
		outputMsg += "format must be: \n" +
			"/edit__cinema__film\n" +
			"{\"id\":0,\n" +
			"\"title\":\"Title_film\",\n" +
			"\"year\":2022,\n" +
			"\"genre\":\"film_genre\",\n" +
			"\"director\":\"SomeOne\"}"

		log.Println(outputMsg)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
		c.bot.Send(msg)
		return
	}

	if film.ID > uint64(len(cinema.AllEntities)) {
		outputMsg := "Incorrect ID id must be < " + strconv.Itoa(len(cinema.AllEntities))
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
		c.bot.Send(msg)
		return
	}

	c.filmService.Update(film.ID, film)

	_, err2 := c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "film edited"))
	if err2 != nil {
		log.Println("Can't send message from Edit")
	}

	return
}
