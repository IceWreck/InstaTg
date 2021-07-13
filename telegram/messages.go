package telegram

import (
	"strings"
	"sync"
	"time"

	"github.com/IceWreck/InstaTG/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Telegram Documentation:
// When sending messages inside a particular chat, avoid sending more than one message per second.
// We may allow short bursts that go over this limit, but eventually you'll begin receiving 429 errors.
// If you're sending bulk notifications to multiple users, the API will not allow more than 30 messages
// per second or so. Consider spreading out notifications over large intervals of 8—12 hours for best results.
// Also note that your bot will not be able to send more than 20 messages per minute to the same group

// They dont specify anything about a channel but Ive seen errors after sending concurrent messages.
// So I'll implement a mutex for photos cause they're small and send fast and no mutex for videos.
// Retrys for both.

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

	var err error

	// Lock mutex, if its stuck in retry hell, then all of the other's sends wont be
	// because when it unlocks's then it means that telegram is ready to accept more msgs.
	m.Lock()
	for i := 0; i < 10; i++ {
		_, err = app.TGBot.Send(msg)

		if err != nil && strings.Contains(err.Error(), "Too Many Requests: retry") {
			app.Logger.Println("Telegram send waiting because ", err)
			time.Sleep(10 * time.Second)
		} else if err != nil {
			// not a rate limit so break
			app.Logger.Println("Error sending telegram photo. Error: ", err)
			break
		} else {
			break
		}
	}
	m.Unlock()
	return err
}

func SendVideo(app *config.Application, videoLocation string, caption string) error {
	msg := tgbotapi.NewVideoUpload(app.Config.TelegramChannel, videoLocation)
	msg.Caption = caption
	var err error
	for i := 0; i < 10; i++ {
		_, err = app.TGBot.Send(msg)
		if err != nil && strings.Contains(err.Error(), "Too Many Requests: retry") {
			app.Logger.Println("Telegram send waiting because ", err)
			time.Sleep(10 * time.Second)
		} else if err != nil {
			app.Logger.Println("Error sending telegram video. Error: ", err)
			break
		} else {
			break
		}
	}
	return err
}
