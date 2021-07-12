package telegram

import (
	"github.com/IceWreck/InstaTG/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendTextMessage(app *config.Application, txt string) error {
	msg := tgbotapi.NewMessage(app.Config.TelegramChannel, txt)
	_, err := app.TGBot.Send(msg)
	if err != nil {
		app.Logger.Println("Error sending telegram message. Error: ", err)
	}
	return err
}

func SendImage(app *config.Application, imgLocation string, caption string) error {
	msg := tgbotapi.NewPhotoUpload(app.Config.TelegramChannel, imgLocation)
	msg.Caption = caption
	_, err := app.TGBot.Send(msg)
	if err != nil {
		app.Logger.Println("Error sending telegram photo. Error: ", err)
	}
	return err
}

func SendVideo(app *config.Application, videoLocation string, caption string) error {
	msg := tgbotapi.NewVideoUpload(app.Config.TelegramChannel, videoLocation)
	msg.Caption = caption
	_, err := app.TGBot.Send(msg)
	if err != nil {
		app.Logger.Println("Error sending telegram video. Error: ", err)
	}
	return err
}
