package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/yaml.v3"
)

var initDataMap map[string]interface{}

func init() {
	data, err := os.ReadFile("../../configs/config.yml")
	if err != nil {
		// Log
	}

	err = yaml.Unmarshal(data, &initDataMap)
	if err != nil {
		// Log
	}
}

func main() {
	botToken, ok := initDataMap["botToken"]
	if !ok {
		// Log
	}
	bot, err := tgbotapi.NewBotAPI(botToken.(string))
	if err != nil {
		// Log
	}
	//bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
		}
	}
}
