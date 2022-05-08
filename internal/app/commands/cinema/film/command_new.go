package film

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
	"log"
)

/*
	/new__cinema__film
	{"title":"Myfilm",
	"year":"2022",
	"genre":"Sci-fi",
	"director":"Khan"}
*/
func (c *FilmCommander) New(inputMessage *tgbotapi.Message) {
	film := cinema.Film{}
	args := inputMessage.CommandArguments()
	err := json.Unmarshal([]byte(args), &film)
	if err != nil {
		outputMsg := "Can't Unmarshal data, err: "
		outputMsg += fmt.Sprintf("%s\n", err)
		outputMsg += "format must be: \n" +
			"/new__cinema__film\n" +
			"{\"title\":\"Title_film\",\n" +
			"\"year\":2022,\n" +
			"\"genre\":\"film_genre\",\n" +
			"\"director\":\"SomeOne\"}"
		log.Println(outputMsg)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
		c.bot.Send(msg)
		return
	}
	lengthEntities := len(cinema.AllEntities)
	film.ID = uint64(lengthEntities)

	c.filmService.Create(film)

	_, err2 := c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "film added"))
	if err2 != nil {
		log.Println("Can't send message from New")
	}

	return
}
