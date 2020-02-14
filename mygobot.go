package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const token = "947475124:AAFcOFjtk30GPhD3H0pMB5ELb41KAur5yEI"

func main() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Println(bot.Self.UserName, bot.Self.ID)

	var updatecfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(1)
	updatecfg.Timeout = 60
	updChan, err := bot.GetUpdatesChan(updatecfg)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case update := <-updChan:
			text := update.Message.Text
			if text != "" {
				replay := fmt.Sprintf("%s? You are so annoying %s!", update.Message.Text, update.Message.From.FirstName)
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, replay))
			}
		}
	}
}
