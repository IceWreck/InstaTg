package telegram

import (
	"github.com/IceWreck/InstaTG/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func ConnectionCheck(app *config.Application) error {
	user, err := app.TGBot.GetMe()
	if err == nil {
		app.Logger.Println("Telegram authenticated. Username: ", user.UserName)
	} else {
		// Panic is intentional because we dont want it to try making requests if it fails.
		// And if connection loss is just temporary, then systemd will restart anyways.
		app.Logger.Panicln("Telegram authentication error: ", err)
	}
	return err
}

func NewConnection(app *config.Application) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(app.Config.TelegramBotToken)
	if err != nil {
		app.Logger.Panicln("Telegram API connection error: ", err)
	}
	bot.Debug = false
	return bot
}
