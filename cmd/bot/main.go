package main

import (
	"RushBananaBet/internal/app"
	user "RushBananaBet/internal/model"
	"RushBananaBet/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"gopkg.in/yaml.v3"
)

var initData struct {
	Admins     []string `yaml:"admins"`
	BotToken   string   `yaml:"botToken"`
	LogLevel   int      `yaml:"log-level"`
	Production bool     `yaml:"production"`
}

func init() {
	data, err := os.ReadFile("../../configs/config.yml")
	if err != nil {
		logger.Fatal("Cant read config", "", "", "main-init()", err)
	}

	err = yaml.Unmarshal(data, &initData)
	if err != nil {
		logger.Fatal("Cant unmarshal data from config", "", "", "main-init()", err)
	}

	// Set admins
	user.Admins = initData.Admins

	// Init logger
	logger.InitLogger(initData.LogLevel, initData.Production)
}

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	bot := app.NewBot(initData.BotToken)
	bot.Start(stop)
}
