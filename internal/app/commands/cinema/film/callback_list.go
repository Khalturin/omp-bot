package film

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
}

func (c *FilmCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("FilmCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}
	const defaultLimit = 4

	fmt.Println("OFFSET: ", parsedData.Offset)

	films, err := c.filmService.List(parsedData.Offset, defaultLimit)
	if err != nil {
		log.Println("Can't take list callback")
	}
	viewSize := defaultLimit
	if len(films) < defaultLimit {
		viewSize = len(films)
	}
	outputMsgText := fmt.Sprintf("List of the films (from %v to %v)\n\n",
		parsedData.Offset+1,
		parsedData.Offset+uint64(viewSize),
	)

	for _, p := range films {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	var msg tgbotapi.MessageConfig
	if len(films) == 0 {
		msg = tgbotapi.NewMessage(
			callback.Message.Chat.ID,
			fmt.Sprintf("No more films"),
		)
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("DemoSubdomainCommander.CallbackList: error sending reply message to chat - %v", err)
		}
		return
	} else {
		msg = tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)
	}

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: parsedData.Offset + defaultLimit,
	})

	newCallbackPath := path.CallbackPath{
		Domain:       "cinema",
		Subdomain:    "film",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", newCallbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
