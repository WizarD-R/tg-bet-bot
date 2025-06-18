package main

import (
	"RushBananaBet/pkg/logger"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/yaml.v3"
)

var initDataMap map[string]interface{}

func init() {
	data, err := os.ReadFile("../../configs/config.yml")
	if err != nil {
		logger.Fatal("Cant read config", "", "", "main-init()", err)
	}

	err = yaml.Unmarshal(data, &initDataMap)
	if err != nil {
		logger.Fatal("Cant unmarshal data from config", "", "", "main-init()", err)
	}
}

func main() {
	botToken, ok := initDataMap["botToken"]
	if !ok {
		logger.Fatal("Cant read 'bot token' from configMap", "", "", "main-main()", nil)
	}
	bot, err := tgbotapi.NewBotAPI(botToken.(string))
	if err != nil {
		logger.Fatal("Error creating newBot", "", "", "main-main()", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
		}
	}
}
