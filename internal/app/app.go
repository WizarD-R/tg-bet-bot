package app

import (
	"RushBananaBet/internal/handler"
	"RushBananaBet/pkg/logger"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	handler *handler.Handler
}

func (b *Bot) Start(stop chan os.Signal) {
	logger.Info("Bot started", "", "", "", nil)

	b.StartPolling()

	<-stop
	logger.Info("Bot stoped", "", "", "", nil)

	b.bot.StopReceivingUpdates()
}

func (b *Bot) StartPolling() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		go b.RouteUpdate(update)
	}
}

func (b *Bot) RouteUpdate(update tgbotapi.Update) {
	if update.Message != nil {
		switch {
		case update.Message.Text == "/start":
			handler.Start()
			// msg := handlers.HandleStart(update.Message)
			// bot.Send(msg)
			// // другие команды...
		case update.Message.Text == "/create-event":
			// ds
		case update.Message.Text == "/add-result":
			//ds
		case update.Message.Text == "/finish-tournament":
			// ыв
		case update.Message.Text == "/my-predictions":
			//ds
		case strings.Contains(update.Message.Text, "/match"):
			// dsd
		}
	}
}

func NewBot(botToken string, handler *handler.Handler) Bot {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		logger.Fatal("Error creating newBot", "", "", "main-main()", err)
	}
	return Bot{
		bot:     bot,
		handler: handler,
	}
}

// func BuildKeyboard(username string) tgbotapi.ReplyKeyboardMarkup {
// 	var rows [][]tgbotapi.KeyboardButton

// 	// Админские кнопки
// 	if user.IsAdmin(username) {
// 		rows = append(rows, tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("➕ Добавить ивент"),
// 		))
// 		rows = append(rows, tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("🎯 Добавить результат"),
// 		))
// 		rows = append(rows, tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("🏁 Завершить турнир"),
// 		))
// 	}

// 	// Пользовательская кнопка
// 	rows = append(rows, tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("📄 Мои ставки"),
// 	))

// 	// Кнопки матчей
// 	rows = append(rows, tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("⚔️ Матч 1"),
// 		tgbotapi.NewKeyboardButton("⚔️ Матч 2"),
// 	))

// 	return tgbotapi.NewReplyKeyboard(rows...)
// }
