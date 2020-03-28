package main

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const token = "947475124:AAFcOFjtk30GPhD3H0pMB5ELb41KAur5yEI"

var configPath = "config.json"
var config = loadConf(configPath)

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
	forward := make(chan tgbotapi.ForwardConfig)
	for {
		select {
		case update := <-updChan:
			text := update.Message.Text

			// chatID := update.Message.Chat.ID

			if update.Message.IsCommand() == true && !update.Message.Chat.IsGroup() {
				message := processCommand(update)
				bot.Send(message)
			} else {
				if text != "" {
					if update.Message.Chat.IsGroup() {
						go findTag(update.Message, forward)
					} else {

						replay := fmt.Sprintf("%s? You are so annoying %s!", update.Message.Text, update.Message.From.FirstName)
						m := tgbotapi.NewMessage(update.Message.Chat.ID, replay)
						m.ReplyToMessageID = update.Message.MessageID
						bot.Send(m)
					}
				}
			}
		case f := <-forward:
			go bot.Send(f)
		}

	}
}

func processCommand(update tgbotapi.Update) tgbotapi.MessageConfig {
	var answer string
	var args []string

	command := update.Message.Command()
	chatID := update.Message.Chat.ID
	args = strings.Fields(update.Message.CommandArguments())

	switch command {
	case "subscribe":
		err := addSubscribtion(chatID, args)
		if err != nil {
			log.Println(err)
		}
		answer = fmt.Sprintf("Got Command %v with args: %v", command, args)
	default:
		answer = "I don't know what do you want"
	}

	return tgbotapi.NewMessage(chatID, answer)
}

func findTag(message *tgbotapi.Message, forward chan tgbotapi.ForwardConfig) {
	msg := message.Text
	tags, err := getTagsList()
	if err != nil {
		log.Println(err)
	}
	for _, t := range tags {
		if strings.Contains(msg, t) {
			chats, _ := getChatByTag(t)
			fmt.Println(chats)
			for _, c := range chats {
				f := tgbotapi.NewForward(c, message.Chat.ID, message.MessageID)
				forward <- f
			}
		}
	}
}
