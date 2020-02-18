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
			chatID := update.Message.Chat.ID
			var waitResponse bool = false
			if update.Message.IsCommand() == true {
				processCommand(bot, update.Message.Command(), chatID, &waitResponse)
				bot.
			} else {
				if text != "" {
					replay := fmt.Sprintf("%s? You are so annoying %s!", update.Message.Text, update.Message.From.FirstName)
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, replay))
				}
			}
		}
	}
}

func processCommand(bot *tgbotapi.BotAPI, command string, chatID int64, wait *bool) {
	bot.Send(tgbotapi.NewMessage(chatID, "got command"))
	var replay string
	switch command {
	case "subscribe":
		fmt.Println(chatID)
		fmt.Println(command)
		replay = command
	default:
		replay = "I don't know what do you want"
	}

	bot.Send(tgbotapi.NewMessage(chatID, replay))
}
