package film

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *FilmCommander) List(inputMessage *tgbotapi.Message) {
	const defaultLimit = 4

	products, err := c.filmService.List(0, defaultLimit)
	if err != nil {
		log.Printf("Can't get list, err - %v", err)
	}

	viewSize := defaultLimit
	if len(products) < defaultLimit {
		viewSize = len(products)
	}
	outputMsgText := fmt.Sprintf("List of the films (from %v to %v)\n\n", 1, viewSize)
	for _, p := range products {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: defaultLimit,
	})

	callbackPath := path.CallbackPath{
		Domain:       "cinema",
		Subdomain:    "film",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err2 := c.bot.Send(msg)
	if err2 != nil {
		log.Printf("DemoSubdomainCommander.List: error sending reply message to chat - %v", err)
	}

}
