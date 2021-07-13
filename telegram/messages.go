package telegram

import (
	"sync"

	"github.com/IceWreck/InstaTG/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// mutex lock to not send files at the same time
// looks like this package isnt concurrency safe
// https://github.com/go-telegram-bot-api/telegram-bot-api/issues/273
var m sync.Mutex

func SendTextMessage(app *config.Application, txt string) error {
	msg := tgbotapi.NewMessage(app.Config.TelegramChannel, txt)
	m.Lock()
	_, err := app.TGBot.Send(msg)
	m.Unlock()
	if err != nil {
		app.Logger.Println("Error sending telegram message. Error: ", err)
	}
	return err
}

func SendImage(app *config.Application, imgLocation string, caption string) error {
	msg := tgbotapi.NewPhotoUpload(app.Config.TelegramChannel, imgLocation)
	msg.Caption = caption
	m.Lock()
	_, err := app.TGBot.Send(msg)
	m.Unlock()
	if err != nil {
		app.Logger.Println("Error sending telegram photo. Error: ", err)
	}
	return err
}

func SendVideo(app *config.Application, videoLocation string, caption string) error {
	msg := tgbotapi.NewVideoUpload(app.Config.TelegramChannel, videoLocation)
	msg.Caption = caption
	m.Lock()
	_, err := app.TGBot.Send(msg)
	m.Unlock()
	if err != nil {
		app.Logger.Println("Error sending telegram video. Error: ", err)
	}
	return err
}
